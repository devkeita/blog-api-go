package main

import (
	"blog-api/config"
	"blog-api/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	userHandler := wire.InitUserAPI(db)

	e := echo.New()
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}

	// 認証
	e.POST("/signup", userHandler.SignUp)
	e.POST("/login", userHandler.Login)

	// users
	u := e.Group("/users")
	u.GET("/", userHandler.GetAllUser)
	u.GET("/:id", userHandler.GetUser, middleware.JWTWithConfig(jwtConfig))

	e.Logger.Fatal(e.Start(":8000"))
}
