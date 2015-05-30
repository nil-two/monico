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

func (m *Moniter) waitContinuousModify() error {
	lastModTime := m.lastModTime
	for {
		time.Sleep(100 * time.Millisecond)
		info, err := os.Stat(m.path)
		if err != nil {
			return err
		}

		newModTime := info.ModTime()
		if lastModTime.Equal(newModTime) {
			break
		}
		lastModTime = newModTime
	}
	return nil
}

func (m *Moniter) Modified() (bool, error) {
	info, err := os.Stat(m.path)
	if err != nil {
		return false, err
	}
	modified := !m.lastModTime.Equal(info.ModTime())
	if !modified {
		return false, nil
	}
	if err = m.waitContinuousModify(); err != nil {
		return false, err
	}
	return true, nil
}

func (m *Moniter) UpdateModTime() error {
	info, err := os.Stat(m.path)
	if err != nil {
		return err
	}
	m.lastModTime = info.ModTime()
	return nil
}
