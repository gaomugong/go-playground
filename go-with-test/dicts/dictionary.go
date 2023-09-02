package dicts

import (
	"errors"
)

type Dict map[string]string

func (d Dict) Search(s string) string {
	return d[s]
}

var DictKeyError = errors.New("could not find the word you were looking for")

func (d Dict) Find(s string) (string, error) {
	if value, ok := d[s]; ok {
		return value, nil
	}
	return "", DictKeyError
}
