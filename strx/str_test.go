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

func TestSnake(t *testing.T) {
	fmt.Println(Snake("fooBar"))      // foo_bar
	fmt.Println(Snake("我是张三AbCDt"))   // 我是张三_ab_c_dt
	fmt.Println(Snake("我是张三AbCDt%d")) // 我是张三_ab_c_dt%d
}

func TestStudly(t *testing.T) {
	fmt.Println(Studly("foo_bar")) // FooBar
}

func TestCamel(t *testing.T) {
	fmt.Println(Camel("foo_bar")) // fooBar
}

func TestCoalesce(t *testing.T) {
	fmt.Println(Coalesce("11", "+", "22", "=", "33"))              // 11+22=33
	fmt.Println(Coalesce([]string{"11", "+", "22", "=", "33"}...)) // 11+22=33
}
