package tmpl

import (
	"io"
	"text/template"
)

// Write text template to stdout. An utility method that wraps text/template
// with our own preferences.
func Write[T any](dest io.Writer, name string, body string, data T) error {
	var tmpl *template.Template
	var err error

	tmpl, err = template.New(name).Option("missingkey=error").Parse(body)
	if err != nil {
		return err
	}

	err = tmpl.Execute(dest, data)
	if err != nil {
		return err
	}

	return nil
}
