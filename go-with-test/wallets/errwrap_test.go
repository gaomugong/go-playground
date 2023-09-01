package wallets

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestErrorWrap(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		//fmt.Println(err)
		t.Errorf("got error: %v", err)
	}
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := os.ReadFile(filepath.Join(home, ".settings.xml"))
	//https://rollbar.com/blog/golang-wrap-and-unwrap-error/
	return config, fmt.Errorf("read config failed: %w", err)
}
