package testexport

type Thing struct {
	B Backend
}

func NewThing() *Thing {
	return &Thing{B: &backend{}}
}

type backend struct {
}

func (b *backend) Do(BackendThing) {
}

func (t *Thing) SetBackend(b Backend) {
	t.B = b
}
