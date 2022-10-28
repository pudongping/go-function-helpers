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

func TestStrToTime(t *testing.T) {
	layout := "2006-01-02 15:04:05"

	t1, _ := StrToTime("2022-10-28 00:42:09")
	// t1 typeof time.Time ====> 2022-10-28 00:42:09 +0800 CST
	fmt.Printf("t1 typeof %T ====> %v \n", t1, t1)

	t2, _ := StrToTime("2022-10-28 00:42:09", layout)
	// t2 typeof time.Time ====> 2022-10-28 00:42:09 +0800 CST
	fmt.Printf("t2 typeof %T ====> %v \n", t2, t2)

	t3, _ := StrToTime("2022-10-28", "2006-01-02")
	// t3 typeof time.Time ====> 2022-10-28 00:00:00 +0800 CST
	fmt.Printf("t3 typeof %T ====> %v \n", t3, t3)

	// 1666886400
	fmt.Println(t3.Unix())
	// 2022-10-28 00:00:00
	fmt.Println(t3.Format(layout))
}
