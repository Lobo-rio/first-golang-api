package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Note struct {
	ID             uint64 `json:"id,omitempty"`
	Title          string `json:"title,omitempty" validate:"required,min=3,max=50"`
	Content        string `json:"content,omitempty" validate:"required,min=10,max=300"`
	AuthorID       uint64 `json:"author_id,omitempty"`
	AuthorNickName string `json:"author_nickname,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

var validate *validator.Validate

// Prepare function used to call the methods to validate and format the received user!
func (note *Note) Prepare(step string) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	if err := note.validateStruct(); err != nil {
		return err
	}

	if err := note.format(step); err != nil {
		return err
	}

	return nil
}

func (note *Note) validateStruct() error {

	err := validate.Struct(note)
	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return err
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				fmt.Println(e.Namespace())
				fmt.Println(e.Field())
				fmt.Println(e.StructNamespace())
				fmt.Println(e.StructField())
				fmt.Println(e.Tag())
				fmt.Println(e.ActualTag())
				fmt.Println(e.Kind())
				fmt.Println(e.Type())
				fmt.Println(e.Value())
				fmt.Println(e.Param())
				fmt.Println()
			}
		}
		return err
	}

	return nil
}

func (note *Note) format(etapa string) error {
	note.Title = strings.TrimSpace(note.Title)
	note.Content = strings.TrimSpace(note.Content)

	return nil
}
