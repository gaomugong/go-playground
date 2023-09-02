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

func (d Dict) Add(key string, value string) {
	//Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。
	//这意味着它拥有对底层数据结构的引用，就像指针一样。
	d[key] = value
}
