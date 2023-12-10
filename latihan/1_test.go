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

func (b *bst) remove(data int) {
	b.removeByNode(data, b.root)
}

func (b *bst) removeByNode(data int, root *node) *node {
	if root == nil {
		return nil
	}

	if data < root.data {
		root.left = b.removeByNode(data, root.left)
	} else if data > root.data {
		root.right = b.removeByNode(data, root.right)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.data = temp.data
			root.left = b.removeByNode(temp.data, root.left)
		}
	}
	return root
}

func (b *bst) search(data int) bool {
	return b.searchByNode(data, b.root)
}
func (b *bst) searchByNode(data int, root *node) bool {
	if root == nil {
		return false
	}
	if data == root.data {
		return true
	} else if data < root.data {
		return b.searchByNode(data, root.left)
	} else {
		return b.searchByNode(data, root.right)
	}
}

func TestLatihan1(t *testing.T) {
	b := &bst{
		root:   nil,
		length: 5,
	}
	b.add(24)
	b.add(12)
	b.add(36)
	b.add(7)
	b.add(15)
	b.add(5)
	b.add(9)
	b.add(11)
	b.print()
	fmt.Println()
	b.remove(12)
	b.print()
	fmt.Println(b.search(30))
}
