package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Product ...
type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageURL    string `json:""`
}

//GetTableName ...
func (p *Product) GetTableName() string {
	return "product"
}

// Validate ...
func (p *Product) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(
			&p.Title,
			validation.Required,
			is.Alphanumeric,
		),
		validation.Field(
			&p.Price,
			validation.Required,
			is.Float,
		),
		// validation.Field(
		// 	&p.Password,
		// 	validation.By(requiredIf(p.EncryptedPassword == "")),
		// 	validation.Length(6, 100),
		// ),
	)
}
