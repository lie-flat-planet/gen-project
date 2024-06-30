package file

import "testing"

func TestProject(t *testing.T) {
	if err := CreateProjectDir("srv-test"); err != nil {
		t.Fatal(err)
	}
}
