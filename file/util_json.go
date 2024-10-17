package file

import (
	"fmt"
	"strings"
)

type UtilJson struct{}

func (json *UtilJson) Name() string {
	return "json.go"
}

func (json *UtilJson) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
package util

import (
	"encoding/json"
	"fmt"
)

func LogJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
`, ))
}
