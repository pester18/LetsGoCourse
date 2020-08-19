package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../person"
)

var target = flag.String("target", "http://localhost:8090", "request target")

func main() {
	flag.Parse()

	wellKnownPerson := person.Person{Name: "Marcus", Surname: "Aurelius"}
	anotherWellKnownPerson := person.Person{Name: "Rene", Surname: "Descartes"}

	id1, _ := AddPerson(wellKnownPerson)
	id2, _ := AddPerson(anotherWellKnownPerson)
	fmt.Println("First id is equal:", id1)
	fmt.Println("Second id is equal:", id2)

	p1, _ := GetPerson(id1)
	p2, _ := GetPerson(id2)

	fmt.Printf("First response: %+v\n", p1)
	fmt.Printf("Second response: %+v\n", p2)
	fmt.Println("Marcus Aurelius equals Marcus Aurelius:", p1 == wellKnownPerson)
	fmt.Println("Rene Descartes equals Rene Descartes:", p2 == anotherWellKnownPerson)
}

func GetPerson(id int) (pers person.Person, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/?id=%d", *target, id))
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &pers)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return pers, nil
}

func AddPerson(pers person.Person) (id int, err error) {
	requestBody, err := json.Marshal(pers)
	if err != nil {
		log.Fatalln(err)
		return
	}

	url := fmt.Sprintf("%s/", *target)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	var res struct{ Id int }
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return res.Id, nil
}
