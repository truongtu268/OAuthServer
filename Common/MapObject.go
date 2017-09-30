package Common

import (
	"fmt"
	"encoding/json"
	"strings"
)

func MapObject(input interface{}, output interface{}) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error:%s ", err)
	}
	if err = json.NewDecoder(strings.NewReader(string(b))).Decode(output); err != nil {
		fmt.Println(err)
	}
}
