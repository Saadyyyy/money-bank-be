package cardrepository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Saadyyyy/money-bank-be/pkg/domain"
	"github.com/jmoiron/sqlx"
)

const (
	queryCreateCard = `
		INSERT INTO cards(card_name,created_at)VALUES($1,$2) returning card_id
	`

	queryGetCard = `
		SELECT card_id,card_name from cards where deleted_at is null 
	`

	queryUpdateCard = `
		UPDATE cards SET card_name = $1, updated_at = $2 where card_id = $3 and deleted_at is null
	`

	queryGetCardById = `
		SELECT card_id,card_name from cards where card_id=$1 and deleted_at is null
	`

	queryDeleteCard = `
		UPDATE cards SET deleted_at = $1 where card_id =$2 and deleted_at is null
	`
)

type CardRepositoryInterfaceImpl struct {
	db *sqlx.DB
}

// CreateCard implements domain.CardRepository.
func (c *CardRepositoryInterfaceImpl) CreateCard(ctx context.Context, card domain.Card) (id int64, err error) {
	createdAt := time.Now()
	err = c.db.QueryRowContext(ctx, queryCreateCard, card.CardName, createdAt).Scan(&id)

	if err != nil {
		err = fmt.Errorf("query queryCreateCard error: %+v", err)
		return
	}
	fmt.Println(id)
	return id, nil
}

// DeleteCard implements domain.CardRepository.
func (c *CardRepositoryInterfaceImpl) DeleteCard(ctx context.Context, id int64) (err error) {
	deletedAt := time.Now()
	_, err = c.db.ExecContext(ctx, queryDeleteCard, deletedAt, id)
	if err != nil {
		return fmt.Errorf("query queryDeleteCard err", err)

	}
	if id == 0 {
		return fmt.Errorf("Fail to get ID")
	}
	return nil
}

// GetCard implements domain.CardRepository.
func (c *CardRepositoryInterfaceImpl) GetCard(ctx context.Context, searchCriteria map[string]interface{}) (result []domain.Card, err error) {
	sqlQuery := queryGetCard + searchCriteria["custom_query"].(string)
	rows, err := c.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("queryGetAllCategory err: %+v", err)
			return
		}
		err = nil
		return
	}
	defer rows.Close()

	var card domain.Card
	for rows.Next() {
		err = rows.Scan(&card.CardId, &card.CardName)
		if err != nil {
			err = fmt.Errorf("fail to scan card", err)
		}
		result = append(result, card)
	}

	if err = rows.Err(); err != nil {
		err = fmt.Errorf("rows iteration err: %+v", err)
		return nil, err
	}
	return result, nil
}

// GetCardById implements domain.CardRepository.
func (c *CardRepositoryInterfaceImpl) GetCardById(ctx context.Context, id int64) (result domain.Card, err error) {
	err = c.db.QueryRowContext(ctx, queryGetCardById, id).Scan(&result.CardId, &result.CardName)
	if err != nil {
		return domain.Card{}, fmt.Errorf("query queryGetCardById error :", err)
	}
	if id == 0 || id < 0 {
		return domain.Card{}, fmt.Errorf("Id not found")
	}
	return
}

// UpdateCard implements domain.CardRepository.
func (c *CardRepositoryInterfaceImpl) UpdateCard(ctx context.Context, card domain.Card) (id int64, err error) {
	updateAd := time.Now()
	_, err = c.db.ExecContext(ctx, queryUpdateCard, card.CardName, updateAd, card.CardId)
	if err != nil {
		return 0, fmt.Errorf("query queryUpdateCard error :", err)
	}
	return
}

func NewCardRepository(db *sqlx.DB) domain.CardRepository {
	return &CardRepositoryInterfaceImpl{db: db}
}
