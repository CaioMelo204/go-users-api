package models

import (
	"devbook-api/pkg/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
)

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) Prepare(stage string) error {
	if err := u.validate(); err != nil {
		return err
	}
	err := u.format(stage)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Nickname == "" {
		return errors.New("nickname is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.Email = strings.TrimSpace(u.Email)

	if stage == "create" {
		hashPass, err := security.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashPass)
	}

	return nil
}
