package person

import (
	"sync"
)

type PersonCache struct {
	cache   map[int]Person
	mux     sync.RWMutex
	counter int
}

func NewPersonCache() *PersonCache {
	var mux sync.RWMutex
	cache := make(map[int]Person)
	return &PersonCache{
		mux:     mux,
		cache:   cache,
		counter: 0,
	}
}

func (personCache *PersonCache) AddPerson(person Person) int {
	personCache.mux.Lock()
	defer personCache.mux.Unlock()
	personId := personCache.counter
	personCache.counter++
	personCache.cache[personId] = person
	return personId
}

func (personCache *PersonCache) GetPerson(personId int) Person {
	personCache.mux.RLock()
	defer personCache.mux.RUnlock()
	person := personCache.cache[personId]
	return person
}
