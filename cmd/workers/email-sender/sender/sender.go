package email_sender

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/wneessen/go-mail"
)

type VerificationEmail struct {
	Email      string `json:"email"`
	ConfirmURL string `json:"confirm_url"`
}

type VerificationEmailData struct {
	ConfirmURL string
}

func (s *Sender) SendMessage(ctx context.Context, body []byte) error {
	var data VerificationEmail

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("unmarshal email message: %w", err)
	}

	tmpl, err := template.ParseFiles(
		"cmd/workers/email-sender/public/index.html",
	)
	if err != nil {
		return fmt.Errorf("parse email template: %w", err)
	}

	var htmlBody bytes.Buffer

	if err := tmpl.Execute(
		&htmlBody,
		VerificationEmailData{
			ConfirmURL: data.ConfirmURL,
		},
	); err != nil {
		return fmt.Errorf("execute email template: %w", err)
	}

	msg := mail.NewMsg()

	if err := msg.From(s.from); err != nil {
		return fmt.Errorf("set sender: %w", err)
	}

	if err := msg.To(data.Email); err != nil {
		return fmt.Errorf("set recipient: %w", err)
	}

	msg.Subject("Подтверждение электронной почты")
	msg.SetBodyString(
		mail.TypeTextHTML,
		htmlBody.String(),
	)

	if err := s.client.DialAndSend(msg); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}
