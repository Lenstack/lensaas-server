package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"time"
)

type IEmail interface {
	Add(email Email) (err error)
	Send() (err error)
	Worker()
	Stop()
}

type Email struct {
	From        string   // Sender
	To          []string // Recipients
	Cc          []string // Copy to other recipients
	Subject     string   // Email subject
	Body        string   // HTML body template
	Queue       []Email  // Queue of emails to be sent
	Attachments []string // File attachments
	log         *zap.Logger
}

func NewEmail(log *zap.Logger) *Email {
	// Initialize email queue to be sent later
	log.Info("Initializing email queue...")
	return &Email{Queue: make([]Email, 0), log: log, From: viper.GetString("EMAIL_USER")}
}

func (e *Email) Add(email Email) (err error) {
	// Start worker to send email
	if len(e.Queue) == 0 {
		e.Worker()
	}
	// Add email to queue to be sent later
	e.log.Info("Adding email to queue...")
	e.Queue = append(e.Queue, email)
	return
}

func (e *Email) Send() (err error) {
	// Send all emails in queue
	for _, email := range e.Queue {
		e.log.Info("Sending email...")
		m := gomail.NewMessage()
		m.SetHeader("From", email.From)
		m.SetHeader("To", email.To...)
		m.SetHeader("Cc", email.Cc...)
		m.SetHeader("Subject", email.Subject)
		m.SetBody("text/html", email.Body)

		// Add attachment if any exist
		if len(email.Attachments) > 0 {
			for _, attachment := range email.Attachments {
				m.Attach(attachment)
			}
		}

		dialer := gomail.NewDialer(viper.GetString("EMAIL_HOST"), viper.GetInt("EMAIL_PORT"),
			viper.GetString("EMAIL_USER"), viper.GetString("EMAIL_PASSWORD"))
		if err := dialer.DialAndSend(m); err != nil {
			e.log.Error("Error sending email: ", zap.Error(err))
			return err
		}

		// Remove email has been sent from queue
		if len(e.Queue) > 0 {
			e.Queue = e.Queue[1:]
		}
	}
	e.log.Info("Email has been sent!")
	return
}

func (e *Email) Worker() {
	e.log.Info("Starting email worker...")
	// Start worker to send email
	go func() {
		for {
			// Wait for 30 seconds
			e.log.Info("Waiting for 30 seconds...")
			time.Sleep(30 * time.Second)

			if len(e.Queue) == 0 {
				e.log.Info("Email worker has been finished. Worker will stop.")
				e.Stop() // Stop worker
				return
			}

			if len(e.Queue) > 0 {
				e.log.Info("Email worker is sending email...")
				// Send email in queue
				if err := e.Send(); err != nil {
					fmt.Println("Error sending email: ", err)
				}
				e.log.Info("Email worker has been finished. Worker will stop.")
				e.Stop() // Stop worker
				return
			}
		}
	}()
}

func (e *Email) Stop() {
	// Stop sending email
	e.Queue = make([]Email, 0)
	e.log.Info("Email worker has been stopped.")
	return
}
