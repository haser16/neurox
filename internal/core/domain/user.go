package core_domain

import (
	"fmt"
	core_errors "neurox/internal/core/errors"
	"regexp"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Phone    *string
	Avatar   *string
	Password string
}

func NewUser(
	id int64,
	username string,
	email string,
	phone *string,
	password string,
) User {
	return User{
		ID:       id,
		Username: username,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
}

func NewUserUninitialized(
	username string,
	email string,
	phone *string,
	password string,
) User {
	return NewUser(
		UnInitializedID,
		username,
		email,
		phone,
		password)
}

func (u *User) Validate() error {
	usernameLength := len([]rune(u.Username))
	if usernameLength < 2 || usernameLength > 100 {
		return fmt.Errorf("`Username` must be between 2 and 100 characters")
	}

	if _, err := regexp.MatchString("@", u.Email); err != nil {
		return fmt.Errorf("invalid user `Email`: %w", err)
	}

	if u.Phone != nil {
		phoneNumberLength := len([]rune(*u.Phone))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalid `Phone` lenght: %d: %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.Phone) {
			return fmt.Errorf(
				"invalid `Phone` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
