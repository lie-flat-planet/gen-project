package example

import (
	"github.com/lie-flat-planet/gen-project/generator"
	"testing"
)

func TestDirGenerate(t *testing.T) {
	err := generator.Generate("./", "github.com/lie-flat-planet/gin-demo")
	//err := dir.Generate("./", "liming")
	if err != nil {
		t.Fatal(err)
	}
}
