package strategy

import (
	"fmt"
)

type NpcBehavior interface {
	Act()
}

type Talk struct {}
type Walk struct {}
type Attack struct {}

func (Talk) Act() {
	fmt.Println("Yo, what's up man")
}
func (Walk) Act() {
	fmt.Println("Walks away...")
}
func (Attack) Act() {
	fmt.Println("*Suddenly attacks you with the longsword!*")
}

type Npc struct {
	name string
	behavior NpcBehavior
}

func (n Npc) New(name string, behavior NpcBehavior) *Npc {
	n.name = name
	n.behavior = behavior
	return &n
}

func (n *Npc) Trigger() {
	n.behavior.Act()
}


