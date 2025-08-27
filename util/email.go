package util

import (
	"github.com/happy3014/happybase/config"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(to []string, subject string, body string, conf config.EmailSenderConfig) error {
	e := email.NewEmail()
	e.From = conf.From
	e.To = to
	e.Subject = subject
	e.Text = []byte(body)
	err := e.Send(conf.SmtpAuth.Host, smtp.PlainAuth(conf.SmtpAuth.Identity, conf.SmtpAuth.Username, conf.SmtpAuth.Password, conf.SmtpAuth.Host))
	return err
}

func SendEmailWithDefaultConfig(subject string, body string) error {
	conf := config.GlobalConfig().EmailSender
	return SendEmail(conf.DefaultTo, subject, body, conf)
}
