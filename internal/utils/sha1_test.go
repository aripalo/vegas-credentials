package utils

import "testing"

func TestGenerateSHA1(t *testing.T) {
	input := "foobar"

	// want generated with https://passwordsgenerator.net/sha1-hash-generator/
	want := "8843d7f92416211de9ebb963ff4ce28125932878"

	output, err := SHA1(input)
	if err != nil {
		t.Fatalf("Got error %q, want nil", err)
	}
	if output != want {
		t.Fatalf(`generateSha1Hash("%s") = %q, want match for %#q`, input, output, want)
	}
}
