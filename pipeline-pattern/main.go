package main

import "fmt"

// 管道，数据处理逻辑1
func square(data []int) <- chan int {
	out := make(chan int)
	go func(){
		defer close(out)
		for _, v := range data {
			out <- v * v
		}
	}()
	return out
}

//管道，数据处理逻辑2
func double(in <- chan int) <- chan int {
	out := make(chan int)
	go func(){
		defer close(out)
		for v:= range in {
			out <- v * 2
		}
	}()
	return out
}

// 管道，数据处理顺序
func pipeline(data []int) <- chan int{
	squareChan := square(data)
	doubleChan := double(squareChan)
	return doubleChan
}

func main(){
	data := []int{1,2,3,4,5,6,7,8,9,10}
	result := pipeline(data)

	for res := range result{
		fmt.Println(res)
	}
}