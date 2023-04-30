package main

import "fmt"

//抽象 - 形状
type Shape interface {
	Draw()
}

//实现部分 - 颜色
type Color interface {
	ApplyColor() string
}

type RedColor struct{}
func(r *RedColor) ApplyColor() string {
	return "Red"
}

type BlueColor struct {}
func(b *BlueColor) ApplyColor() string{
	return "Blue"
}

//桥接部分 - 形状和颜色的组合
type Circle struct {
	color Color
}
func(c *Circle) Draw(){
	fmt.Printf("Drawing a %s circle\n",c.color.ApplyColor())
}

type Square struct {
	Color
}
func(s *Square) Draw(){
	fmt.Printf("Drawing a %s square",s.ApplyColor())
}

func main(){
	red := &RedColor{}
	blue := &BlueColor{}

	circleRed := &Circle{color: red}
	circleBlue := &Circle{color: blue}

	squareRed := &Square{red}
	squareBlue := &Square{blue}

	circleRed.Draw()
	circleBlue.Draw()

	squareRed.Draw()
	squareBlue.Draw()
}