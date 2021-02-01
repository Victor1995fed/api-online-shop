package model

import (
	"errors"
	"strconv"

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

//GetSupportedScenarioValidation ...
func (p *Product) GetSupportedScenarioValidation() map[string]string {
	return map[string]string{
		"CREATE": "create",
		"UPDATE": "update",
		"DELETE": "delete",
	}
}

// Validate ...
func (p *Product) Validate(scenario string) error {

	supportScenario := p.GetSupportedScenarioValidation()
	var result error
	switch scenario {
	case supportScenario["CREATE"]:
		result = validation.ValidateStruct(p,
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
		)
	case supportScenario["UPDATE"]:
		var id string
		if p.ID > 0 {
			id = strconv.Itoa(p.ID)
		}
		result = validation.Validate(id,
			validation.Required.Error("ID is required"),
			is.Int,
		)

	case supportScenario["DELETE"]:
		var id string
		if p.ID > 0 {
			id = strconv.Itoa(p.ID)
		}
		result = validation.Validate(id,
			validation.Required.Error("ID is required"),
			is.Int,
		)
	default:
		result = errors.New("unknown scenario")
	}
	return result
}
