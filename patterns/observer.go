package patterns

import (
	"fmt"
	"time"
)



type IObservable interface {
	Add(IObserver)
	Remove(IObserver)
	Notify(time.Time)
}

//implementation of the IObservable interface
type Observable struct {
	Observers []IObserver
	Date time.Time
}

func (obs *Observable) Add(iobs IObserver) {
	obs.Observers = append(obs.Observers, iobs)
}

func (obs *Observable) Remove(iobs IObserver) {
	var j int
	for i, e := range obs.Observers {
		if iobs == e {
			j = i
			break
		}
	}
	obs.Observers = append(obs.Observers[:j], obs.Observers[j+1:]...)
}

func (obs *Observable) Notify() {
	for _, observer := range obs.Observers {
		observer.Update()
	}
}

func (obs *Observable) SetState() {
	obs.Date = time.Now()
}

func (obs *Observable) GetState() time.Time{
	return obs.Date
}


type IObserver interface {
	Update()
}

//implementation of the IObserver interface
type Obser struct {
	//instance to comunicate with the observable
	Obsvbl *Observable
	
}

func (obs *Obser) Update() {
	//get the state of the observable
	fmt.Printf("\r%s",obs.Obsvbl.GetState().Format("01-02-2006 15:04:05 Mon"))
}

func Observer(){
	obs := Observable{}

	//subscriber
	sub := Obser{}
	sub.Obsvbl = &obs
	obs.Add(&sub)

	for {
		time.Sleep(time.Second)
		obs.SetState()
		//notify the subscribers for each 1 second
		obs.Notify()
	}
}