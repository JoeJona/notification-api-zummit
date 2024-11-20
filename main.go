package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendNotification(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Course is available. You can continue your studies, enjoy."})
}

func main() {

	router := gin.Default()

	router.GET("/send-notification", SendNotification)

	router.Run("localhost:8080")

}
