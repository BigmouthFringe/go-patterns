package decorator

import "fmt"

// Decorator pattern attaches additional responsibilities
// to an object dynamically (at runtime)
// Decorators provide a flexible alternative to
// subclassing for extending functionality

type Beverage interface {
	Description() string
	Cost() float64
}

type Espresso struct {}
func (Espresso) New() *Espresso {
	return new(Espresso)
}
func (*Espresso) Description() string {
	return "Espresso"
}
func (*Espresso) Cost() float64 {
	return 1.99
}

type HouseBlend struct {}
func (HouseBlend) New() *HouseBlend {
	return new(HouseBlend)
}
func (*HouseBlend) Description() string {
	return "House Blend Coffee"
}
func (*HouseBlend) Cost() float64 {
	return 0.89
}

// Each decorator wraps the component
// or in other words has-a component
type Condiment interface {
	New(beverage Beverage)
}

// Concrete Decorator has an instance
// variable for the thing it decorates
type Mocha struct {
	beverage Beverage
}
func (Mocha) New(beverage Beverage) *Mocha {
	var m = new(Mocha)
	m.beverage = beverage
	return m
}
func (m *Mocha) Description() string {
	return m.beverage.Description() + ", Mocha"
}
func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}

type Whip struct {
	beverage Beverage
}
func (Whip) New(beverage Beverage) *Whip {
	var w = new(Whip)
	w.beverage = beverage
	return w
}
func (w *Whip) Description() string {
	return w.beverage.Description() + ", Whip"
}
func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.10
}

func ServeCoffee() {
	var espresso = Espresso{}.New()
	fmt.Println(espresso.Description(), "$", espresso.Cost())

	var houseBlend Beverage
	houseBlend = HouseBlend{}.New()
	houseBlend = Mocha{}.New(houseBlend)
	houseBlend = Mocha{}.New(houseBlend)
	houseBlend = Whip{}.New(houseBlend)
	fmt.Println(houseBlend.Description(), "$", houseBlend.Cost())
}
