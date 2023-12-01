package main

import (
	"fmt"
	"log"

	// "log"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}
func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	// log.Printf("inOrderTraversalByNode(sb,root(%s))", root)
	if root == nil {
		return
	}
	// log.Printf("inOrderTraversalByNode(sb,root.left(%s))", root.left)
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	// log.Printf("inOrderTraversalByNode(sb,root.right(%s))", root.right)
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}
func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}
func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b *bst) remove(value int) {
	b.root = b.removeByNode(b.root, value)
}
func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}
	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.value = temp.value
			log.Println("Value temp", temp.value)
			root.left = b.removeByNode(root.left, temp.value)
		}
	}
	return root
}

func main() {
	// n := &node{value: 1, left: nil, right: nil}
	// n.left = &node{value: 0, left: nil, right: nil}
	// n.right = &node{value: 2, left: nil, right: nil}

	// b := &bst{
	// 	root: n,
	// 	len:  0,
	// }
	// b.add(5)
	// b.add(4)
	// b.add(6)
	// fmt.Println(b)
	// b.remove(0)
	// fmt.Println(b)
	bst := &bst{}
	bst.add(5)
	bst.add(3)
	bst.add(7)
	bst.add(2)
	bst.add(4)
	bst.add(6)
	bst.add(8)

	fmt.Println(bst)
	bst.remove(3)
	fmt.Println(bst)
}
