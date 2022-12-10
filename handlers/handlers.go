package handlers

// Import the necessary modules and dependencies
import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/hmmelton/firebasedemoapi/models"
)

// GetRoot is a route handler for HTTP GET requests to the "/" path
func GetRoot(context *gin.Context) {
  // Return a static message as the response to the GET request
  context.JSON(http.StatusOK, gin.H{
    "message": "Hello, world!",
  })
}

// GetUsers is a route handler for HTTP GET requests to the "/users" path
func GetUsers(context *gin.Context) {
  // Get a list of users from the data store
  users := models.GetUsers()

  // Return the list of users as the response to the GET request
  context.JSON(http.StatusOK, gin.H{
    "users": users,
  })
}

// PostUser is a route handler for HTTP POST requests to the "/users" path
func PostUser(context *gin.Context) {
  // Parse the request body to get the user data
  var user models.User
  context.BindJSON(&user)

  // Save the user data to the data store
  models.SaveUser(user)

  // Return a success message as the response to the POST request
  context.JSON(http.StatusCreated, gin.H{
    "message": "User created successfully",
  })
}