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

func main() {
	n := &node{value: 1, left: nil, right: nil}
	n.left = &node{value: 2, left: nil, right: nil}
	n.left.left = &node{value: 8, left: nil, right: nil}
	n.right = &node{value: 3, left: nil, right: nil}

	b := &bst{
		root: n,
		len:  3,
	}
	fmt.Println(b)
}
