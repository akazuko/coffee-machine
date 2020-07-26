package main

import (
	"sync"

	"github.com/akazuko/coffee-machine/machine"
)

// Client who will use the coffee-machine
type Client struct {
	Name string
}

// RequestBeverage allows the client to request a beverage from the machine
func (client *Client) RequestBeverage(wg *sync.WaitGroup, m *machine.Machine, beverageType string) error {
	wg.Add(1)
	defer wg.Done()

	beverage, err := machine.CreateBeverage(beverageType)
	if err != nil {
		return err
	}

	err = m.PrepareBeverage(beverage)
	if err != nil {
		return err
	}

	return nil
}
