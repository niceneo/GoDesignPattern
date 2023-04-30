package main

import "fmt"

//适配器源
type Adaptee interface {
	SpecificRequest() string
}
type AdapteeImpl struct{}
func (a *AdapteeImpl) SpecificRequest() string{
	return "Specific Request"
}
//适配器目标
type Target interface {
	Request() string
}
//适配器, 适配器结构体继承适配器源,并且实现适配器目标接口方法
type Adapter struct {
	Adaptee
}
func(a *Adapter) Request() string{
	return a.SpecificRequest()
}
//适配器实现
func Client(t Target) string{
	return t.Request()
}

func main(){
	adaptee := &AdapteeImpl{}
	target := &Adapter{Adaptee:adaptee}
	output := Client(target)
	fmt.Println(output)
}
