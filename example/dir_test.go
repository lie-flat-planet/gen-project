package example

import (
	"github.com/lie-flat-planet/gen-project/dir"
	"testing"
)

func TestDirGenerate(t *testing.T) {
	err := dir.Generate("./", "liming")
	if err != nil {
		t.Fatal(err)
	}
}
