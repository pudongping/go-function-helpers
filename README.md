# go-function-helpers

本项目收录一些高频、实用的 Go 语言助手函数，函数名称尽可能的和 PHP 中同功能函数一致，避免重复造轮子，提高开发效率。  
由于 Go 语言为强类型语言，因此本项目中的助手函数可能只会提供一种数据类型的作为参考，
要想兼容其他数据类型，请找到符合需求的助手函数之后，复制一份，然后自行调整为你所期望的数据类型。  
**因此，本项目更像是一个资料库，而不像是一个插件库。**

## 实用的工具推荐

- [JSON-to-Go](https://mholt.github.io/json-to-go/) —— 在线将 json 结构转换为 Go 类型定义
- [php2golang](https://www.php2golang.com/) —— PHP 函数到 Go 语言的转化

## 可用方法

### 数组 & 切片

[InArray](#method-InArray)
[ArrayChunk](#method-ArrayChunk)

## 方法列表

#### InArray
<p id="method-InArray"></p>

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

#### ArrayChunk
<p id="method-ArrayChunk"></p>

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