package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/usernamesalah/game-currency/pkg/models"
)

// CurrencyService service interface.
type CurrencyService interface {
	ListCurrency(ctx context.Context) ([]models.Currency, error)
	GetCurrency(ctx context.Context, id int64) (models.Currency, error)
	CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error)
	DeleteCurrency(ctx context.Context, id int64) error
	UpdateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error)
}

type currencyService struct {
	db *sqlx.DB
}

// NewCurrencyService returns an initialized CurrencyService implementation.
func NewCurrencyService(db *sqlx.DB) CurrencyService {
	return &currencyService{db: db}
}

func (s *currencyService) ListCurrency(ctx context.Context) ([]models.Currency, error) {
	query := `
		SELECT
			id
			, name
			, created_at
			, updated_at
		FROM currency`

	var currency []models.Currency
	if err := s.db.SelectContext(ctx, &currency, query); err != nil {
		return nil, fmt.Errorf("get the list of currency: %s", err)
	}

	return currency, nil
}

func (s *currencyService) GetCurrency(ctx context.Context, id int64) (models.Currency, error) {
	query := `
		SELECT
			id
			, name
			, created_at
			, updated_at
		FROM currency
		WHERE id = $1`

	var currency models.Currency
	if err := s.db.GetContext(ctx, &currency, query, id); err != nil {
		return models.Currency{}, fmt.Errorf("get an currency: %s", err)
	}

	return currency, nil
}

func (s *currencyService) CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error) {
	query := "INSERT INTO currency (name) VALUES ($1) RETURNING id"

	var id int64
	if err := s.db.QueryRowxContext(ctx, query, currency.Name).Scan(&id); err != nil {
		return models.Currency{}, fmt.Errorf("insert new currency: %s", err)
	}

	newCurrency, err := s.GetCurrency(ctx, id)
	if err != nil {
		return models.Currency{}, fmt.Errorf("get new currency: %s", err)
	}

	return newCurrency, nil
}

func (s *currencyService) DeleteCurrency(ctx context.Context, id int64) error {
	query := `DELETE FROM currency WHERE id = $1`

	if _, err := s.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete an currency: %s", err)
	}

	return nil
}

func (s *currencyService) UpdateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error) {
	query := `UPDATE currency SET name=$1 Where id=$2`

	if _, err := s.db.ExecContext(ctx, query, currency.Name, currency.ID); err != nil {
		return models.Currency{}, fmt.Errorf("Update currency: %s", err)
	}

	Currency, err := s.GetCurrency(ctx, currency.ID)
	if err != nil {
		return models.Currency{}, fmt.Errorf("get currency: %s", err)
	}

	return Currency, nil
}
