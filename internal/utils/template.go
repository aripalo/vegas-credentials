package utils

import (
	"io"
	"text/template"
)

// Print text template to stdout.
func PrintTemplate[T any](dest io.Writer, name string, body string, data T) error {
	var tmpl *template.Template
	var err error

	tmpl, err = template.New(name).Parse(body)

	if err != nil {
		return err
	}

	err = tmpl.Execute(dest, data)
	if err != nil {
		return err
	}

	return nil
}
