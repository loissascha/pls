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
	buildJob, found := f.Jobs["build"]
	if !found {
		t.Fatal("job 'build' not found")
	}
	fmt.Println("buildJob", buildJob)
	testprintJob, found := f.Jobs["testprint"]
	if !found {
		t.Fatal("job 'testprint' not found")
	}
	fmt.Println("testprintJob", testprintJob)
}
