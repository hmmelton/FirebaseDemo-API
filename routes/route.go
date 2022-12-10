package routes

// Import the necessary modules and dependencies
import (
  "github.com/gin-gonic/gin"
  "github.com/hmmelton/firebasedemoapi/handlers"
)

// SetupRoutes defines the API routes and attaches the route handlers
func SetupRoutes(router *gin.Engine) {
  router.GET("/", handlers.GetRoot)
  router.GET("/users", handlers.GetUsers)
  router.POST("/users", handlers.PostUser)
}