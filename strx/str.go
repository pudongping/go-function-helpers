package strx

import (
	"crypto/rand"
	"strings"
	"unicode"
)

const (
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbol    = "!@#$%^&*()-_=+,.?/:;{}[]`~"
	Numeric   = "0123456789"
)

// StrRandom 随机生成指定长度的字符串
// length := 16  // 生成长度
// s := LowerCase + UpperCase + Numeric  // 种子
// output: iAgDdAe9JMo0wlrY
func StrRandom(length int, s string) string {
	var chars = []byte(s)
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length")
	}
	max := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > max {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

// Snake 将给定的字符串转换为 snake_case「蛇式」
// s := "fooBar"
// output: "foo_bar"
func Snake(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// Studly 将给定字符串「蛇式」转化为 camelCase「大驼峰式」
// s := "foo_bar"
// output: "FooBar"
func Studly(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s) // 首字母大写的格式
	return strings.Replace(s, " ", "", -1)
}

// Camel 将给定字符串「蛇式」转化为 camelCase「小驼峰式」
// s := "foo_bar"
// output: "fooBar"
func Camel(s string) string {
	s = Studly(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// Coalesce 高性能字符串拼接
// s1 := Coalesce("11", "+", "22", "=", "33")
// s2 := Coalesce([]string{"11", "+", "22", "=", "33"}...)
// s1 和 s2 都输出 "11+22=33"
func Coalesce(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	var str strings.Builder
	for _, v := range s {
		str.WriteString(v)
	}

	return str.String()
}
