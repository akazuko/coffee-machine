package machine

import (
	"fmt"
)

// Valid beverages that can be served by the machine
const (
	BlackTea  = "black_tea"
	GreenTea  = "green_tea"
	HotCoffee = "hot_coffee"
	HotTea    = "hot_tea"
)

// Beverage holds the metadata of a valid beverage
type Beverage struct {
	Name        string
	Ingredients map[string]int
}

// CreateBeverage implements the factory that allows creation of a valid beverage
func CreateBeverage(beverageType string) (*Beverage, error) {
	switch beverageType {
	case HotTea:
		ingridients := map[string]int{
			HotWater:       200,
			HotMilk:        100,
			TeaLeavesSyrup: 30,
			GingerSyrup:    10,
			SugarSyrup:     10,
		}
		return &Beverage{Name: beverageType, Ingredients: ingridients}, nil
	case HotCoffee:
		ingridients := map[string]int{
			HotWater:       100,
			HotMilk:        400,
			TeaLeavesSyrup: 30,
			GingerSyrup:    30,
			SugarSyrup:     50,
		}
		return &Beverage{Name: beverageType, Ingredients: ingridients}, nil
	case BlackTea:
		ingridients := map[string]int{
			HotWater:       300,
			GingerSyrup:    30,
			TeaLeavesSyrup: 30,
			SugarSyrup:     50,
		}
		return &Beverage{Name: beverageType, Ingredients: ingridients}, nil
	case GreenTea:
		ingridients := map[string]int{
			HotWater:        100,
			GingerSyrup:     30,
			SugarSyrup:      50,
			"green_mixture": 30,
		}
		return &Beverage{Name: beverageType, Ingredients: ingridients}, nil
	default:
		err := fmt.Errorf("alas, no such beverage => %v is available today", beverageType)
		return nil, err
	}
}
