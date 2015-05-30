package monico

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	e := m.Run()
	defer os.Exit(e)
}
