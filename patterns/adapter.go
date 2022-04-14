package patterns

import "fmt"

type Client struct {
	target ITarget
}

func (c *Client) Init(){
	c.target = &Adapter{Adaptee{}}
	c.target.request()
}

type ITarget interface {
	request()
}

type Adapter struct {
	Adaptee Adaptee
}

func (a *Adapter) request() {
	a.Adaptee.SpecificRequest()
}

type Adaptee struct{}

func (a *Adaptee) SpecificRequest() {
	fmt.Println("adapted!")
}

func AdapterPattern(){
	c := Client{}
	c.Init()
}