package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/arrayx"
)

func main() {
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	size := 3

	r1 := arrayx.ArrayChunk(i, size)
	fmt.Println(r1)

	r1[2][1] = 100
	fmt.Println(r1)

	fmt.Println(i)
}
