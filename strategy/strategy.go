package strategy

func TriggerDucks() {
	var shortWingsDuck = Duck{}.New(Quack{}, FlyNoWay{})
	shortWingsDuck.PerformFly()
	shortWingsDuck.PerformQuack()

	var weirdDuck = Duck{}.New(Squeak{}, FlyWithWings{})
	weirdDuck.PerformFly()
	weirdDuck.PerformQuack()
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
		var notGivingAFuckNpc = Npc{}.New("Tanner", Walk{})
		notGivingAFuckNpc.Trigger()
	}
}
