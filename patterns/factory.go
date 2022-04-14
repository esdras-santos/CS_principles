package patterns

import "fmt"

type animal struct {
	name  string
	sound string
}

func (a animal) MakeSound() string {
	return a.sound
}

func AnimalFactory(name string) animal {
	animal := animal{}
	animal.name = name
	if name == "dog" {
		animal.sound = "AUAU"
	} else if name == "cat" {
		animal.sound = "MIAU"
	}
	return animal
}

func Factory() {
	cat := AnimalFactory("cat")
	fmt.Println(cat.MakeSound())
	dog := AnimalFactory("dog")
	fmt.Println(dog.MakeSound())
}
