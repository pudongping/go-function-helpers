package arrayx

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	haystack := []int{1, 2, 3, 4}

	if InArray(3, haystack) {
		fmt.Println("3 存在数组中")
	} else {
		fmt.Println("3 不存在数组中")
	}
}

func TestArrayChunk(t *testing.T) {
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	size := 3

	r1 := ArrayChunk(i, size)
	// [[0 1 2] [3 4 5] [6 7 8] [9]]
	fmt.Println(r1)

	r1[2][1] = 100
	// [[0 1 2] [3 4 5] [6 100 8] [9]]
	fmt.Println(r1)

	// [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(i)
}
