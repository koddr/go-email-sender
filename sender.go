package sender

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
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
func (s *Sender) SendHTMLEmail(templatePath string, to, cc []string, subject string, data interface{}, files []string) error {
	if templatePath == "" {
		return fmt.Errorf("Template not found in the given path!")
	}
	tmpl, err := ParseTemplate(templatePath, data)
	if err != nil {
		return err
	}
	body := s.writeEmail(to, cc, "text/html", subject, tmpl, files)
	if err := s.sendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}

// SendPlainEmail func for send plain text email with data.
func (s *Sender) SendPlainEmail(to, cc []string, subject, data string, files []string) error {
	body := s.writeEmail(to, cc, "text/plain", subject, data, files)
	if err := s.sendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}

func (s *Sender) attachFile(file string) string {
	rawFile, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(rawFile)
}

// writeEmail method for prepare email header, body and files to send.
func (s *Sender) writeEmail(to, cc []string, ct, subj, body string, files []string) string {
	// Define variables.
	var message string
	var encodedBody bytes.Buffer

	// Create delimiter.
	delimiter := fmt.Sprintf("**=mail%d", time.Now().UnixNano())

	// Create writer for make encoding the email's body.
	result := quotedprintable.NewWriter(&encodedBody)
	if _, err := result.Write([]byte(body)); err != nil {
		return ""
	}
	defer result.Close()

	// Create message.
	message += fmt.Sprintf("From: %s\r\n", s.Login)
	message += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))
	if len(cc) > 0 {
		// If CC is specified.
		message += fmt.Sprintf("Cc: %s\r\n", strings.Join(cc, ";"))
	}
	message += fmt.Sprintf("Subject: %s\r\n", subj)
	message += fmt.Sprintf("MIME-Version: 1.0\r\n")
	message += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimiter)

	// Add body of the message (from html template or plain text).
	message += fmt.Sprintf("--%s\r\n", delimiter)
	message += fmt.Sprintf("Content-Transfer-Encoding: quoted-printable\r\n")
	message += fmt.Sprintf("Content-Type: %s; charset=\"utf-8\"\r\n", ct)
	message += fmt.Sprintf("Content-Disposition: inline\r\n")
	message += fmt.Sprintf("%s\r\n", encodedBody.String())

	// If files are specified.
	if len(files) > 0 {
		// Add body of the each file to the message.
		for _, file := range files {
			message += fmt.Sprintf("--%s\r\n", delimiter)
			message += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n")
			message += fmt.Sprintf("Content-Type: text/plain; charset=\"utf-8\"\r\n")
			message += fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", file)
			message += fmt.Sprintf("\r\n%s", s.attachFile(file))
		}
	}

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
