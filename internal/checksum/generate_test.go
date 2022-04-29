package checksum

import "testing"

func TestGenerate(t *testing.T) {
	input := "foobar"

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	want := "5f6f3065208dde5f4624d7dfafc36a296a526590"

	output, err := Generate(input)
	if err != nil {
		t.Fatalf("Got error %q, want nil", err)
	}
	if output != want {
		t.Fatalf(`generateSha1Hash("%s") = %q, want match for %#q`, input, output, want)
	}
}
