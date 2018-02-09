package main

import "golang.org/x/tour/tree"
import "fmt"


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walking(t, ch)
	close(ch)
}

func Walking(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walking(t.Left, ch)
	ch <- t.Value
	Walking(t.Right, ch)	
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1_ch := make(chan int)
	t2_ch := make(chan int)
	
	go Walk(t1, t1_ch)
	go Walk(t2, t2_ch)
	for {
		t1_val, ok1 := <-t1_ch
		t2_val, ok2 := <-t2_ch
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if t1_val != t2_val {
			return false
		}
	}
}


func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	
	for i := range ch {
		fmt.Println(i)
	}
	
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

