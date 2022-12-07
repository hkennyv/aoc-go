package utils

import (
	"errors"
	"fmt"
)

type Node struct {
	Name     string
	Size     int
	Parent   *Node
	Children []*Node
}

func (n *Node) PrintNode(i int) {

	if i > 10 {
		return
	}

	s := ""

	for j := 0; j < i; j++ {
		s += "  "
	}

	fmt.Printf(s + n.Name + " (" + fmt.Sprint(n.Size) + ")" + "\n")
	for _, c := range n.Children {
		c.PrintNode(i + 1)
	}
}

func (n *Node) AddChild(name string, size int) (*Node, error) {
	for _, c := range n.Children {
		if c.Name == name {
			return c, errors.New(name + " - child with that name already exists")
		}
	}

	new := &Node{
		Name:     name,
		Size:     size,
		Parent:   n,
		Children: []*Node{},
	}

	n.Children = append(n.Children, new)
	return new, nil
}

func (n *Node) CalculateSize() int {
	sum := n.Size

	for _, c := range n.Children {

		size := c.CalculateSize()
		if c.Size != size {
			c.Size = size
		}

		sum += size
	}

	if n.Size != sum {
		n.Size = sum
	}

	return sum
}
