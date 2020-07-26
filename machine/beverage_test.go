package machine

import (
	"fmt"
)

func ExampleCreateBeverage() {
	b, _ := CreateBeverage(HotTea)
	fmt.Printf("%v %v\n", b.Name, b.Ingredients)

	b, _ = CreateBeverage(HotCoffee)
	fmt.Printf("%v %v\n", b.Name, b.Ingredients)

	b, _ = CreateBeverage(BlackTea)
	fmt.Printf("%v %v\n", b.Name, b.Ingredients)

	b, _ = CreateBeverage(GreenTea)
	fmt.Printf("%v %v\n", b.Name, b.Ingredients)

	_, err := CreateBeverage("RandomTea")
	fmt.Println(err)

	// Output:
	// hot_tea map[ginger_syrup:10 hot_milk:100 hot_water:200 sugar_syrup:10 tea_leaves_syrup:30]
	// hot_coffee map[ginger_syrup:30 hot_milk:400 hot_water:100 sugar_syrup:50 tea_leaves_syrup:30]
	// black_tea map[ginger_syrup:30 hot_water:300 sugar_syrup:50 tea_leaves_syrup:30]
	// green_tea map[ginger_syrup:30 green_mixture:30 hot_water:100 sugar_syrup:50]
	// alas, no such beverage => RandomTea is available today
}
