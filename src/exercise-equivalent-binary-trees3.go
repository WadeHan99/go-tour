//

package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func walkImpl(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkImpl(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkImpl(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channle ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	// ch := make(chan int)
	// go func() {
	// 	Walk(tree.New(1), ch)
	// 	ch <- 0
	// }()

	// for {
	// 	t := <-ch
	// 	if t == 0 {
	// 		break
	// 	}
	// 	fmt.Println(t)
	// }

	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for n := range ch {
		fmt.Printf("%v", n)
		fmt.Println()
	}

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
