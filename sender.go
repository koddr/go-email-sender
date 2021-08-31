package sender

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

// Sender struct for describe auth data for sending emails.
type Sender struct {
	Login          string
	Password       string
	SMTPServer     string
	SMTPServerPort int
}

// NewEmailSender func for create new email sender.
// Includes login, password, SMTP server and port.
func NewEmailSender(login, password, server string, port int) *Sender {
	return &Sender{login, password, server, port}
}

// SendHTMLEmail func for send email with given HTML template and data.
// If template is empty string, it will throw error.
func (s *Sender) SendHTMLEmail(templatePath string, dest []string, subject string, data interface{}) error {
	if templatePath == "" {
		return fmt.Errorf("Template not found in the given path!")
	}
	tmpl, errParseTemplate := ParseTemplate(templatePath, data)
	if errParseTemplate != nil {
		return errParseTemplate
	}
	body := s.writeEmail(dest, "text/html", subject, tmpl)
	if errSendEmail := s.sendEmail(dest, subject, body); errSendEmail != nil {
		return errSendEmail
	}
	return nil
}

// SendPlainEmail func for send plain text email with data.
func (s *Sender) SendPlainEmail(dest []string, subject, data string) error {
	body := s.writeEmail(dest, "text/plain", subject, data)
	if errSendEmail := s.sendEmail(dest, subject, body); errSendEmail != nil {
		return errSendEmail
	}
	return nil
}

// writeEmail method for prepare email header and body to send.
func (s *Sender) writeEmail(dest []string, contentType, subject, body string) string {
	// Define variables.
	var message string
	var header map[string]string
	var encodedMessage bytes.Buffer

	// Create header.
	header["From"] = s.Login
	header["To"] = strings.Join(dest, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	// Add header values to the message.
	for key, value := range header {
		message += fmt.Sprintf("%s:%s\r\n", key, value)
	}

	// Create writer for make encoding the message.
	result := quotedprintable.NewWriter(&encodedMessage)
	result.Write([]byte(body))
	result.Close()

	// Return the encoded message string.
	message += "\r\n" + encodedMessage.String()
	return message
}

// sendEmail method for send email to destination addresses with subject and body.
func (s *Sender) sendEmail(dest []string, subject, body string) error {
	if err := smtp.SendMail(
		fmt.Sprintf("%s:%d", s.SMTPServer, s.SMTPServerPort),
		smtp.PlainAuth("", s.Login, s.Password, s.SMTPServer),
		s.Login, dest, []byte(body),
	); err != nil {
		return err
	}
	return nil
}
