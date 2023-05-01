package utils

import "testing"

func TestSendEmail(t *testing.T) {
	email := NewEmail()
	err := email.Add(Email{
		From:    "asesin-blood@gmail.com",
		To:      []string{"asesin-blood@gmail.com"},
		Subject: "Test",
		Body:    "Test",
	})

	err = email.Add(Email{
		From:    "asesin-blood@gmail.com",
		To:      []string{"asesin-blood@gmail.com"},
		Subject: "Test2",
		Body:    "Test2",
	})

	err = email.Add(Email{
		From:    "asesin-blood@gmail.com",
		To:      []string{"asesin-blood@gmail.com"},
		Subject: "Test3",
		Body:    "Test3",
	})

	if len(email.Queue) == 0 {
		t.Errorf("Email queue is empty")
	}

	if err != nil {
		t.Errorf("Error sending email: %v", err)
	}

	for len(email.Queue) > 0 {
		t.Log("Email sent! " + email.Queue[0].Subject)
		email.Queue = email.Queue[1:]
	}
}
