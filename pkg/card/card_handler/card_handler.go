package cardhandlers

import (
	"fmt"
	"strconv"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
	"github.com/Saadyyyy/money-bank-be/utils/https"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CardHandlerInterfaceImpl struct {
	service domain.CardService
}

func NewCardHandler(service domain.CardService) *CardHandlerInterfaceImpl {
	return &CardHandlerInterfaceImpl{service: service}
}

func (h *CardHandlerInterfaceImpl) CreateCard(e echo.Context) error {
	fName := "CardHttpHandler.CreateCard"
	ctx := e.Request().Context()
	type ReqBody struct {
		CardId   int64  `json:"card_id"`
		CardName string `json:"card_name"`
	}

	req := ReqBody{}
	if err := e.Bind(&req); err != nil {
		return https.WriteBadRequestResponse(e, https.ResponseBadRequestError)
	}

	// validate request body
	if err := validator.New().Struct(&req); err != nil {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, err)
	}

	resp := domain.Card{
		CardId:   req.CardId,
		CardName: req.CardName,
	}
	id, err := h.service.CreateCard(ctx, resp)

	if err != nil {
		return https.WriteServerErrorResponse(e, fName, err)

	}
	return https.WriteOkResponse(e, fmt.Sprintf("Berhasil membuat akun dengan id %d", id))
}
func (h *CardHandlerInterfaceImpl) UpdateCard(e echo.Context) error {
	fName := "CardHttpHandler.UpdateCard"
	ctx := e.Request().Context()
	CardIdStr := e.QueryParam("card_id")
	if CardIdStr == "" {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, fmt.Errorf("missing or invalid card_id parameter"))
	}

	cardId, err := strconv.ParseInt(CardIdStr, 10, 64)
	if err != nil || cardId <= 0 {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, fmt.Errorf("invalid card_id parameter"))
	}

	type ReqBody struct {
		CardName string `json:"card_name"`
	}

	req := ReqBody{}
	if err := e.Bind(&req); err != nil {
		return https.WriteBadRequestResponse(e, https.ResponseBadRequestError)
	}

	// validate request body
	if err := validator.New().Struct(&req); err != nil {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, err)
	}

	resp := domain.Card{
		CardId:   cardId,
		CardName: req.CardName,
	}

	_, errService := h.service.UpdatedCard(ctx, resp)
	if errService != nil {
		return https.WriteServerErrorResponse(e, fName, err)

	}
	return https.WriteOkResponse(e, fmt.Sprintf("Succes Update Card with ID: %d", cardId))
}

func (h *CardHandlerInterfaceImpl) DeleteCard(e echo.Context) error {
	fName := "CardHttpHandler.DeleteCard"
	ctx := e.Request().Context()

	cardIdStr := e.QueryParam("card_id")
	if cardIdStr == "" {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, fmt.Errorf("missing or invalid card_id parameter"))
	}

	cardId, err := strconv.ParseInt(cardIdStr, 10, 64)
	if err != nil {
		return https.WriteBadRequestResponseWithErrMsg(e, https.ResponseBadRequestError, fmt.Errorf("invalid card_id parameter"))
	}
	if cardId == 0 || cardId < 0 {
		return https.WriteNotFoundResponse(e, "card_id Not Found")
	}

	errServ := h.service.DeleteCard(ctx, cardId)
	if errServ != nil {
		return https.WriteServerErrorResponse(e, fName, err)
	}
	return https.WriteOkResponse(e, fmt.Sprintf("Successc delete card with id : %d", cardId))

}
