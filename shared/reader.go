package shared

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func ReadFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
