package dicts

import (
	"fmt"
	"testing"
)

const testStr = "this is just a test"

func TestSearch(t *testing.T) {
	dict := Dict{
		"test": testStr,
	}

	got := dict.Search("test")
	want := testStr

	assertString(t, got, want)
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v, want %#v, given %#v", got, want, "test")
	}
}

func assertError(t *testing.T, want, got error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %#v, want %#v", got, want)
	}
}

func TestFind(t *testing.T) {
	dict := Dict{"test": testStr}

	t.Run("test: found", func(t *testing.T) {
		got, _ := dict.Find("test")
		want := testStr
		assertString(t, got, want)
	})

	t.Run("unknown: not found", func(t *testing.T) {
		_, err := dict.Find("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, DictKeyNotFound)
	})

}

func TestAdd(t *testing.T) {
	dict := Dict{}
	key := "test"
	want := testStr

	dict.Add(key, testStr)

	got, err := dict.Find("test")
	if err != nil {
		t.Fatal("add new word failed -> ", err)
	}

	if want != got {
		t.Errorf("got #%v, want #%v", got, want)
	}
}

func TestMap(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from: ", err)
		}
	}()

	//正确的初始化方式
	m1 := map[string]string{}
	m2 := make(map[string]string)
	m1["a"] = "b"
	m2["a"] = "b"
	fmt.Printf("m1 = %#v, m2 = %#v\n", m1, m2)

	//引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，
	//你会得到一个 nil 指针异常，这将导致程序终止运行
	// panic: assignment to entry in nil map
	var m map[string]string
	m["a"] = "b"

}

// assertDefination 将定义断言移到了自己的辅助函数
func assertDefination(t *testing.T, dict Dict, word, defination string) {
	got, err := dict.Find(word)
	if err != nil {
		t.Fatal("add new word failed -> ", err)
	}

	if defination != got {
		t.Errorf("got %s, want %s", got, defination)
	}
}

func TestAddErr(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		word, defination := "test", testStr
		err := dict.AddErr(word, defination)

		assertError(t, err, nil)
		assertDefination(t, dict, word, defination)
	})

	t.Run("exist word", func(t *testing.T) {
		word, defination := "test", testStr

		dict := Dict{word: defination}

		err := dict.AddErr(word, defination)

		assertError(t, err, DictKeyExist)
		assertDefination(t, dict, word, defination)
	})
}
