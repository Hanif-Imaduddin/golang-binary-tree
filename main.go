package main

import (
	"fmt"
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

func main() {
	b := &bst{
		root: nil,
		len:  0,
	}
	b.add(65)
	b.add(32)
	b.add(17)
	b.add(14)
	fmt.Println(b)
	fmt.Println(b.len)
}
