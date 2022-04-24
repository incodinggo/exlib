package stringex

import (
	"github.com/gogf/gf/v2/text/gstr"
	"regexp"
	"strconv"
	"strings"
)

// TrimHtml 去除字符串的HTML
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func Int64(s string) int64 {
	v, err := strconv.ParseInt(s, 0, 0)
	if err == nil {
		return v
	}
	return 0
}

func Int(s string) int {
	v, err := strconv.ParseInt(s, 0, 0)
	if err == nil {
		return int(v)
	}
	return 0
}

// Split 切分字符串
// needSpace是否保留空字符，默认false
func Split(s, sep string, needSpace ...bool) *T {
	var ns bool
	if len(needSpace) != 0 {
		ns = needSpace[0]
	}
	var arr []string
	for _, v := range strings.Split(s, sep) {
		if v == "" && !ns {
			continue
		}
		arr = append(arr, v)
	}
	return &T{arr}
}

func Convert(s []string) {

}

func Join(in interface{}, sep string) string {
	return gstr.JoinAny(in, sep)
}
