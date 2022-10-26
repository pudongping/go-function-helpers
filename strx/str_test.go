package strx

import (
	"fmt"
	"testing"
	"time"
)

func TestStrRandom(t *testing.T) {
	var r1 = make([]string, 0, 20)
	a := LowerCase + UpperCase + Symbol + Numeric
	for i := 0; i <= 20; i++ {
		go func(j int, xx string) {
			r2 := StrRandom(j, xx)
			r1 = append(r1, r2)

			fmt.Printf("%v   len ====> %s \n", len(r2), r2)
		}(i, a)
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("len %+v cap %+v \n rr %v \n", len(r1), cap(r1), r1)
}
