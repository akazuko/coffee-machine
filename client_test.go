package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/akazuko/coffee-machine/machine"
)

// TestClient tests the functional flow of the coffee-machine
func TestClient(t *testing.T) {
	m := machine.NewMachine("Coffee Machine", 4)
	m.AddIngredient(machine.HotWater, 500)
	m.AddIngredient(machine.HotMilk, 500)
	m.AddIngredient(machine.GingerSyrup, 100)
	m.AddIngredient(machine.SugarSyrup, 100)
	m.AddIngredient(machine.TeaLeavesSyrup, 100)

	var wg sync.WaitGroup
	expectedErrorCount := 0
	beverages := [...]string{machine.HotTea, machine.HotCoffee, machine.BlackTea, machine.GreenTea, "invalid_drink"}

	for i, beverageType := range beverages {
		c := &Client{Name: "client " + strconv.Itoa(i)}
		go func(beverageType string) {
			err := c.RequestBeverage(&wg, m, beverageType)
			if err != nil {
				fmt.Println(err)
				expectedErrorCount++
			} else {
				fmt.Printf("%v is prepared\n", beverageType)
			}
		}(beverageType)
	}

	time.Sleep(10 * time.Millisecond)
	wg.Wait()

	if expectedErrorCount != 3 {
		t.Errorf("there should be three induces errors out of five scenarios")
	}
}
