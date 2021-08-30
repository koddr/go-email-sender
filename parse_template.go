package sender

import (
	"bytes"
	"html/template"
)

// ParseTemplate ...
func ParseTemplate(file string, data interface{}) (string, error) {
	tmpl, errParseFiles := template.ParseFiles(file)
	if errParseFiles != nil {
		return "", errParseFiles
	}
	buffer := new(bytes.Buffer)
	if errExecute := tmpl.Execute(buffer, data); errExecute != nil {
		return "", errExecute
	}
	return buffer.String(), nil
}
