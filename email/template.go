package email

import (
	"bytes"
	"html/template"
)

func LoadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}

func RenderTemplate(tmpl *template.Template, fullName string) (string, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, struct {
		FullName        string
		UnsubscribeLink string
	}{
		FullName:        fullName,
		UnsubscribeLink: "https://smart-stack.ru/unsubscribe", // 👈 сюда можно вставлять токенизированный URL
	})
	return buf.String(), err
}
