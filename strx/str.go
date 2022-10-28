package strx

import (
	"crypto/rand"
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
