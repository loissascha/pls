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

func TestRemoveComment(t *testing.T) {

	// start str with a comment
	startStr := "	go run . // this runs the code "
	endStr := removeCommentFromLine(startStr)
	if endStr != "	go run . " {
		t.Errorf("failed to remove comment from line: '%s' .. result: '%s'", startStr, endStr)
	}

	// start str without a comment
	startStr = "go run ."
	endStr = removeCommentFromLine(startStr)
	if endStr != "go run ." {
		t.Errorf("failed to parse line without comment: '%s' .. result: '%s'", startStr, endStr)
	}

	// comment line
	startStr = "// this is a comment line"
	endStr = removeCommentFromLine(startStr)
	if endStr != "" {
		t.Errorf("failed to parse comment line: '%s' .. result: '%s'", startStr, endStr)
	}
}
