package worker

import (
	"../person"
	"../profession"
)

type Worker struct {
	person.Person
	profession.Profession
}
