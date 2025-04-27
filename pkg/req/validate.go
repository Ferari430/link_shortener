package req

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		log.Println("Validator bad")

		return err

	}
	return nil
}
