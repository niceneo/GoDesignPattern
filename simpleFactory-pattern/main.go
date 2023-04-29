package main

import "fmt"

// Product 接口定义
type Product interface {
	GetName() string
}

//ProductA A具体类
type ProductA struct {}
func (p *ProductA) GetName() string{
	return "Product A"
}

//ProductB B具体类
type ProductB struct {}
func (p *ProductB) GetName() string {
	return "Product B"
}

type Factory struct {}
func (f *Factory) CreateProduct(name string) Product{
	switch name {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	default:
		return nil
	}
}

func main(){
	factory := &Factory{}
	pA:=factory.CreateProduct("A")
	pB:=factory.CreateProduct("B")
	fmt.Println(pA.GetName())
	fmt.Println(pB.GetName())
}