package util

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	fmt.Println(Empty(""))            // true
	fmt.Println(Empty(0))             // true
	fmt.Println(Empty(0.0))           // true
	fmt.Println(Empty([]int{}))       // true
	fmt.Println(Empty(map[int]int{})) // true
	fmt.Println(Empty(struct {
	}{})) // true

	fmt.Println(Empty(" "))  // false
	fmt.Println(Empty(0.01)) // false
}

func TestMd5(t *testing.T) {
	s1 := "hello I'm Alex"
	s2 := Md5(s1)
	s3 := Md5(s1)

	fmt.Println(s2)       // 7c3e7d99dfccf33b6a10de103957694e
	fmt.Println(s3)       // 7c3e7d99dfccf33b6a10de103957694e
	fmt.Println(s2 == s3) // true
}

func TestBcryptHash(t *testing.T) {
	pwd := "123456"

	hash := BcryptHash(pwd)
	// $2a$14$KGnHvQVNo4qSnoFRKVaJXuRNP9F45IlCI6CowkqXQxKT8wO5ta7nG
	fmt.Println(hash)

	check := BcryptCheck(pwd, hash)
	// true
	fmt.Println(check)

}
