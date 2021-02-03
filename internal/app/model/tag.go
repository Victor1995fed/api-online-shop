package model

import (
	"errors"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

//Tag ...
type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//GetTableName ...
func (p *Tag) GetTableName() string {
	return "tag"
}

//GetSupportedScenarioValidation ...
func (p *Tag) GetSupportedScenarioValidation() map[string]string {
	return map[string]string{
		"CREATE": "create",
		"UPDATE": "update",
		"DELETE": "delete",
	}
}

// Validate ...
func (p *Tag) Validate(scenario string) error {

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
