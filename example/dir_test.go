package example

import (
	"github.com/lie-flat-planet/gen-project/dir/Generator"
	"testing"
)

func TestDirGenerate(t *testing.T) {
	err := Generator.Generate("./", "liming")
	//err := dir.Generate("./", "liming")
	if err != nil {
		t.Fatal(err)
	}
}
