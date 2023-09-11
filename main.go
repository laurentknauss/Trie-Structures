package main

import "fmt"

// Node represents a node in the binary search tree
type Node struct {
	Value int  // The value stored in the node
	Left  *Node // Pointer to the left child node
	Right *Node // Pointer to the right child node
}

// Insert inserts a new node into the binary search tree
func (n *Node) Insert(value int) {
	if n == nil { // Check if the node is nil
		return
	}
	if value < n.Value { // Check if the value should go to the left subtree
		if n.Left == nil { // If left child is nil, create a new node
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value) // Recursively insert into the left subtree
		}
	} else { // The value should go to the right subtree
		if n.Right == nil { // If right child is nil, create a new node
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value) // Recursively insert into the right subtree
		}
	}
}

// InOrderTraverse traverses the tree in in-order
func (n *Node) InOrderTraverse(f func(int)) {
	if n != nil { // Check if the node is nil
		n.Left.InOrderTraverse(f)  // Recursively traverse the left subtree
		f(n.Value)                 // Apply the function to the current node's value
		n.Right.InOrderTraverse(f) // Recursively traverse the right subtree
	}
}

func main() {
	root := &Node{Value: 10} // Create the root node with value 10

	// Insert nodes based on the structure you provided
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	// Traverse and print the tree in in-order
	root.InOrderTraverse(func(value int) {
		fmt.Println(value)
	})
}
