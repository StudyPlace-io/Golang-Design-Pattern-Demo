package main

import (
	"fmt"
	"golanglearning/new_project/for-gong-zhong-hao/Design-Patterns-practice/Structural/DecoratorPattern"
)

func main() {
	pizza := DecoratorPattern.NewVeggeMania(20)

	//Add cheese topping
	pizzaWithCheese := &DecoratorPattern.CheeseTopping{
		Pizza: pizza,
		CheeseToppingPrice: 12,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &DecoratorPattern.TomatoTopping{
		Pizza: pizzaWithCheese,
		TomatoToppingPrice: 10,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}



