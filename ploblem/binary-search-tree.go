package ploblem

import (
	"fmt"
	"sort"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(v int) {
	if v < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: v}
		} else {
			n.Left.Insert(v)
		}
	} else if v > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			n.Right.Insert(v)
		}
	}
}

func (n *Node) Search(v int) bool {
	if n == nil {
		return false
	}

	fmt.Printf("กำลังเช็กที่ Node: %d\n", n.Value)

	if v == n.Value {
		return true
	}

	if v < n.Value {
		return n.Left.Search(v)
	}

	return n.Right.Search(v)
}

func BinarySearchTree(data []int, target int) bool {
	sort.Ints(data)
	root := &Node{Value: data[0]}
	for i := 1; i < len(data); i++ {
		root.Insert(data[i])
	}
	return root.Search(target)
}
