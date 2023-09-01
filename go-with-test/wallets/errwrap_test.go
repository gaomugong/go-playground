package wallets

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestErrorWrap(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		originalError := errors.Unwrap(err)
		fmt.Println(originalError)
		fmt.Println(reflect.TypeOf(originalError))
		t.Errorf("got error: %v", err)
	}
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := os.ReadFile(filepath.Join(home, ".settings.xml"))
	//https://rollbar.com/blog/golang-wrap-and-unwrap-error/
	return config, fmt.Errorf("read config failed: %w", err)
}
