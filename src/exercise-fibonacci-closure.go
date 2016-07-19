package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		first, second = second, first + second
		return first
	}
}

func main() {
	f := fibonacci()
	fmt.Println("f =", f);
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
