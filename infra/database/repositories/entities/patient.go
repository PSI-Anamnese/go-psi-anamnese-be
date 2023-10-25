package entities

var _ Entity = (*Patient)(nil)

type Patient struct {
	id string
}

func (p Patient) GetId() string {
	return p.id
}
