package patterns

import "fmt"

//Component interface
type TodoList interface{
	Operation() string
	//add()
	//remove()
}

//Composite
type Project struct{
	tl []TodoList
	title string
}

func (p *Project) Operation() string{
	tasks := "->"+p.title+"\n"
	for _,t := range p.tl{
		tasks += "	"+t.Operation()+"\n"
	}

	return tasks
}

//Leaf
type Task struct{
	task string
}

func (t *Task) Operation() string{
	return t.task
}


func CompositePattern(){
	
	t11 := Task{}
	t21 := Task{}
	t31 := Task{}

	t11.task = "  eat"
	t21.task = "  sleep"
	t31.task = "  repeat"

	p11 := Project{}
	p11.title = "evening tasks"
	p11.tl = append(p11.tl, &t11,&t21,&t31)
	
	t1 := Task{}
	t2 := Task{}
	t3 := Task{}

	t1.task = "wake up"
	t2.task = "eat"
	t3.task = "code"
	

	p1 := Project{}
	p1.title = "daily tasks"
	p1.tl = append(p1.tl, &t1,&t2,&t3,&p11)
	fmt.Println(p1.Operation())
}