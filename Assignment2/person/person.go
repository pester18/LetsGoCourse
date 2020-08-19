package person

import (
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("Nice to meet you, I am %s %s", p.Name, p.Surname)
}
