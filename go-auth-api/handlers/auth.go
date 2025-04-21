package handlers

import (
    "context"
    "net/http"

    "go-auth-api/config"
    "go-auth-api/models"
    "go-auth-api/utils"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
)

func Signup(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Hash error"})
        return
    }
    user.Password = string(hashedPassword)

    collection := config.DB.Collection("users")
    _, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert error"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func Login(c *gin.Context) {
    var input models.User
    var dbUser models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    collection := config.DB.Collection("users")
    err := collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&dbUser)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

    token, err := utils.GenerateJWT(dbUser.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
