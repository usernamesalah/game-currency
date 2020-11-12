package models

// Currency model.
type Currency struct {
	CreatedUpdated

	ID   int64  `json:"id" db:"id" swaggerignore:"true"`
	Name string `json:"name" db:"name" valid:"required"`
}
