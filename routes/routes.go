package routes

import (
	controller "task1/controllers"

	"github.com/gin-gonic/gin"
)

func RouterHandling() {
	router := gin.Default()

	router.POST("/encode",controller.Encode)
	router.POST("/decode",controller.Decode)
	
	router.Run("localhost:8080")
}