/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package ds

import (
	"fmt"
)

// BST Node
type Node struct {
	data   int
	count int
	left  *Node
	right *Node
}


func NewBinarySearchTree(data int) *Node {
	return &Node{data, 1, nil, nil}
}

func (node *Node)Insert(data int) *Node {
	return _insert(node,data)
}

func _insert(root *Node, data int) *Node{
	if root == nil {
		return NewBinarySearchTree(data)
	}
	if root.data == data { root.count++ } else
	if root.data > data { root.left = _insert(root.left,data) } else
	if root.data < data { root.right = _insert(root.right,data) }
	return root
}

func (node *Node) Search(data int) *Node {
	return _search(node,data)
}

func _search(node *Node, data int) *Node  {
	if node == nil || node.data == data {
		return node
	}

	if node.data > data {
		return _search(node.left,data)
	}

	return _search(node.right,data)
}


func (node *Node) Delete(data int) *Node {
	return _delete(node,data)
}

func _delete(node *Node, data int) *Node {
	if node == nil {
		return node
	}

	if node.data > data {
		node.left = _delete(node.left, data)
	} else if node.data < data {
		node.right = _delete(node.right, data)
	} else {

		// process count
		if node.count > 1 {
			node.count--
			return node
		}

		// delete current node
		if node.left == nil { return node.right }
		if node.right == nil { return node.left }
		n := _minNode(node.right)
		node.data = n.data
		node.right = _delete(node.right,n.data)
	}

	return node
}


func _minNode(node *Node) *Node {
	for node.left != nil {
		node = maxNode(node.left)
	}
	return node
}

func maxNode(node *Node) *Node {
	for node.right != nil {
		node = maxNode(node.right)
	}
	return node
}



func InOrder(node *Node) {
	if node != nil {
		InOrder(node.left)
		fmt.Printf("Node{%d count:%d}\t", node.data, node.count)
		InOrder(node.right)
	}
}


func PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("Node{%d count:%d}\t", node.data, node.count)
		PreOrder(node.left)
		PreOrder(node.right)
	}
}
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.left)
		PostOrder(node.right)
		fmt.Printf("Node{%d count:%d}\t", node.data, node.count)
	}
}


func (node *Node) String() string {
	return fmt.Sprintf("Node{data:%d,count:%d,left:%+v,right:%+v}", node.data, node.count, node.left, node.right)
}
