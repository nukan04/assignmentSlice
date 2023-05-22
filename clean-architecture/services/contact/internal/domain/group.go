package domain

type Group struct {
	Id   int
	Name string
}

func NewGroup(name string) *Group {
	if len(name) > 250 {
		return nil
	}

	return &Group{Name: name}
}
