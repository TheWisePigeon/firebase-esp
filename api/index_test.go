package handler

import "testing"

func TestHandler(t *testing.T) {
	sender := NewSender("hamanecisse@esp.sn", "cgdseyhhoeujfyzu")
	receiver := []string{"hamanecisse2@gmail.com"}

	subject := "Signal from Arduino"
	message := `Hello`
	bodyMessage := sender.WritePlainEmail(receiver, subject, message)

	sender.SendMail(receiver, subject, bodyMessage)
}
