package models

// Import the necessary modules and dependencies
import (
  "sync"
)

// User represents a user in the system
type User struct {
  ID   int    `json:"id"`
  Name string `json:"name"`
  Age  int    `json:"age"`
}

var (
  // users is a slice of User objects
  users []User
  // mu is a mutex for synchronizing access to the users slice
  mu sync.Mutex
)

// SaveUser saves a new user to the data store
func SaveUser(user User) {
  // Lock the mutex to synchronize access to the users slice
  mu.Lock()
  defer mu.Unlock()

  // Append the user to the slice
  users = append(users, user)
}

// GetUsers returns a list of all users from the data store
func GetUsers() []User {
  // Lock the mutex to synchronize access to the users slice
  mu.Lock()
  defer mu.Unlock()

  // Return a copy of the users slice
  return append([]User{}, users...)
}
