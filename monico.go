package monico

import (
	"os"
	"time"
)

type Moniter struct {
	path        string
	lastModTime time.Time
}

func NewMoniter(path string) (*Moniter, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &Moniter{
		path:        path,
		lastModTime: info.ModTime(),
	}, nil
}

func NewMoniterWithWD() (*Moniter, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return NewMoniter(wd)
}

func (m *Moniter) Path() string {
	return m.path
}

func (m *Moniter) Modified() (bool, error) {
	info, err := os.Stat(m.path)
	if err != nil {
		return false, err
	}
	return !m.lastModTime.Equal(info.ModTime()), nil
}

func (m *Moniter) LastModTime() time.Time {
	return m.lastModTime
}
