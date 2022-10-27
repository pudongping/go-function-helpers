# 时间

a | b | c
--- | --- | --- 
[Date](#method-Date) | |

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