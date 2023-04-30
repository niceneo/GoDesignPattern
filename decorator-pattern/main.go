package main

import "fmt"

// 接口，定义了一个操作
type Component interface {
	Operation() string
}

//具体组件，实现了Component接口
type ConcreteComponent struct {}
func(c *ConcreteComponent) Operation() string{
	return "ConcreteComponent"
}

type Decorator func(component Component) Component

type decoratorA struct {
	component Component
}

func (d *decoratorA) Operation() string{
	return fmt.Sprintf("ConcreteDecoratorA(%s)",d.component.Operation())
}

// 具体装饰器,给Compoent 添加新的行为
func ConcreteDecoratorA(component Component) Component{
	return &decoratorA{component:component}
}

// 具体另外一个装饰器，给Component添加新的行为
type decoratorB struct {
	component Component
}

func(d *decoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)",d.component.Operation())
}

func ConcreteDecoratorB(component Component) Component{
	return &decoratorB{component:component}
}

func main(){
	component := &ConcreteComponent{}

	decoratedComponentA := ConcreteDecoratorA(component)
	decoratedComponentB := ConcreteDecoratorB(decoratedComponentA)

	fmt.Println(component.Operation())
	fmt.Println(decoratedComponentA.Operation())
	fmt.Println(decoratedComponentB.Operation())
}