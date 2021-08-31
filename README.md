# 📮 Go Email Sender

[![Go Reference](https://pkg.go.dev/badge/github.com/koddr/go-email-sender.svg)](https://pkg.go.dev/github.com/koddr/go-email-sender)

Simple (but useful) email sender written in pure Go `v1.17`. Yes, _yet another_ email package here! 😅

Support **HTML templates** and **attachments**.

## Send HTML email

Method signature:

```go
func (s *Sender) SendHTMLEmail(
    templatePath string,
    dest []string,
    subject string,
    data interface{},
) error
```

Example:

```go
// Create a new struct for the email data.
type HTMLEmailData struct {
    Name    string
    Website string
}

// Create a new SMTP sender instance with your auth params.
sender := NewEmailSender("mail@test.com", "secret", "smtp.test.com", 25)

// Send the email with HTML template.
if err := sender.SendHTMLEmail(
    "my/templates/welcome.html",  // path to the HTML template
    []string{"mail@example.com"}, // slice of the emails to send
    "It's a test email!",         // subject of the email
    &HTMLEmailData{
        Name:    "Vic",
        Website: "https://shostak.dev/",
    },
); err != nil {
    // Throw error message, if something went wrong.
    return fmt.Errorf("Something went wrong: %v", err)
}
```

## Send plain text email

Method signature:

```go
func (s *Sender) SendPlainEmail(
    dest []string,
    subject string,
    data string,
) error
```

Example:

```go
// Create a new SMTP sender instance with your auth params.
sender := NewEmailSender("mail@test.com", "secret", "smtp.test.com", 25)

// Send the email with a plain text.
if err := sender.SendPlainEmail(
    "my/templates/welcome.html",  // path to the HTML template
    []string{"mail@example.com"}, // slice of the emails to send
    "It's a test email!",         // subject of the email
    "Here is a plain text body.", // body of the email
); err != nil {
    // Throw error message, if something went wrong.
    return fmt.Errorf("Something went wrong: %v", err)
}
```

## ⚠️ License

Apache-2.0 © [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
