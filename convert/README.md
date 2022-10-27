### [类型转换](#class-Convert)

a | b | c
--- | --- | --- 
[Struct2Map](#method-Struct2Map) | [Map2Struct](#method-Map2Struct) |

### 类型转换
<a name="class-Convert"></a>

#### `Struct2Map`
<a name="method-Struct2Map"></a>

结构体转字典

#### `Map2Struct`
<a name="method-Map2Struct"></a>

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