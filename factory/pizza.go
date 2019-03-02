package factory

type PizzaSort int
const (
	Cheese PizzaSort = 1 << iota
	Pepperoni
	Clam
)

type PizzaProvider interface {
	GetPizza() *Pizza
}
type Pizza struct {
	title string
	rating  int
}
func (Pizza) New(title string, rating int) *Pizza {
	p := new(Pizza)
	p.title = title
	p.rating = rating
	return p
}
func (p *Pizza) GetPizza() *Pizza {
	return p
}

type CheesePizza struct {
	Pizza
}
type PepperoniPizza struct {
	Pizza
}
type ClamPizza struct {
	Pizza
}

func CreatePizza(sort PizzaSort) *PizzaProvider {
	var pizza PizzaProvider
	if sort == Cheese {
		pizza = CheesePizza{}.New("Cheesy Circle", 4)
	} else if sort == Pepperoni {
		pizza = PepperoniPizza{}.New("Pepperoni Surprise", 3)
	} else if sort == Clam {
		pizza = ClamPizza{}.New("Clamsy Chef", 5)
	}
	return &pizza
}
