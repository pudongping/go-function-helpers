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

func TestIsNumeric(t *testing.T) {
	v1 := 123
	v2 := "alex123"
	v3 := "123"
	v4 := "alex"
	v5 := 123.45
	v6 := []int{1, 2, 3}
	v7 := map[int]int{1: 11, 2: 22}
	v8 := struct {
	}{}

	fmt.Printf("v1 ==> %v typeof ==> %T is_numeric ==> %v \n", v1, v1, IsNumeric(v1)) // v1 ==> 123 typeof ==> int is_numeric ==> true
	fmt.Printf("v2 ==> %v typeof ==> %T is_numeric ==> %v \n", v2, v2, IsNumeric(v2)) // v2 ==> alex123 typeof ==> string is_numeric ==> false
	fmt.Printf("v3 ==> %v typeof ==> %T is_numeric ==> %v \n", v3, v3, IsNumeric(v3)) // v3 ==> 123 typeof ==> string is_numeric ==> true
	fmt.Printf("v4 ==> %v typeof ==> %T is_numeric ==> %v \n", v4, v4, IsNumeric(v4)) // v4 ==> alex typeof ==> string is_numeric ==> false
	fmt.Printf("v5 ==> %v typeof ==> %T is_numeric ==> %v \n", v5, v5, IsNumeric(v5)) // v5 ==> 123.45 typeof ==> float64 is_numeric ==> true
	fmt.Printf("v6 ==> %v typeof ==> %T is_numeric ==> %v \n", v6, v6, IsNumeric(v6)) // v6 ==> [1 2 3] typeof ==> []int is_numeric ==> false
	fmt.Printf("v7 ==> %v typeof ==> %T is_numeric ==> %v \n", v7, v7, IsNumeric(v7)) // v7 ==> map[1:11 2:22] typeof ==> map[int]int is_numeric ==> false
	fmt.Printf("v8 ==> %v typeof ==> %T is_numeric ==> %v \n", v8, v8, IsNumeric(v8)) // v8 ==> {} typeof ==> struct {} is_numeric ==> false
}
