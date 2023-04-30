package main

import "fmt"

type Node struct {
	Value string
	Children []*Node
}
func(n *Node) AddChild(child *Node){
	n.Children =  append(n.Children,child)
}
func(n *Node) Print(indent string){
	fmt.Println(indent + n.Value)
	for _,child :=range n.Children{
		child.Print(indent+"    ")
	}
}
func main(){
	root:=&Node{Value:"Root"}
	//添加一级子节点
	child1 := &Node{Value: "child 1"}
	child2 := &Node{Value: "child 2"}
	root.AddChild(child1)
	root.AddChild(child2)

	//添加二级子节点
	grandchild1 := &Node{Value: "Grandchild 1"}
	grandchild2 := &Node{Value: "Grandchild 2"}
	child1.AddChild(grandchild1)
	child2.AddChild(grandchild2)

	//打印树结构
	root.Print("")
}