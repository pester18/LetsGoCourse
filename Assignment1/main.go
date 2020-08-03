package main

import (
	"fmt"
	"reflect"

	"./person"
	"./profession"
	"./worker"
)

func pollWorker(cache map[string]worker.Worker) map[string]reflect.Type {
	professionTypeCache := make(map[string]reflect.Type)

	for name, worker := range cache {
		professionTypeCache[name] = reflect.TypeOf(worker.Profession)
	}

	return professionTypeCache
}

func main() {
	wellKnownPerson := person.Person{Name: "Marcus", Surname: "Aurelius"}
	anotherWellKnownPerson := person.Person{Name: "Rene", Surname: "Descartes"}

	emperor := worker.Worker{
		Person: &wellKnownPerson,
		Profession: &profession.Emperor{
			Dynasty:      "Antonini",
			TermOfOffice: "161-180 A.D.",
		},
	}
	philosopher := worker.Worker{
		Person: &anotherWellKnownPerson,
		Profession: &profession.Philosopher{
			Direction: "rationalism",
		},
	}

	workersCache := map[string]worker.Worker{
		"Marcus": emperor,
		"Rene":   philosopher,
	}

	professionTypeCache := pollWorker(workersCache)

	fmt.Println(professionTypeCache)
}
