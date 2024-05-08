package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Client struct {
	config     Config
	connection *gorm.DB
	mu         *sync.RWMutex
}

func (c *Client) Shutdown() error {
	c.mu.RLock()
	db, err := c.connection.DB()
	c.mu.RUnlock()

	if err != nil {
		return err
	}

	return db.Close()
}

func New(config Config) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	c := &Client{
		config: config,
		mu:     &sync.RWMutex{},
	}

	db, err := gorm.Open(c.dialect(), &gorm.Config{})

	if err != nil {
		err = fmt.Errorf("failed to open database connection: %w", err)
	}

	c.connection = db
	connection, err := db.DB()

	if err != nil {
		err = fmt.Errorf("failed to get underlying database connection: %w", err)
	}

	for {
		if err = connection.Ping(); err == nil {
			break
		}

		select {
		case <-time.After(500 * time.Millisecond):
			continue
		case <-ctx.Done():
			defer connection.Close()
			return nil, errors.New("unable to connect to postgres client, context deadline exceeded")
		}
	}

	return c, nil
}

func (c *Client) dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		c.config.Host,
		c.config.Username,
		c.config.Password,
		c.config.DB,
		c.config.Port,
	)
}

func (c *Client) dialect() gorm.Dialector {
	return postgres.Open(c.dsn())
}

func (c *Client) DB() *sql.DB {
	c.mu.RLock()
	defer c.mu.RUnlock()

	db, _ := c.connection.DB()

	return db
}

func (c *Client) Ping() error {
	return c.DB().Ping()
}

func (c *Client) Debug() *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Debug()

	return c
}

func (c *Client) AutoMigrate(dst ...any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.connection.AutoMigrate(dst...)
}

func (c *Client) Close() error {
	return c.DB().Close()
}

func (c *Client) Error() error {
	return c.connection.Error
}
