package main

import "fmt"

type Product struct {
	PartA string
	PartB string
	PartC string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

type ConcreateBuilder struct {
	product *Product
}

func (b *ConcreateBuilder) BuildPartA(){
	b.product.PartA ="Part A"
}
func (b *ConcreateBuilder) BuildPartB(){
	b.product.PartB = "Part B"
}
func (b *ConcreateBuilder) BuildPartC(){
	b.product.PartC = "Part C"
}
func (b *ConcreateBuilder) GetProduct() *Product{
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder){
	d.builder = builder
}

func(d *Director) Construct(){
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main(){
	concreateBuilder := &ConcreateBuilder{product: &Product{}}
	director := &Director{builder: concreateBuilder}
	director.Construct()
	product := concreateBuilder.GetProduct()
	fmt.Println(product.PartA,product.PartB,product.PartC)
}