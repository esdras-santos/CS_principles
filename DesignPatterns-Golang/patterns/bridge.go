package patterns

import "fmt"

type InterfaceClass interface {
	DogName() string
}

type AbstractClass interface {
	Show()
}

//////////////////////////////////////////////////////////////////////////////////
//IMPLEMENTATION
//////////////////////////////////////////////////////////////////////////////////

//abstract class(struct) implementation
type View struct {
	IC   InterfaceClass
	
}

func (v *View) Show() string{
	return "you have a dog named " + v.IC.DogName() + " and you'll go the beach, do you take " + v.IC.DogName() + " or leave " + v.IC.DogName() + "?"
}


//class(struct) that implements the InterfaceClass
type Resource struct {
	dogName string
}

//interface method(function)
func (r Resource) DogName() string {
	return r.dogName
}


func Bridge() {
	resource := Resource{"nabunda"}
	view := View{resource}
	fmt.Println(view.Show())
}