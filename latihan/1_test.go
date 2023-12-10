package latihan

import (
	"fmt"
	"testing"
)

type node struct {
	data  int
	left  *node
	right *node
}

type bst struct {
	root   *node
	length int
}

func (b bst) print() {
	b.printByNode(b.root)
}
func (b bst) printByNode(root *node) {
	if root == nil {
		return
	}
	b.printByNode(root.left)
	fmt.Printf("%d ", root.data)
	b.printByNode(root.right)
}

func (b *bst) add(data int) {
	b.root = b.addByNode(data, b.root)
}
func (b *bst) addByNode(data int, root *node) *node {
	if root == nil {
		return &node{data: data, left: nil, right: nil}
	} else if data < root.data {
		root.left = b.addByNode(data, root.left)
	} else {
		root.right = b.addByNode(data, root.right)
	}
	return root
}

func TestLatihan1(t *testing.T) {
	// n := &node{data: 5, left: nil, right: nil}
	// n.left = &node{data: 3, left: nil, right: nil}
	// n.left.left = &node{data: 2, left: nil, right: nil}
	// n.left.right = &node{data: 4, left: nil, right: nil}
	// n.right = &node{data: 6, left: nil, right: nil}

	b := &bst{
		root:   nil,
		length: 5,
	}
	b.add(6)
	b.add(1)
	b.add(2)
	b.add(-100)
	b.print()
}
