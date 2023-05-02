package main

import (
	"fmt"
)

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
	name string
}

func (c *ConcreteObserver) Update(msg string){
	fmt.Printf("%s, recevied: %s\n",c.name,msg)
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Notify(string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject)Register(observer Observer){
	s.observers = append(s.observers,observer)
}

func (s *ConcreteSubject)Unregister(observer Observer){
	for i,obs := range s.observers{
		if obs == observer{
			s.observers = append(s.observers[:i],s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject)Notify(msg string){
	for _, observer:=range s.observers{
		observer.Update(msg)
	}
}

func main(){
	subject := &ConcreteSubject{}
	ob1:=&ConcreteObserver{name: "ob1"}
	ob2:=&ConcreteObserver{name: "ob2"}
	subject.Register(ob1)
	subject.Register(ob2)

	subject.Notify("Hello Observer")

	subject.Unregister(ob1)

	subject.Notify("Another message")

}