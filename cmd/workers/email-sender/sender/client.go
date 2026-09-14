package email_sender

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

type Sender struct {
	client *mail.Client
	from   string
}

func NewSender(config Config) (*Sender, error) {
	client, err := mail.NewClient(
		config.Host,
		mail.WithPort(config.Port),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(config.User),
		mail.WithPassword(config.Password),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize sender client: %w", err)
	}

	return &Sender{
		client: client,
		from:   config.From,
	}, nil
}
