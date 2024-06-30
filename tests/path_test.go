package tests

import (
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	t.Run("#path", func(t *testing.T) {
		path := "global/sa/"
		//p := filepath.Dir(path)
		//fmt.Println(p)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			t.Fatal(err)
		}
	})
}