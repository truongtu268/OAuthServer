package Common

import (
	"fmt"
	"encoding/json"
	"strings"
)

func MapObject(dto interface{},entity interface{}) {
	b,err := json.Marshal(dto)
	if err != nil {
		fmt.Println("Error:%s ",err)
	}
	if err = json.NewDecoder(strings.NewReader(string(b))).Decode(entity);err != nil {
		fmt.Println(err)
	}
}
