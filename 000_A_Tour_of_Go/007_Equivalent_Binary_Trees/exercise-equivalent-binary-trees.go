package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	fmt.Println("Walking")
	//https://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
	// Create a stack - a slice of integers here
	stack := []*tree.Tree{}

	curr := t
	for curr != nil || len(stack) > 0 {

		// Reach the left most Node of the
		// curr Node
		for curr != nil {

			// Place pointer to a tree node on
			// the stack before traversing
			// the node's left subtree
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Current must be NULL at this point
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ch <- curr.Value

		// we have visited the node and its
		// left subtree.  Now, it's right
		// subtree's turn
		curr = curr.Right

	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1channel := make(chan int)
	t2channel := make(chan int)
	go Walk(t1, t1channel)
	go Walk(t2, t2channel)
	t1slice := make([]int, 10)
	t2slice := make([]int, 10)

	for i := 0; i < 10; i++ {
		t1slice[i] = <-t1channel
	}

	for i := 0; i < 10; i++ {
		t2slice[i] = <-t2channel
	}

	for i := 0; i < 10; i++ {
		if t1slice[i] != t2slice[i] {
			return false
		}
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
