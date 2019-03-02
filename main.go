package main

import (
	"fmt"
	"github.com/bigmouthfringe/go-patterns/factory"
)

func main() {
	pizzaSort := factory.Clam
	pizza := factory.CreatePizza(pizzaSort)
	fmt.Printf("%T, %#v", pizza, *pizza)
}
