package strategy

// The Strategy Pattern defines a family of algorithms,
// (in our case, how do ducks fly or quack, or how NPCs behave)
// encapsulates each one of them, and makes them interchangeable
// (that means we can't access them directly, and also can't change those algorithms)

// What we get in the result is that our Subjects behavior varies
// depending on the given to them algorithms at runtime through composition
func TriggerDucks() {
	var shortWingsDuck = Duck{}.New(Quack{}, FlyNoWay{})
	shortWingsDuck.PerformFly()
	shortWingsDuck.PerformQuack()

	var weirdDuck = Duck{}.New(Squeak{}, FlyWithWings{})
	weirdDuck.PerformFly()
	weirdDuck.PerformQuack()

	var rubberDuck = Duck{}.New(Squeak{}, FlyNoWay{})
	rubberDuck.PerformFly()
	rubberDuck.PerformQuack()
}
func TriggerNpcByName(name string) {
	switch name {
	case "Mark":
		var viciousNpc = Npc{}.New("Mark", Attack{})
		viciousNpc.Trigger()
	case "Tom":
		var friendlyNpc = Npc{}.New("Tom", Talk{})
		friendlyNpc.Trigger()
	case "Tanner":
		var silentNpc = Npc{}.New("Tanner", Walk{})
		silentNpc.Trigger()
	}
}
