package main

import (
  "golang.org/x/tour/tree"
	"fmt"
)

func main() {
  var i interface{} = "hello"
  
  s := i.(string)
  fmt.Println(s)
  
  f, ok := i.(float64)
  fmt.Println(f, ok)

  ch := make(chan int)
	go Walk(tree.New(1), ch) 
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1))) 
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkInner(t.Left, ch)
	ch <- t.Value
	WalkInner(t.Right, ch)
	close(ch)
}

func WalkInner(t *tree.Tree, ch chan int) {
	if (t != nil) {
		WalkInner(t.Left, ch)
		ch <- t.Value
		WalkInner(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1) 
	ch2 := make(chan int)
	go Walk(t2, ch2) 
	for i := 0; i < 10; i++ {
		tmp1 := <- ch1
		tmp2 := <- ch2
		if (tmp1 != tmp2) {
			return false
		}
	}
	return true
}
