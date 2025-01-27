package routers

import (
	userhandler "github.com/Saadyyyy/money-bank-be/pkg/user/user_handler"
	userrepository "github.com/Saadyyyy/money-bank-be/pkg/user/user_repository"
	userservice "github.com/Saadyyyy/money-bank-be/pkg/user/user_service"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Register(db *sqlx.DB, echo *echo.Echo, rdb *redis.Client) {
	RegisterUser(db, echo, rdb)
}

func RegisterUser(db *sqlx.DB, echo *echo.Echo, rdb *redis.Client) {
	repo := userrepository.NewUserRepository(db)
	service := userservice.NewUserService(repo, db)
	handler := userhandler.NewUserHandler(service)

	user := echo.Group("/user")
	user.POST("/register", handler.RegisterUser).Name = "RegisterUser"
	user.POST("/login", handler.LoginUser).Name = "LoginUser"
}
