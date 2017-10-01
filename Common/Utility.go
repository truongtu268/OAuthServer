package Common

import (
	"fmt"
	"encoding/json"
	"strings"
	"crypto/rand"
	"encoding/base64"
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

func StringContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
