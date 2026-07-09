// Package regexpx 提供常用正则表达式的便捷封装：
// - 提取数字、替换空白、分割字符串、判断字母数字
package regexpx

import "regexp"

var (
	reDigits       = regexp.MustCompile(`\d+`)
	reWhitespace   = regexp.MustCompile(`\s+`)
	reAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
)

// MatchDigits 提取字符串中所有连续数字。
func MatchDigits(s string) []string {
	return reDigits.FindAllString(s, -1)
}

// ReplaceWhitespace 将连续空白字符替换为单个空格。
func ReplaceWhitespace(s string) string {
	return reWhitespace.ReplaceAllString(s, " ")
}

// SplitByComma 按逗号分割字符串。
func SplitByComma(s string) []string {
	return regexp.MustCompile(`,`).Split(s, -1)
}

// IsAlphanumeric 判断字符串是否只包含字母和数字。
func IsAlphanumeric(s string) bool {
	return reAlphanumeric.MatchString(s)
}
