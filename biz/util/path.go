package util

import (
	"path"
	"runtime"
)

//获取绝对文件路径
func GetAbsolutePath() string {
	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}
