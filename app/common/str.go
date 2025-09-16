package common

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var StrTool = strTool{}

type strTool struct {
}

func (cls *strTool) Int(str string, defaultValue int) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		result = defaultValue
	}
	return result
}

func (cls *strTool) Int64(str string, defaultValue int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		result = defaultValue
	}
	return result
}

func (cls *strTool) Float64(str string, defaultValue float64) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		result = defaultValue
	}
	return result
}

// 首字母大写的函数
func (cls *strTool) CapitalizeFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

// 合并字符串
func (cls *strTool) Merge(sep string, data ...string) (str string) {
	if len(data) == 0 {
		return
	}
	for _, val := range data {
		str += sep + val
	}

	return str[len(sep):]
}

// 判断首字母是否为大写
func (cls *strTool) IsUppercaseFirstLetter(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsUpper(rune(s[0]))
}

// 字符串模版替换
func (cls *strTool) Expand(s string, data map[string]any) string {
	mapper := func(placeholderName string) string {
		value, exists := data[placeholderName]
		if !exists {
			return ""
		}

		strValue, ok := value.(string)
		if !ok {
			return ""
		}

		return strValue
	}
	s = os.Expand(s, mapper)
	return s
}

// md5生成
func (cls *strTool) MD5String(s string) string {
	sum := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", sum)
}
