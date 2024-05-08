package user

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	router := gin.New()
	handler := &Handler{}
	router.POST("/api/v1/users", handler.Create)
	// Create a recorder to capture the response

	m := Request{
		Name:    "foo",
		Surname: "bar",
	}

	data, err := json.Marshal(&m)

	assert.NoError(t, err)

	reader := bytes.NewReader(data)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/users",
		reader,
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var response map[string]any
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
}
