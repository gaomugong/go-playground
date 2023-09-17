package server

import (
	"os"
)

// tape 创建一个新类型来封装我们的「当写入时，从头部开始」功能
type tape struct {
	//file io.ReadWriteSeeker
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
