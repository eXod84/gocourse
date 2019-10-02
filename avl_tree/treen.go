package main

import (
	"fmt"
)

type Node struct {
	Value  int
	height int
	Left   *Node
	Right  *Node
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *Node) bFactor() int {
	if n == nil {
		panic("Nil pointer")
	}

	return n.Right.Height() - n.Left.Height()
}

func (n *Node) fixHeight() {
	if n == nil {
		panic("Nil pointer")
	}

	hl := n.Left.Height() + 1
	hr := n.Right.Height() + 1

	if hl > hr {
		n.height = hl
	} else {
		n.height = hr
	}
}

func (p *Node) rotateRight() *Node {
	q := p.Left
	p.Left = q.Right
	q.Right = p

	p.fixHeight()
	q.fixHeight()

	return q
}

func (q *Node) rotateLeft() *Node {
	p := q.Right
	q.Right = p.Left
	p.Left = q

	p.fixHeight()
	q.fixHeight()

	return p
}

func (p *Node) balance() *Node {
	p.fixHeight()

	if p.bFactor() == 2 {
		if p.Right.bFactor() < 0 {
			p.Right = p.Right.rotateRight()
		}
		return p.rotateLeft()
	}

	if p.bFactor() == -2 {
		if p.Left.bFactor() > 0 {
			p.Left = p.Left.rotateLeft()
		}
		return p.rotateRight()
	}

	return p // балансировка не нужна
}

func (p *Node) Insert(v int) *Node {
	if p == nil {
		return &Node{Value: v, height: 1}
	}
	if v < p.Value {
		p.Left = p.Left.Insert(v)
	} else {
		p.Right = p.Right.Insert(v)
	}
	return p.balance()
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

func print(n *Node, level int) {
	if n != nil {
		format := ""

		for i := 0; i < level; i++ {
			format += "       "
		}

		level++

		print(n.Right, level)
		fmt.Printf(format+"%d|%d [\n", n.Value, n.height)
		print(n.Left, level)
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value, height: 1}
		return
	}
	t.Root = t.Root.Insert(value)
}

func (t *Tree) Print() {
	print(t.Root, 0)
}

func main() {
	t := &Tree{}

	t.Insert(10)
	t.Insert(9)
	t.Insert(8)
	t.Insert(7)

	t.Print()

	fmt.Print("\n\n=================================\n\n")

	t2 := &Tree{}

	t2.Insert(1)
	t2.Insert(2)
	t2.Insert(3)
	t2.Insert(4)
	t2.Insert(5)
	t2.Insert(6)
	t2.Insert(7)
	t2.Insert(8)
	t2.Insert(9)

	t2.Print()

}
