package observer

// Observer Pattern defines one-to-many dependency between objects
// (to be specific - between Subject or Observable and Observers)
// so that when one object changes state, all its dependents are notified and updated automatically
type Observer interface {
	Update(temp, humidity, pressure float32)
}

type Observable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type Observers []Observer
func (slice Observers) indexOf(observer Observer) int {
	for pos, o := range slice {
		if o == observer {
			return pos
		}
	}
	return -1
}
func (slice Observers) remove(pos int) Observers {
	return append(slice[:pos], slice[pos+1:]...)
}
