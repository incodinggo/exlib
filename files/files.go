package files

import (
	"os"
)

// CheckFileIsExist 判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// FileHash 文件夹散列   6512bd43d9caa6e02c990b0a82652dca => 6/5/1/2/
// split 需要几位
func FileHash(hash string, split int) (path string) {
	if split < 0 || split > 32 {
		return ""
	}
	rs := []rune(hash)
	for i := 0; i < split; i++ {
		end := i + 1
		path += string(rs[i:end]) + string(os.PathSeparator)
	}
	return
}
