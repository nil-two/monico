package monico

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
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
	actual := m.path
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
	actual := m.path
	if actual != expect {
		t.Errorf("got %q, want %q",
			actual, expect)
	}
}

func TestModified(t *testing.T) {
	m, err := NewMoniter(tempDir)
	if err != nil {
		t.Errorf("NewMoniter returns %q, want nil", err)
	}
	now := time.Now()
	if err = os.Chtimes(tempDir, now, now); err != nil {
		t.Errorf("Failed os.Chtimes(%q, %v, %v)",
			tempDir, now, now)
	}

	expect := true
	actual, err := m.Modified()
	if err != nil {
		t.Errorf("Modified returns %q, want nil", err)
	}
	if actual != expect {
		t.Errorf("got %v, want %v", actual, expect)
	}
}

func TestNotModified(t *testing.T) {
	m, err := NewMoniter(tempDir)
	if err != nil {
		t.Errorf("NewMoniter returns %q, want nil", err)
	}

	expect := false
	actual, err := m.Modified()
	if err != nil {
		t.Errorf("Modified returns %q, want nil", err)
	}
	if actual != expect {
		t.Errorf("got %v, want %v", actual, expect)
	}
}

func TestUpdateModTime(t *testing.T) {
	m, err := NewMoniter(tempDir)
	if err != nil {
		t.Errorf("NewMoniter returns %q, want nil", err)
	}
	now := time.Now()
	if err = os.Chtimes(tempDir, now, now); err != nil {
		t.Errorf("Failed os.Chtimes(%q, %v, %v)",
			tempDir, now, now)
	}
	if err = m.UpdateModTime(); err != nil {
		t.Errorf("UpdateModTime returns %q, want nil", err)
	}

	expect := now
	actual := m.lastModTime
	if actual != expect {
		t.Errorf("got %v, want %v",
			actual, expect)
	}
}
