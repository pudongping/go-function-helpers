### [字符串](#class-Strx)

a | b | c
--- | --- | --- 
[StrRandom](#method-StrRandom) | |


### 字符串
<a name="class-Strx"></a>

#### `StrRandom`
<a name="method-StrRandom"></a>

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