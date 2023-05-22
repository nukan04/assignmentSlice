package domain

type Contact struct {
	Id       int
	fullName fullName
	Phone    int
}

type fullName struct {
	Last   string
	First  string
	Middle string
}

func (c Contact) FullName() fullName {
	return c.fullName
}

func NewContact(last, first, middle string) *Contact {
	fname := fullName{
		Last:   last,
		First:  first,
		Middle: middle,
	}

	return &Contact{fullName: fname}
}
