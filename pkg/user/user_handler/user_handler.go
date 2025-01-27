package userhandler

import (
	"fmt"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
	"github.com/Saadyyyy/money-bank-be/utils/https"
	"github.com/Saadyyyy/money-bank-be/utils/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase domain.UserService
}

func NewUserHandler(usecase domain.UserService) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) RegisterUser(e echo.Context) error {
	fName := "UserHttpHandler.RegisterUser"
	ctx := e.Request().Context()

	type ReqBody struct {
		UserId   int64  `json:"user_id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := ReqBody{}
	if err := e.Bind(&req); err != nil {
		return https.WriteBadRequestResponse(e, https.ResponseBadRequestError)
	}

	// validate request body
	if err := validator.New().Struct(&req); err != nil {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, err)
	}

	resp := domain.Users{
		UserId:   req.UserId,
		Username: req.Username,
		Password: req.Password,
	}

	userId, err := h.usecase.RegisterUser(ctx, resp)
	if err != nil {
		return https.WriteServerErrorResponse(e, fName, err)
	}

	return https.WriteOkResponse(e, fmt.Sprintf("Berhasil membuat akun dengan id %d", userId))
}

func (h *UserHandler) LoginUser(e echo.Context) error {
	fName := "UserHttpHandler.LoginUser"
	ctx := e.Request().Context()

	type reqBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := reqBody{}
	if err := e.Bind(&req); err != nil {
		return https.WriteBadRequestResponse(e, https.ResponseBadRequestError)

	}
	if err := validator.New().Struct(&req); err != nil {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, err)
	}

	users, token, err := h.usecase.LoginUser(ctx, req.Username)
	if err != nil {
		return https.WriteServerErrorResponse(e, fName, err)
	}
	middleware.SetTokenCookie(e, token)

	respons := domain.LoginUser{
		UserId:   users.UserId,
		Username: users.Username,
		Token:    token,
	}

	return https.WriteOkResponse(e, respons)
}
