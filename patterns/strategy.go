package patterns

import "fmt"

type IFlyBeahavior interface {
	Fly()
}

type IQuackBehavior interface {
	Quack()
}

type Duck struct {
	fb IFlyBeahavior
	qb IQuackBehavior
}

func (d *Duck) Fly() {
	d.fb.Fly()
}

func (d *Duck) Quack() {
	d.qb.Quack()
}

//////////////////////////////////////////////////////////////////////////////////
//IMPLEMENTATION
//////////////////////////////////////////////////////////////////////////////////

type SimpleDuck struct {
	Duck
}

type WildDuck struct {
	Duck
}

type SimpleFly struct{}

func (sf SimpleFly) Fly(){
	fmt.Println("ugly fly")
}

type WildFly struct{}

func (wf WildFly) Fly(){
	fmt.Println("Flying like a super cool rocket")
}

type SimpleQuack struct{}

func (sq SimpleQuack) Quack() {
	fmt.Println("quack")
}

type WildQuack struct{}

func (wq WildQuack) Quack() {
	fmt.Println("QUACK!!!!!!!!!!!")
}

func Strategy(){
	sd := SimpleDuck{}
	wd := WildDuck{}
	
	sd.fb = SimpleFly{}
	sd.qb = SimpleQuack{}

	wd.fb = WildFly{}
	wd.qb = WildQuack{}

	sd.Quack()
	sd.Fly()

	wd.Quack()
	wd.Fly()
}

