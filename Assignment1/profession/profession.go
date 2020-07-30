package profession

import "fmt"

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

func (e *Emperor) Position() string {
	return fmt.Sprintf("Emperor from %s dynasty served %s", e.Dynasty, e.TermOfOffice)
}

func (ph *Philosopher) Position() string {
	return fmt.Sprintf("Philosopher, representative of %s school", ph.Direction)
}
