package main

import "fmt"

//抽象产品 button
type Button interface {
	Click()
}

//抽象产品 checkbox
type CheckBox interface {
	Check()
}

//抽象工厂
type GUIFactory interface {
	CreateButton() Button
	CreateCheckBox() CheckBox
}

//具体产品
type WinButton struct {}
func (w *WinButton) Click(){
	fmt.Println("Windows button clicked")
}

type WinCheckBox struct {}
func (w *WinCheckBox) Check(){
	fmt.Println("Windows checkbox checked")
}

type MacButton struct{}
func (m *MacButton) Click(){
	fmt.Println("Mac Button clicked")
}

type MacCheckBox struct{}
func (m *MacCheckBox) Check(){
	fmt.Println("Mac CheckBox checked")
}

//具体工厂
type WinFactory struct{}
func(w *WinFactory) CreateButton() Button{
	return &WinButton{}
}
func (w *WinFactory) CreateCheckBox() CheckBox{
	return &WinCheckBox{}
}

type MacFactory struct{}
func (m *MacFactory) CreateButton() Button{
	return &MacButton{}
}
func (m *MacFactory) CreateCheckBox() CheckBox{
	return &MacCheckBox{}
}

func main(){
	var factory GUIFactory
	//根据需要选择具体工厂
	factory = &WinFactory{}

	button := factory.CreateButton()
	checkbox := factory.CreateCheckBox()

	button.Click()
	checkbox.Check()
}