these libraries installed by me 

```
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
```

setup

```
go-auth-api/
│
├── main.go
├── config/
│   └── db.go       # MongoDB connection
├── models/
│   └── user.go     # User struct
├── handlers/
│   └── auth.go     # Signup/Login logic
└── utils/
    └── jwt.go      # Token generation & validation

```