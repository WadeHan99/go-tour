package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("value = %d\n", value)
	}
	for i, _ := range pow {
		fmt.Printf("pow[%d] = %d\n", i, pow[i])
	}
}

