package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	now1 := Date("Y-m-d H:i:s")
	now2 := Date("Y-m-d H:i:s", time.Now())
	now3 := Date("Y-m-d H:i:s", time.Now().AddDate(0, 1, 0))

	fmt.Println(now1)
	fmt.Println(now2)
	fmt.Println(now3)
}