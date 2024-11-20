package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSendNotification(t *testing.T) {
	r := SetUpRouter()
	r.GET("/send-notification", SendNotification)
	http.NewRequest("GET", "/send-notification", nil)
	w := httptest.NewRecorder()

	assert.Equal(t, w.Code, http.StatusOK)

}
