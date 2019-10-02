package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value  int
	Data   string
	Left   *Node
	Right  *Node
	parent *Node
}

func (n *Node) Insert(value int, data string) error {
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

func (n *Node) Find(value int) (*Node, bool) {
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

func (n *Node) Delete(value int) error {
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

func (t *Tree) Insert(value int, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s int) (*Node, bool) {
	if t.Root == nil {
		return nil, false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s int) error {

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

func print(n *Node, level int) {
	if n != nil {
		format := ""

		for i := 0; i < level; i++ {
			format += "       "
		}

		level++

		print(n.Left, level)
		fmt.Printf(format+"%d [\n", n.Value)
		print(n.Right, level)
	}
}

func main() {
	// n := &Node{Value: 5, Data: "123"}

	// fmt.Println(n)

	// n.Insert(2, "321")

	// fmt.Println(n)

	// fmt.Println(n.Find(2))

	// n.Delete(2)

	// fmt.Println(n)

	t := &Tree{}
	t.Insert(10, "1")
	t.Insert(5, "2")
	t.Insert(15, "3")
	t.Insert(3, "3")
	t.Insert(7, "3")
	t.Insert(6, "3")
	t.Insert(2, "3")
	t.Insert(1, "3")
	t.Insert(0, "3")

	print(t.Root, 0)

}
