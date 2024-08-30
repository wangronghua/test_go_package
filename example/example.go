package example

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multi(a, b int) int {
	return a * b
}

func init() {
	fmt.Println("测试包已初始化2")
}
