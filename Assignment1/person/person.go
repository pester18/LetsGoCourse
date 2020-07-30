package person

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("Nice to meet you, I am %s %s", p.Name, p.Surname)
}
