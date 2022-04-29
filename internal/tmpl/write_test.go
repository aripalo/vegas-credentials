package tmpl

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tmplData struct {
	Version string
	DirPath string
}

func Test(t *testing.T) {

	name := "flag parsing"
	template := "{{.Version}} @ {{.DirPath}}"
	input := tmplData{
		Version: "v2.47.113",
		DirPath: "/tmp/foo/bar",
	}
	expected := "v2.47.113 @ /tmp/foo/bar"

	t.Run(name, func(t *testing.T) {
		var output bytes.Buffer
		err := Write(&output, "test", template, input)
		actual := string(output.Bytes())
		assert.Equal(t, err, nil)
		assert.Equal(t, expected, actual)
	})

}
