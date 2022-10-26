package convert

import (
	"encoding/json"
)

// Map2Struct 字典转结构体
func Map2Struct(input interface{}, output interface{}) error {
	res, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, output)
}
