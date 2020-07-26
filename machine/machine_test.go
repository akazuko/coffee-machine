package machine

import (
	"fmt"
)

// ExampleMachine tests machine state when created and an ingredient is added
func ExampleMachine() {
	m := NewMachine("test_machine", 1)
	fmt.Printf("%v %v %v\n", m.Name, m.Outlets, m.store)

	m.AddIngredient("ingredient1", 10)
	fmt.Printf("%v %v %v\n", m.Name, m.Outlets, m.store)

	// Output:
	// test_machine 1 map[]
	// test_machine 1 map[ingredient1:10]
}

func ExampleMachine_PrepareBeverage() {
	m := NewMachine("test_machine", 1)

	m.AddIngredient(HotMilk, 100)
	m.AddIngredient(HotWater, 200)
	m.AddIngredient(SugarSyrup, 10)
	m.AddIngredient(TeaLeavesSyrup, 30)

	b, _ := CreateBeverage(HotTea)
	fmt.Printf("%v %v\n", b.Name, b.Ingredients)

	var err error
	err = m.PrepareBeverage(b)
	if err != nil {
		fmt.Println(err)
	}

	m.AddIngredient(GingerSyrup, 5)

	err = m.PrepareBeverage(b)
	if err != nil {
		fmt.Println(err)
	}

	m.AddIngredient(GingerSyrup, 5)
	err = m.PrepareBeverage(b)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// hot_tea map[ginger_syrup:10 hot_milk:100 hot_water:200 sugar_syrup:10 tea_leaves_syrup:30]
	// hot_tea can not be prepared because ginger_syrup is not available
	// hot_tea can not be prepared because ginger_syrup is not sufficient
}
