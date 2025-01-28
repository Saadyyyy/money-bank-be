package routers

import (
	cardhandlers "github.com/Saadyyyy/money-bank-be/pkg/card/card_handler"
	cardrepository "github.com/Saadyyyy/money-bank-be/pkg/card/card_repository"
	cardservice "github.com/Saadyyyy/money-bank-be/pkg/card/card_service"
	userhandler "github.com/Saadyyyy/money-bank-be/pkg/user/user_handler"
	userrepository "github.com/Saadyyyy/money-bank-be/pkg/user/user_repository"
	userservice "github.com/Saadyyyy/money-bank-be/pkg/user/user_service"
	"github.com/Saadyyyy/money-bank-be/utils/middleware"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Register(db *sqlx.DB, echo *echo.Echo, rdb *redis.Client) {
	RegisterUser(db, echo)
	RegisterCard(db, echo)
}

func RegisterUser(db *sqlx.DB, echo *echo.Echo) {
	repo := userrepository.NewUserRepository(db)
	service := userservice.NewUserService(repo, db)
	handler := userhandler.NewUserHandler(service)

	user := echo.Group("/user")
	user.POST("/register", handler.RegisterUser).Name = "RegisterUser"
	user.POST("/login", handler.LoginUser).Name = "LoginUser"
}

func RegisterCard(db *sqlx.DB, echo *echo.Echo) {
	repo := cardrepository.NewCardRepository(db)
	service := cardservice.NewCardService(repo)
	handler := cardhandlers.NewCardHandler(service)

	card := echo.Group("/card")
	card.POST("/create", handler.CreateCard, middleware.JWTMiddleware()).Name = "CreateCard"
	card.POST("/update", handler.UpdateCard, middleware.JWTMiddleware()).Name = "UpdateCard"
	card.POST("/delete", handler.DeleteCard, middleware.JWTMiddleware()).Name = "DeleteCard"
}
