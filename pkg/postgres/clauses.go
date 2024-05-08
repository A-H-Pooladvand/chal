package postgres

import (
	"context"
	"fmt"
)

func (c *Client) Where(query any, args ...any) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Where(query, args...)

	return c
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.WithContext(ctx)

	return c
}

func (c *Client) Create(value any) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Create(value)

	return c
}

func (c *Client) HasError() bool {
	return c.connection.Error != nil
}

func (c *Client) Affected() bool {
	return c.connection.RowsAffected > 0
}

func (c *Client) InsertError() error {
	if c.HasError() {
		return c.Error()
	}

	if !c.Affected() {
		return fmt.Errorf("failed to insert into the database %v rows affected", c.connection.RowsAffected)
	}

	return nil
}

func (c *Client) Limit(limit int) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Limit(limit)

	return c
}

func (c *Client) Offset(offset int) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Offset(offset)

	return c
}

func (c *Client) Find(dest any, conds ...any) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.Find(dest, conds...)

	return c
}

func (c *Client) First(dest any, conds ...any) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connection = c.connection.First(dest, conds...)

	return c
}
