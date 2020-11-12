package models

// Conversion model.
type Conversion struct {
	CreatedUpdated

	ID             int64 `json:"id" db:"id"  swaggerignore:"true"`
	ConversionFrom int64 `json:"conversion_from" db:"conversion_from" valid:"required"`
	ConversionTo   int64 `json:"conversion_to" db:"conversion_to" valid:"required"`
	Rate           int64 `json:"rate" db:"rate" valid:"required"`
}
