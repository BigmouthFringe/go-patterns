package strategy

import "fmt"

type Quackable interface {
	Quack()
}
type Flyable interface {
	Fly()
}

type Quack struct {}
type Squeak struct {}

type FlyWithWings struct {}
type FlyNoWay struct {}

func (Quack) Quack() {
	fmt.Println("Quack!")
}
func (Squeak) Quack() {
	fmt.Println("Squeak!")
}

func (FlyWithWings) Fly() {
	fmt.Println("*Flying away towards sunset...*")
}
func (FlyNoWay) Fly() {
	fmt.Println("*It just can't fly, I'm sorry*")
}

type Duck struct {
	quackable Quackable
	flyable Flyable
}

func (d Duck) New(quackable Quackable, flyable Flyable) *Duck {
	d.quackable = quackable
	d.flyable = flyable
	return &d
}

func (d *Duck) PerformFly() {
	d.flyable.Fly()
}
func (d *Duck) PerformQuack() {
	d.quackable.Quack()
}

