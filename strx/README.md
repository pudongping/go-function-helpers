# 字符串

a | b                 | c
--- |-------------------| --- 
[StrRandom](#method-StrRandom) | [Snake](#method-Snake) |

## `StrRandom`
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

## `Snake`
<a name="method-Snake"></a>

将给定的字符串转换为 snake_case「蛇式」

```go
package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/strx"
)

func main() {
	fmt.Println(strx.Snake("fooBar"))      // foo_bar
	fmt.Println(strx.Snake("我是张三AbCDt"))   // 我是张三_ab_c_dt
	fmt.Println(strx.Snake("我是张三AbCDt%d")) // 我是张三_ab_c_dt%d
}
```