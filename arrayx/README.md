# 数组 & 切片

a | b | c
--- | --- | --- 
[InArray](#method-InArray) | [ArrayChunk](#method-ArrayChunk) | [ArrayUnique](#method-ArrayUnique)


<a name="method-InArray"></a>
## `InArray`

判断某个元素是否在某个数组中，存在则为 `true` 不存在则为 `false`

```go
package main

import (
    "fmt"

    "github.com/pudongping/go-function-helpers/arrayx"
)

func main() {
    is1 := arrayx.InArray(3, []int{1, 2, 3})
    is2 := arrayx.InArray(4, []int{1, 2, 3})

    fmt.Println(is1) // true
    fmt.Println(is2) // false
}
```

## `ArrayChunk`
<a name="method-ArrayChunk"></a>

将一个切片分割成多个子切片

```go
package main

import (
    "fmt"

    "github.com/pudongping/go-function-helpers/arrayx"
)

func main() {
    i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    size := 3

    r1 := arrayx.ArrayChunk(i, size)
	// [[0 1 2] [3 4 5] [6 7 8] [9]]
    fmt.Println(r1)

    r1[2][1] = 100
	// [[0 1 2] [3 4 5] [6 100 8] [9]]
    fmt.Println(r1)

	// [0 1 2 3 4 5 6 7 8 9]
    fmt.Println(i)
}
```

## `ArrayUnique`
<a name="method-ArrayUnique"></a>

移除切片中重复的值

```go
package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/arrayx"
)

func main() {
	// [1 2 3 4]
	fmt.Println(arrayx.ArrayUnique([]int{1, 2, 3, 4, 3, 2, 1}))
}
```