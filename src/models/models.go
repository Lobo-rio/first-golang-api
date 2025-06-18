package models

import (
	"errors"
	"modules/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

// Prepare vai chamar os métodos para validar e formatar o usuário recebido
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if user.NickName == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if step == "registration" && user.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "registration" {
		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hash)
	}

	return nil
}