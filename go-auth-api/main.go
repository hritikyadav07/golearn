
package main

import (
    "log" // Import log package
    "go-auth-api/config"
    "go-auth-api/handlers"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv" // Import godotenv
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: Could not load .env file") // Log warning instead of fatal if .env is optional
    }

    config.ConnectDB() // Connect DB after loading env

    r := gin.Default()

    r.POST("/signup", handlers.Signup)
    r.POST("/login", handlers.Login)

    r.Run(":4000")
}