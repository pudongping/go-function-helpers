package convert

import (
	"fmt"
	"testing"
)

func TestMap2Struct(t *testing.T) {
	input := map[string]interface{}{
		"LikeDoing": []string{"watch TV", "play basketball", "play ping-pong ball"},
		"age":       18,
		"name":      "alex",
	}

	type student struct {
		Name      string `json:"name"`
		Age       uint   `json:"age"`
		LikeDoing []string
	}

	var alex student
	_ = Map2Struct(input, &alex)
	// typeof convert.student ===> {Name:alex Age:18 LikeDoing:[watch TV play basketball play ping-pong ball]}
	fmt.Printf("typeof %T ===> %+v \n", alex, alex)
}
