package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodosRoute(t *testing.T) {
	router := gin.Default()
	RegisterTodoEndpoints(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
