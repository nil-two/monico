package monico

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var tempDir string

func TestMain(m *testing.M) {
	path, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Fprintln(os.Stderr, "monico_test:", err)
	}
	tempDir = path

	e := m.Run()
	defer os.Exit(e)

	os.RemoveAll(tempDir)
}

func TestPath(t *testing.T) {
	m, err := NewMoniter(tempDir)
	if err != nil {
		t.Errorf("NewMoniter returns %q, want nil", err)
	}
	expect := tempDir
	actual := m.Path()
	if actual != expect {
		t.Errorf("got %q, want %q",
			actual, expect)
	}
}

func TestErrorPath(t *testing.T) {
	_, err := NewMoniter("")
	if err == nil {
		t.Errorf("NewMoniter returns nil, want err")
	}
}

func TestDefaultPath(t *testing.T) {
	m, err := NewMoniterWithWD()
	if err != nil {
		t.Errorf("NewMoniter returns %q, want nil", err)
	}
	expect, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed get working directory")
	}
	actual := m.Path()
	if actual != expect {
		t.Errorf("got %q, want %q",
			actual, expect)
	}
}
