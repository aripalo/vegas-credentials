package diskcache

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgraph-io/badger/v3"
)

// createTempDir creates a temporary directory for testing the disk d
// and returns the path for the directory and a function to remove the dir after test
func createTempDir() (string, func()) {
	dir, err := ioutil.TempDir("", strings.Join([]string{"vegas-credentials", "test", ""}, "-"))
	if err != nil {
		panic(err)
	}
	return dir, func() { os.RemoveAll(dir) }
}

func TestDatabaseNew(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})

	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	// Use the Close method on the actual badger instance
	err = d.db.Close()
	if err != nil {
		t.Fatalf("Could not close the d: %q", err)
	}
}

func TestDatabaseClose(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})

	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	// use the Close method on d struct
	err = d.Close()
	if err != nil {
		t.Fatalf("Could not close the d: %q", err)
	}
}

func TestDatabaseWrite(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Microsecond)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}
}

func TestDatabaseWriteAndRead(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}

	output, err := d.Read(key)
	if err != nil {
		t.Fatalf("Could not read from d: %q", err)
	}

	got := string(output)
	want := string(data)

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDatabaseDelete(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}

	err = d.Delete(key)
	if err != nil {
		t.Fatalf("Could not delete from d: %q", err)
	}

	_, err = d.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseDeleteByPrefix(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}

	err = d.DeleteByPrefix("f")
	if err != nil {
		t.Fatalf("Could not delete all from d: %q", err)
	}

	_, err = d.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseDeleteAll(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}

	err = d.DeleteAll()
	if err != nil {
		t.Fatalf("Could not delete all from d: %q", err)
	}

	_, err = d.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseTTL(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	d, err := New(tempDir, Options{})
	defer func() {
		err := d.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the d: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = d.Write(key, data, time.Second)
	if err != nil {
		t.Fatalf("Could not write to d: %q", err)
	}

	output, err := d.Read(key)
	if err != nil {
		t.Fatalf("Could not read from d: %q", err)
	}

	got := string(output)
	want := string(data)

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}

	time.Sleep(time.Second)

	_, err = d.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
