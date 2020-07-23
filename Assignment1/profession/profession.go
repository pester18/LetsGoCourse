package profession

type Profession interface {
	Position() string
}

type Emperor struct {
	Dynasty      string
	TermOfOffice string
}

type Philosopher struct {
	Direction string
}

func (e Emperor) Position() string {
	return "Emperor"
}

func (e Philosopher) Position() string {
	return "Philosopher"
}
