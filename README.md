# 📮 Go Email Sender

[![Go Reference](https://pkg.go.dev/badge/github.com/koddr/go-email-sender.svg)](https://pkg.go.dev/github.com/koddr/go-email-sender)

Simple (but useful) email sender written in pure Go `v1.17`. Yes, _yet another_ email package here! 😅

Support **HTML templates** and **attachments**.

## Send HTML email

Method signature:

```go
func (s *Sender) SendHTMLEmail(templatePath string, dest []string, subject string, data interface{}) error
```

Example:

```go
// Create a new struct for the email data.
type HTMLEmailData struct {
    Name    string
    Website	string
}

// Create a new SMTP sender instance with your auth params.
sender := NewEmailSender("mail@test.com", "secret", "smtp.test.com", 25)

// Send the email with HTML template.
if err := sender.SendHTMLEmail(
    "my/templates/welcome.html",
    []string{"mail@example.com"},
    "It's a test email!",
    &HTMLEmailData{
        Name:    "Vic",
        Website: "https://shostak.dev/",
    },
); err != nil {
    return fmt.Errorf("Something went wrong: %v", err)
}
```

## Send plain text email

Method signature:

```go
func (s *Sender) SendPlainEmail(dest []string, subject, data string) error
```

Example:

```go
// Create a new struct for the email data.
type PlainEmailData struct {
    Company string
    Website	string
}

// Create a new SMTP sender instance with your auth params.
sender := NewEmailSender("mail@test.com", "secret", "smtp.test.com", 25)

// Send the email with HTML template.
if err := sender.SendHTMLEmail(
    "my/templates/welcome.html",
    []string{"mail@example.com"},
    "It's a test email!",
    &PlainEmailData{
        Company: "True web artisans",
        Website: "https://1wa.co/",
    },
); err != nil {
    return fmt.Errorf("Something went wrong: %v", err)
}
```

## ⚠️ License

Apache-2.0 © [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
