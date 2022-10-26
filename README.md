# go-function-helpers

本项目收录一些高频、实用的 Go 语言助手函数，函数名称尽可能的和 PHP 中同功能函数一致，避免重复造轮子，提高开发效率。  
由于 Go 语言为强类型语言，因此本项目中的助手函数可能只会提供一种数据类型的作为参考，
要想兼容其他数据类型，请找到符合需求的助手函数之后，复制一份，然后自行调整为你所期望的数据类型。  
**因此，本项目更像是一个资料库，而不像是一个插件库。**

## 实用的工具推荐

- [JSON-to-Go](https://mholt.github.io/json-to-go/) —— 在线将 json 结构转换为 Go 类型定义
- [php2golang](https://www.php2golang.com/) —— PHP 函数到 Go 语言的转化

## 可用方法

### [数组 & 切片](#class-Arrayx)

a | b | c
--- | --- | --- 
[InArray](#method-InArray) | [ArrayChunk](#method-ArrayChunk) | [ArrayUnique](#method-ArrayUnique)

### [字符串](#class-Strx)

a | b | c
--- | --- | --- 
[StrRandom](#method-StrRandom) | |

### [类型转换](#class-Convert)

a | b | c
--- | --- | --- 
[Struct2Map](#method-Struct2Map) | [Map2Struct](#method-Map2Struct) |

## 方法列表

### 数组 & 切片
<p id="class-Arrayx"></p>

<a name="method-InArray"></a>
#### `InArray`

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

#### `ArrayChunk`
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

#### `ArrayUnique`
<p id="method-ArrayUnique"></p>

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

### 字符串
<p id="class-Strx"></p>

#### `StrRandom`
<p id="method-StrRandom"></p>

根据指定的种子数生成指定长度的随机字符串

```go
package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/strx"
)

func main() {
	// iAgDdAe9JMo0wlrY
	fmt.Println(strx.StrRandom(16, strx.LowerCase+strx.UpperCase+strx.Numeric))
}
```

### 类型转换
<p id="class-Convert"></p>

#### `Struct2Map`
<p id="method-Struct2Map"></p>

结构体转字典

#### `Map2Struct`
<p id="method-Map2Struct"></p>

字典转结构体

```go
package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/convert"
)

func main() {
	type student struct {
		Name      string `json:"name"`
		Age       uint   `json:"age"`
		LikeDoing []string
	}

	alex := student{
		Name:      "alex",
		Age:       18,
		LikeDoing: []string{"watch TV", "play basketball", "play ping-pong ball"},
	}

	r1 := convert.Struct2Map(alex)
	// typeof map[string]interface {} ==> map[LikeDoing:[watch TV play basketball play ping-pong ball] age:18 name:alex]
	fmt.Printf("typeof %T ==> %+v \n", r1, r1)

	var harry student
	_ = convert.Map2Struct(r1, &harry)
	// typeof main.student ==> {Name:alex Age:18 LikeDoing:[watch TV play basketball play ping-pong ball]}
	fmt.Printf("typeof %T ==> %+v \n", harry, harry)
}
```