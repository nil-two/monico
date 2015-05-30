package monico

type Moniter struct {
	path string
}

func NewMoniter(path string) (*Moniter, error) {
	return &Moniter{
		path: path,
	}, nil
}

func (m *Moniter) Path() string {
	return m.path
}
