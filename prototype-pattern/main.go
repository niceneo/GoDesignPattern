package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcreatePrototype struct {
	name string
	age int
}

func (c *ConcreatePrototype) Clone() Prototype{
	return &ConcreatePrototype{
		name: c.name,
		age: c.age,
	}
}

func NewPrototype(name string,age int) Prototype{
	return &ConcreatePrototype{
		name:name,
		age: age,
	}
}
func main(){
	prototype := NewPrototype("NEO",33)
	clone1 := prototype.Clone()
	clone2 := prototype.Clone()

	fmt.Println(prototype)
	fmt.Println(clone1)
	fmt.Println(clone2)
}
