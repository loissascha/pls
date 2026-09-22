package plsfile

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := ReadFile("test.plsfile")
	if err != nil {
		t.Fatal(err)
	}
	for n, j := range f.Jobs {
		fmt.Println("job name:", n)
		fmt.Println(j)
	}
}
