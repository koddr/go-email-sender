# 📮 Go Email Sender

[![Go Reference](https://pkg.go.dev/badge/github.com/koddr/go-email-sender.svg)](https://pkg.go.dev/github.com/koddr/go-email-sender)

Simple (but useful) email sender written in pure Go `v1.17`. Yes, _yet another_ email package here! 😅

Support **HTML templates** and **attachments**.

## Send HTML email

Method signature:

```go
func (s *Sender) SendHTMLEmail(
    templatePath string,
    to []string,
    cc []string,
    subject string,
    data interface{},
    files []string,
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

// Send the email with an HTML template.
if err := sender.SendHTMLEmail(
    "my/templates/welcome.html", // path to the HTML template
    []string{
        "mail@example.com"       // slice of the emails to send
    },
    []string{
        "copy-mail@example.com"  // slice of the emails to send email copy
    },
    "It's a test email!",        // subject of the email
    &HTMLEmailData{
        Name:    "Vic",
        Website: "https://shostak.dev/",
    },
    []string{
        "my/files/image.jpg",    // slice of the files to send
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
    to []string,
    cc []string,
    subject string,
    data interface{},
    files []string,
) error
```

Example:

```go
// Create a new SMTP sender instance with your auth params.
sender := NewEmailSender("mail@test.com", "secret", "smtp.test.com", 25)

// Send the email with a plain text.
if err := sender.SendPlainEmail(
    []string{
        "mail@example.com"        // slice of the emails to send
    },
    []string{
        "copy-mail@example.com"   // slice of the emails to send email copy
    },
    "It's a test email!",         // subject of the email
    "Here is a plain text body.", // body of the email
    []string{
        "my/files/image.jpg",     // slice of the files to send
    },
); err != nil {
    // Throw error message, if something went wrong.
    return fmt.Errorf("Something went wrong: %v", err)
}
```

## ⚠️ License

Apache-2.0 © [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
