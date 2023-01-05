package main

import "fmt"

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4]) // 从这里 panic 跳转 defer 内逻辑
	fmt.Println("after panic")
	fmt.Println(arr[5])
	fmt.Println("after panic 2")
}

func main() {
	test_recover()
	fmt.Println("after recover")
}
