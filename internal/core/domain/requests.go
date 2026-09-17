package core_domain

import (
	"fmt"
)

type Request struct {
	ID     int64
	Prompt string
	Image  string
	UserID int64
}

func NewRequest(
	id int64,
	prompt string,
	image string,
	userID int64,
) Request {
	return Request{
		ID:     id,
		Prompt: prompt,
		Image:  image,
		UserID: userID,
	}
}

func NewRequestUninitialized(
	prompt string,
	userID int64,
) Request {
	return NewRequest(
		UnInitializedID,
		prompt,
		UnInitializedImage,
		userID,
	)
}

func (r *Request) Validate() error {
	promptLength := len([]rune(r.Prompt))
	if promptLength >= 1000 {
		return fmt.Errorf("`Prompt` must be less then 1000 characters")
	}

	return nil
}
