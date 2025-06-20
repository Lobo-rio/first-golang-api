package models

import (
	"errors"
	"fmt"
	"modules/src/security"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required,min=3,max=50"`
	NickName  string    `json:"nickname,omitempty" validate:"required,min=3,max=30"`
	Email     string    `json:"email,omitempty" validate:"required,email,min=8,max=100"`
	Phone     string    `json:"phone,omitempty" validate:"required,min=8,max=15"`
	Password  string    `json:"password,omitempty" validate:"required,min=6,max=100"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}


var validateUser *validator.Validate

// Prepare function used to call the methods to validate and format the received user!
func (user *User) Prepare(step string) error {
	validateUser = validator.New(validator.WithRequiredStructEnabled())

	if err := user.validateStruct(); err != nil {
		return err
	}

	if err := user.validateVariable(); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validateStruct() error {

	err := validateUser.Struct(user)
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

func (user *User) validateVariable() error {
    err := validateUser.Var(user.Email, "required,email")

	if err != nil {
		fmt.Println(err) 
		return err
	}
	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = strings.TrimSpace(user.Phone)

	if etapa == "registration" {
		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hash)
	}

	return nil
}