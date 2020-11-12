package services

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/usernamesalah/game-currency/pkg/models"
)

// ConversionsService service interface.
type ConversionsService interface {
	ListConversions(ctx context.Context) ([]models.Conversion, error)
	GetConversion(ctx context.Context, id int64) (models.Conversion, error)
	GetConversionBy(ctx context.Context, from, to int64) (models.Conversion, error)
	CreateConversion(ctx context.Context, conversion models.Conversion) (models.Conversion, error)
	DeleteConversion(ctx context.Context, id int64) error
	UpdateConversion(ctx context.Context, conversion models.Conversion) (models.Conversion, error)
}

type conversionsService struct {
	db *sqlx.DB
}

// NewConversionsService returns an initialized ConversionsService implementation.
func NewConversionsService(db *sqlx.DB) ConversionsService {
	return &conversionsService{db: db}
}

func (s *conversionsService) ListConversions(ctx context.Context) ([]models.Conversion, error) {
	query := `
		SELECT
			id
			, conversion_from
			, conversion_to
			, rate
			, created_at
			, updated_at
		FROM conversions`

	var conversions []models.Conversion
	if err := s.db.SelectContext(ctx, &conversions, query); err != nil {
		return nil, fmt.Errorf("get the list of conversions: %s", err)
	}

	return conversions, nil
}

func (s *conversionsService) GetConversion(ctx context.Context, id int64) (models.Conversion, error) {
	query := `
		SELECT
			id
			, conversion_from
			, conversion_to
			, rate
			, created_at
			, updated_at
		FROM conversions
		WHERE id = $1`

	var conversion models.Conversion
	if err := s.db.GetContext(ctx, &conversion, query, id); err != nil {
		return models.Conversion{}, fmt.Errorf("get an conversion: %s", err)
	}

	return conversion, nil
}

func (s *conversionsService) GetConversionBy(ctx context.Context, from, to int64) (models.Conversion, error) {
	query := `
		SELECT
			id
			, conversion_from
			, conversion_to
			, rate
			, created_at
			, updated_at
		FROM conversions
		WHERE (conversion_from = $1 AND conversion_to = $2) OR (conversion_to = $3 AND conversion_from = $4)`

	var conversion models.Conversion
	if err := s.db.GetContext(ctx, &conversion, query, from, to, to, from); err != nil {
		return models.Conversion{}, fmt.Errorf("get an conversion: %s", err)
	}

	return conversion, nil
}

func (s *conversionsService) CreateConversion(ctx context.Context, conversion models.Conversion) (models.Conversion, error) {

	data, err := s.GetConversionBy(ctx, conversion.ConversionFrom, conversion.ConversionTo)
	if err != nil || data.ID == 0 {
		return models.Conversion{}, fmt.Errorf("Convesion Already Exist")
	}

	query := "INSERT INTO conversions (conversion_from, conversion_to, rate) VALUES ($1, $2 , $3) RETURNING id"

	var id int64
	if err := s.db.QueryRowxContext(ctx, query, conversion.ConversionFrom, conversion.ConversionTo, conversion.Rate).Scan(&id); err != nil {
		return models.Conversion{}, fmt.Errorf("insert new conversion: %s", err)
	}

	newConversion, err := s.GetConversion(ctx, id)
	if err != nil {
		return models.Conversion{}, fmt.Errorf("get new conversion: %s", err)
	}

	return newConversion, nil
}

func (s *conversionsService) DeleteConversion(ctx context.Context, id int64) error {
	query := `DELETE FROM conversions WHERE id = $1`

	if _, err := s.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("delete an conversion: %s", err)
	}

	return nil
}

func (s *conversionsService) UpdateConversion(ctx context.Context, conversion models.Conversion) (models.Conversion, error) {
	query := `UPDATE conversions SET 
	conversion_from = $1
	, conversion_to = $2
	, rate = $3
	Where id=$4`

	if _, err := s.db.ExecContext(ctx, query, conversion.ConversionFrom, conversion.ConversionTo, conversion.Rate, conversion.ID); err != nil {
		return models.Conversion{}, fmt.Errorf("Update conversion: %s", err)
	}

	Conversion, err := s.GetConversion(ctx, conversion.ID)
	if err != nil {
		return models.Conversion{}, fmt.Errorf("get conversion: %s", err)
	}

	return Conversion, nil
}
