package database

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/dgraph-io/badger/v3"
)

// createTempDir creates a temporary directory for testing the disk database
// and returns the path for the directory and a function to remove the dir after test
func createTempDir() (string, func()) {
	dir, err := ioutil.TempDir("", strings.Join([]string{config.APP_NAME, "test", ""}, "-"))
	if err != nil {
		panic(err)
	}
	return dir, func() { os.RemoveAll(dir) }
}

func TestDatabaseOpen(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})

	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	// Use the Close method on the actual badger instance
	err = database.db.Close()
	if err != nil {
		t.Fatalf("Could not close the database: %q", err)
	}
}

func TestDatabaseClose(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})

	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	// use the Close method on database struct
	err = database.Close()
	if err != nil {
		t.Fatalf("Could not close the database: %q", err)
	}
}

func TestDatabaseWrite(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Microsecond)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}
}

func TestDatabaseWriteAndRead(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}

	output, err := database.Read(key)
	if err != nil {
		t.Fatalf("Could not read from database: %q", err)
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

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}

	err = database.Delete(key)
	if err != nil {
		t.Fatalf("Could not delete from database: %q", err)
	}

	_, err = database.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseDeleteByPrefix(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}

	err = database.DeleteByPrefix("f")
	if err != nil {
		t.Fatalf("Could not delete all from database: %q", err)
	}

	_, err = database.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseDeleteAll(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Minute)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}

	err = database.DeleteAll()
	if err != nil {
		t.Fatalf("Could not delete all from database: %q", err)
	}

	_, err = database.Read(key)
	if err.Error() != badger.ErrKeyNotFound.Error() {
		t.Fatalf("Invalid response: %q", err)
	}
}

func TestDatabaseTTL(t *testing.T) {
	tempDir, tempDirRemove := createTempDir()
	defer tempDirRemove()

	database, err := Open(tempDir, DatabaseOptions{})
	defer func() {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatalf("Could not open the database: %q", err)
	}

	key := "foo"
	data := []byte(randStringBytes(1000))

	err = database.Write(key, data, time.Second)
	if err != nil {
		t.Fatalf("Could not write to database: %q", err)
	}

	output, err := database.Read(key)
	if err != nil {
		t.Fatalf("Could not read from database: %q", err)
	}

	got := string(output)
	want := string(data)

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}

	time.Sleep(time.Second)

	_, err = database.Read(key)
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
