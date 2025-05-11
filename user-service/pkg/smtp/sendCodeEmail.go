package smtp

import (
	"fmt"
	"net/smtp"
)

func SendVerificationCode(toEmail, code string) error {
	from := "egore.chip@gmail.com"
	password := "ssqiavutgcrjrcjm"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", code)
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	return err
}
