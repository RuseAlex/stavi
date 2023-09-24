package mail

import (
	"fmt"
	"net/smtp"
)

func SendEmail(options map[string]string, msg string, to []string) error {

	// Sender data
	from, exists := options["from"]
	if !exists {
		return fmt.Errorf("can't send email; invalid options %v", options)
	}
	password, exists := options["password"]
	if !exists {
		return fmt.Errorf("can't send email; invalid options %v", options)
	}

	// smtp server config
	stmpHost, exists := options["stmpHost"]
	if !exists {
		return fmt.Errorf("can't send email; invalid options %v", options)
	}
	stmpPort, exists := options["stmpPort"]
	if !exists {
		return fmt.Errorf("can't send email; invalid options %v", options)
	}

	// message
	message := []byte(msg)

	// authentication
	auth := smtp.PlainAuth("", from, password, stmpHost)

	// sending email
	err := smtp.SendMail(stmpHost+":"+stmpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
