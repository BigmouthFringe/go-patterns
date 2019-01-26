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
	fmt.Println("Squeak! *weird noises following...*")
}

func (FlyWithWings) Fly() {
	fmt.Println("*flying away towards sunset...*")
}
func (FlyNoWay) Fly() {
	fmt.Println("*tries to fly without success...*")
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

