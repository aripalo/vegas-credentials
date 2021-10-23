package response

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	r := New()

	defer func() {
		err := r.cache.Disconnect()
		if err != nil {
			panic(err)
		}
	}()

	wantedVersion := 1
	if r.Version != wantedVersion {
		t.Fatalf("Got Version %q, want %q", r.Version, wantedVersion)
	}

	if r.cache == nil {
		t.Fatalf("cache not defined")
	}

	if r.destination != os.Stdout {
		t.Fatal("destination should be /dev/stdout")
	}
}

func TestTeardown(t *testing.T) {
	r := New()

	err := r.Teardown()
	if err != nil {
		t.Fatalf("Got err %q, want nil", err)
	}

}
