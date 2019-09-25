package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value  string
	Data   string
	Left   *Node
	Right  *Node
	parent *Node
}

func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("Must be not empty")
	}

	if n.Value == value {
		return nil
	}

	if n.Value < value {
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data, parent: n}
			return nil
		}

		return n.Left.Insert(value, data)
	}

	if n.Value > value {
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data, parent: n}
			return nil
		}

		return n.Right.Insert(value, data)
	}

	return nil
}

func (n *Node) Find(value string) (*Node, bool) {
	if n.Value == value {
		return n, true
	}

	if n.Value < value {
		return n.Left.Find(value)
	}

	if n.Value > value {
		return n.Right.Find(value)
	}

	return nil, false
}

func (n *Node) findMax() *Node {
	if n.Right == nil {
		return n
	}

	return n.Right.findMax()
}

func (n *Node) Replace(replacement *Node) {
	if n == n.parent.Left {
		n.parent.Left = replacement
	} else {
		n.parent.Right = replacement
	}
}

func (n *Node) Delete(value string) error {
	if n == nil {
		return errors.New("Empty")
	}

	removeNode, ok := n.Find(value)
	if !ok {
		return errors.New("Can't find node")
	}

	if removeNode.Left == nil && removeNode.Right == nil {
		removeNode.Replace(nil)
		return nil
	}

	if removeNode.Left == nil {
		removeNode.Replace(removeNode.Right)
		return nil
	}

	if removeNode.Right == nil {
		removeNode.Replace(removeNode.Left)
		return nil
	}

	maxNode := removeNode.Left.findMax()

	n.Value = maxNode.Value
	n.Data = maxNode.Data

	maxNode.Delete(maxNode.Value)

	return nil
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) (*Node, bool) {
	if t.Root == nil {
		return nil, false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s string) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	err := t.Root.Delete(s)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {
	n := &Node{Value: "f", Data: "123"}

	fmt.Println(n)

	n.Insert("a", "321")

	fmt.Println(n)

	fmt.Println(n.Find("a"))

	n.Delete("a")

	fmt.Println(n)

	t := &Tree{}

	t.Insert("f", "1")
	t.Insert("a", "2")
	t.Insert("z", "3")

	t.Traverse(t.Root, func(n *Node) { fmt.Println(n.Value, ": ", n.Data, " | ") })
}
