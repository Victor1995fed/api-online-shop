package model

import (
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Order ...
type Order struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string `json:""`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Comment     string `json:"comment"`
	Phone		string `json:"phone"`
}

//GetTableName ...
func (p *Order) GetTableName() string {
	return "product"
}

//GetSupportedScenarioValidation ...
func (p *Order) GetSupportedScenarioValidation() map[string]string {
	return map[string]string{
		"CREATE": "create",
		"UPDATE": "update",
		"DELETE": "delete",
	}
}

// Validate ...
func (p *Order) Validate(scenario string) error {

	supportScenario := p.GetSupportedScenarioValidation()
	var result error
	switch scenario {
	case supportScenario["CREATE"]:
		result = validation.ValidateStruct(p,
			validation.Field(
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