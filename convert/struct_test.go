package convert

import (
	"fmt"
	"testing"
)

func TestStruct2Map(t *testing.T) {
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

	r1 := Struct2Map(alex)
	// typeof map[string]interface {} ==> map[LikeDoing:[watch TV play basketball play ping-pong ball] age:18 name:alex]
	fmt.Printf("typeof %T ==> %+v \n", r1, r1)
}
