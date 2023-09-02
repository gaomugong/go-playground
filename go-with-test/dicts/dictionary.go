package dicts

import (
	"errors"
)

type Dict map[string]string

func (d Dict) Search(s string) string {
	return d[s]
}

var (
	DictKeyNotFound = errors.New("could not find the word you were looking for")
	DictKeyExist    = errors.New("could not find the word you were looking for")
)

func (d Dict) Find(s string) (string, error) {
	if value, ok := d[s]; ok {
		return value, nil
	}
	return "", DictKeyNotFound
}

func (d Dict) Add(key string, value string) {
	//Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。
	//这意味着它拥有对底层数据结构的引用，就像指针一样。
	d[key] = value
}

func (d Dict) AddErr(word string, defination string) error {

	//if _, exist := d[word]; exist {
	//	return DictKeyExist
	//}
	_, err := d.Find(word)

	switch err {
	case DictKeyNotFound:
		d[word] = defination
	case nil:
		return DictKeyExist
	default:
		return err
	}
	return nil
}
