package domain

import (
	"context"
)

// default strut card same like tabel Card
type Card struct {
	CardId    int64
	CardName  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

// For respons card
type CardRespons struct {
	CardId   int64  `json:"card_id"`
	CardName string `json:"card_name"`
}

// Repository for card
type CardRepository interface {
	CreateCard(ctx context.Context, card Card) (id int64, err error)
	UpdateCard(ctx context.Context, card Card) (id int64, err error)
	GetCardById(ctx context.Context, id int64) (result Card, err error)
	DeleteCard(ctx context.Context, id int64) (err error)
	GetCard(ctx context.Context, searchCriteria map[string]interface{}) (result []Card, err error)
}

type CardService interface {
	CreateCard(ctx context.Context, card Card) (id int64, err error)
	UpdatedCard(ctx context.Context, card Card) (id int64, err error)
	GetCardById(ctx context.Context, id int64) (result Card, err error)
	DeleteCard(ctx context.Context, id int64) error
	GetCard(ctx context.Context, searchCriteria map[string]interface{}) (result []Card, err error)
}
