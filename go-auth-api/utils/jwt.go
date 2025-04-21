package utils

import (

    "time"
    "os"
    
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_secret")


func GenerateJWT(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(jwtKey)
}
