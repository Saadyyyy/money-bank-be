package cardservice

import (
	"context"
	"fmt"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
)

type CardServiceInterfaceImpl struct {
	repo domain.CardRepository
}

// DeleteCard implements domain.CardService.
func (c *CardServiceInterfaceImpl) DeleteCard(ctx context.Context, id int64) error {
	err := c.repo.DeleteCard(ctx, id)
	if err != nil {
		return fmt.Errorf("Error", err)
	}
	return nil
}

// GetCard implements domain.CardService.
func (c *CardServiceInterfaceImpl) GetCard(ctx context.Context, searchCriteria map[string]interface{}) (result []domain.Card, err error) {
	panic("unimplemented")
}

// GetCardById implements domain.CardService.
func (c *CardServiceInterfaceImpl) GetCardById(ctx context.Context, id int64) (result domain.Card, err error) {
	panic("unimplemented")
}

// CreateCard implements domain.CardService.
func (c *CardServiceInterfaceImpl) CreateCard(ctx context.Context, card domain.Card) (id int64, err error) {
	if card.CardName == "" {
		return 0, fmt.Errorf("card_name cannot null")
	}
	id, err = c.repo.CreateCard(ctx, card)
	if err != nil {
		return
	}
	fmt.Println(id)

	return id, nil
}

// UpdatedCard implements domain.CardService.
func (c *CardServiceInterfaceImpl) UpdatedCard(ctx context.Context, card domain.Card) (id int64, err error) {
	id, err = c.repo.UpdateCard(ctx, card)
	if err != nil {
		return
	}

	return id, err
}

func NewCardService(repo domain.CardRepository) domain.CardService {
	return &CardServiceInterfaceImpl{repo: repo}
}
