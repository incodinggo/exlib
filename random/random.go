package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type R struct{}

var (
	dirLetter  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	dirNum     = "0123456789"
	dirSpacial = "~!@#$%^&*()_+-=*.,;:"
)

// Strings 字母和符号
func (r *R) Strings(l int) string {
	bytes := make([]byte, l)
	b := strings.Builder{}
	b.WriteString(dirLetter)
	b.WriteString(dirSpacial)
	dir := b.String()
	dl := len(dir)
	for i := 0; i < l; i++ {
		bytes[i] = dir[rand.Intn(dl)]
	}
	return string(bytes)
}

// Letters 仅字母
func (r *R) Letters(l int) string {
	bytes := make([]byte, l)
	dl := len(dirLetter)
	for i := 0; i < l; i++ {
		bytes[i] = dirLetter[rand.Intn(dl)]
	}
	return string(bytes)
}

// Mixes 字母数字符号
func (r *R) Mixes(l int, dirEx string) string {
	bytes := make([]byte, l)
	b := strings.Builder{}
	b.WriteString(dirLetter)
	b.WriteString(dirSpacial)
	b.WriteString(dirNum)
	b.WriteString(dirEx)
	dir := b.String()
	dl := len(dir)
	for i := 0; i < l; i++ {
		bytes[i] = dir[rand.Intn(dl)]
	}
	return string(bytes)
}

// Custom 自定义
func (r *R) Custom(l int, dir string) string {
	bytes := make([]byte, l)
	dl := len(dir)
	for i := 0; i < l; i++ {
		bytes[i] = dir[rand.Intn(dl)]
	}
	return string(bytes)
}

// Nums 数字0-9
func (r *R) Nums(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(rand.Intn(10))
	}
	return string(bytes)
}

// Num 数字0-9 1个
func (r *R) Num(l int) int {
	return rand.Intn(10)
}

// Range 范围
func (r *R) Range(min, max int) string {
	if min == max {
		return fmt.Sprint(min)
	}
	if min < max {
		min, max = max, min
	}
	return fmt.Sprint(rand.Intn(max-min) + min)
}

// RangeNum Range 范围数字
func (r *R) RangeNum(min, max int) int {
	if min == max {
		return min
	}
	if min < max {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

// Rand 随机0-max的整数或{a-z|A-Z}区间的字母，max必须>0
func Rand() *R {
	rand.Seed(time.Now().UnixNano())
	return &R{}
}
