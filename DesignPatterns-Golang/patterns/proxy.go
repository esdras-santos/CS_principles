//VIRTUAL PROXY

package patterns

import (
	"fmt"
	"sync"
)

var onceproxy sync.Once

type ISubject interface{
	request() string
}

type RealSubject struct{
	text string
}

func (rs *RealSubject) request() string{
	return rs.text
}

type Proxy struct{
	rs RealSubject
}

func (p *Proxy) request() string{
	once.Do(func() {
		p.rs = RealSubject{}
		p.rs.text = "proxy access!"
	})
	return p.rs.text
}

func ProxyPattern(){
	p := Proxy{}
	fmt.Println(p.request())
}