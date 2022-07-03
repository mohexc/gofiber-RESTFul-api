package utils

import (
	"encoding/json"
	"fmt"
)

func Print2Json(data interface{}) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
