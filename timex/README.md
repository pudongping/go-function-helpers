# 时间

a | b              | c
--- |----------------| --- 
[Date](#method-Date) | [StrToTime](#method-StrToTime) |

## `Date`
<a name="method-Date"></a>

格式化时间或者日期

```go
package main

import (
    "fmt"
    "time"

    "github.com/pudongping/go-function-helpers/timex"
)

func main() {
    now1 := timex.Date("Y-m-d H:i:s")
    now2 := timex.Date("Y-m-d H:i:s", time.Now())
    now3 := timex.Date("Y-m-d H:i:s", time.Now().AddDate(0, 1, 0))

    fmt.Println(now1) // 2022-10-28 00:42:09
    fmt.Println(now2) // 2022-10-28 00:42:09
    fmt.Println(now3) // 2022-11-28 00:42:09
}
```

## `StrToTime`
<a name="method-StrToTime"></a>

时间字符串转时间格式

```go
package main

import (
	"fmt"

	"github.com/pudongping/go-function-helpers/timex"
)

func main() {
	t, _ := timex.StrToTime("2022-10-28 00:42:09")
	// typeof time.Time ====> 2022-10-28 00:42:09 +0800 CST
	fmt.Printf("typeof %T ====> %v \n", t, t)
}
```