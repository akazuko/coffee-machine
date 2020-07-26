package machine

import (
	"fmt"
	"sync"
	"time"
)

const (
	preparationTime = 1 // unit: seconds
)

// Machine to be used by the client
type Machine struct {
	// config
	Name    string
	Outlets int

	// synchronization helpers
	sem chan string
	mx  sync.Mutex

	// in-memory store
	store map[string]int
}

// NewMachine creates machines with required name & outlets.
func NewMachine(name string, outlets int) *Machine {
	m := &Machine{
		Name:    name,
		Outlets: outlets,
		sem:     make(chan string, outlets),
		store:   make(map[string]int),
	}
	return m
}

// AddIngredient allows addition of the valid ingredient in store
func (m *Machine) AddIngredient(name string, amount int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.store[name] += amount
}

// PrepareBeverage prepares the beverage in the machine only if the ingredients
// in the sufficient amount exists
func (m *Machine) PrepareBeverage(b *Beverage) error {
	m.sem <- ""
	m.mx.Lock()
	defer func() {
		m.mx.Unlock()
		<-m.sem
	}()

	for ingredient, requiredAmount := range b.Ingredients {
		currAmount, ingredientExists := m.store[ingredient]

		// check if the ingredient exists
		if !ingredientExists {
			return fmt.Errorf("%v can not be prepared because %v is not available", b.Name, ingredient)
		}

		// check if the ingredient is enough
		if requiredAmount > currAmount {
			return fmt.Errorf("%v can not be prepared because %v is not sufficient", b.Name, ingredient)
		}
	}

	// all OK, time to prepare the beverage
	for ingredient, requiredAmount := range b.Ingredients {
		currAmount, _ := m.store[ingredient]
		m.store[ingredient] = currAmount - requiredAmount
	}

	// some constant time for beverage preparation
	time.Sleep(preparationTime * 1000 * time.Millisecond)

	return nil
}
