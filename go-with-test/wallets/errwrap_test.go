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
		// 还原被包装后的错误，得到错误的"根源"
		originalError := errors.Unwrap(err)
		fmt.Println(originalError)
		fmt.Println(reflect.TypeOf(err))
		fmt.Println(reflect.TypeOf(originalError))
		t.Errorf("got error: %v", err)
	}
}

// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand.
// If there is more than one %w verb, the returned error will implement
// an Unwrap method returning a []error containing all the %w operands
// in the order they appear in the arguments.
func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := os.ReadFile(filepath.Join(home, ".settings.xml"))
	//https://rollbar.com/blog/golang-wrap-and-unwrap-error/
	// 利用%w包装错误 *fs.PathError -> *fmt.wrapError
	return config, fmt.Errorf("read config failed: %w", err)
	//return config, fmt.Errorf("read config failed: %w, %w", err, errors.New("ANOTHER ERROR"))
}
