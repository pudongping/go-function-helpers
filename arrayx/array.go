package arrayx

// InArray 判断某个值是否在切片中存在
// needle := 3
// haystack := []int{1, 2, 3}
// output: true
func InArray(needle int, haystack []int) bool {
	if len(haystack) == 0 {
		return false
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

// ArrayChunk 将一个切片分割成多个子切片
// i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// size := 3
// output: [[0 1 2] [3 4 5] [6 7 8] [9]]
func ArrayChunk(j []int, size int) [][]int {
	var i = make([]int, len(j))
	copy(i, j)
	batches := make([][]int, 0, (len(i)+size-1)/size)

	for size < len(i) {
		i, batches = i[size:], append(batches, i[0:size:size])
	}
	batches = append(batches, i)

	return batches
}
