package patterns

import "fmt"

type Beverage interface {
	Cost() int
}

type Expresso struct {
}

func (e *Expresso) Cost() int {
	return 4
}

//decorator
type AddonDecorator interface {
	/* with the cost function in the decorator interface we make that compatible 
	   with the Bevarage interface
	*/
	Cost() int
	Name() string
}

type Chocolate struct {
	beverage Beverage
}

func (c *Chocolate) Cost() int {
	return c.beverage.Cost()+4
}

func (c *Chocolate) Name() string {
	return "Chololate"
}

type Milk struct {
	beverage Beverage
}

func (m *Milk) Cost() int {
	return m.beverage.Cost()+4
}

func (m *Milk) Name() string {
	return "Milk"
}


func Decorator() {
	var c Chocolate
	ex := Expresso{}
	c.beverage = &ex

	fmt.Println(c.Cost())
	fmt.Println(c.Name())
}