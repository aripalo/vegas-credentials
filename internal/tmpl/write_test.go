package tmpl

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

type tmplData struct {
	Version string
	DirPath string
}

type incorrectTmplData struct{}

func TestWrite(t *testing.T) {
	tests := []struct {
		name     string
		template string
		input    any
		expected string
		err      error
	}{
		{
			name:     "no data",
			template: "{{.Version}} @ {{.DirPath}}",
			input:    incorrectTmplData{},
			expected: "",
			err: template.ExecError{
				Name: "test",
				Err:  errors.New(`template: test:1:2: executing "test" at <.Version>: can't evaluate field Version in type tmpl.incorrectTmplData`),
			},
		},
		{
			name:     "incorrect template",
			template: "{{{}.Version}} @ {{.DirPath}}",
			input: tmplData{
				Version: "v2.47.113",
				DirPath: "/tmp/foo/bar",
			},
			expected: "",
			err:      errors.New(`template: test:1: unexpected "{" in command`),
		},
		{
			name:     "success",
			template: "{{.Version}} @ {{.DirPath}}",
			input: tmplData{
				Version: "v2.47.113",
				DirPath: "/tmp/foo/bar",
			},
			expected: "v2.47.113 @ /tmp/foo/bar",
			err:      nil,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer
			err := Write(&output, "test", test.template, test.input)
			actual := output.String()
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
