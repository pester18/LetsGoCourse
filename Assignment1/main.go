package main

import (
	"fmt"

	"./person"
	"./profession"
	"./worker"
)

func main() {
	wellKnownPerson := person.Person{Name: "Marcus", Surname: "Aurelius"}
	anotherWellKnownPerson := person.Person{Name: "Rene", Surname: "Descartes"}

	fmt.Println(wellKnownPerson.Name, wellKnownPerson.Surname)
	fmt.Println(anotherWellKnownPerson.Name, anotherWellKnownPerson.Surname)

	emperor := worker.Worker{
		Person: &wellKnownPerson,
		Profession: &profession.Emperor{
			Dynasty:      "Antonini",
			TermOfOffice: "161-180 A.D.",
		},
	}
	philosopher := worker.Worker{
		Profession: &profession.Philosopher{
			Direction: "rationalism",
		},
	}
	// fmt.Println(philosopher.Introduce())	  //Uncommenting this line will make the code panic (nil pointer dereferencing)
	philosopher.Person = &anotherWellKnownPerson

	fmt.Println(emperor.Introduce())
	fmt.Println(philosopher.Introduce())
	fmt.Printf("%s %s job is: %s\n", emperor.Name, emperor.Surname, emperor.Position())
	fmt.Printf("%s %s job is: %s\n", philosopher.Name, philosopher.Surname, philosopher.Position())
}
