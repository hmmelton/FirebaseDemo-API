package main

// Import the necessary modules and dependencies
import (
  "github.com/gin-gonic/gin"
  "github.com/hmmelton/firebasedemoapi/routes"
)

func main() {
  // Create a new Gin router
  router := gin.Default()

  // Define the API routes and attach the route handlers
  routes.SetupRoutes(router)

  // Start the server and listen for incoming requests
  router.Run(":443")
}