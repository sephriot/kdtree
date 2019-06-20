package kdtree

import (
	"fmt"
)

type KDTree struct {
	root *node
}

type node struct {
	Point
	Left *node
	Right *node
}

func New() *KDTree {
	return &KDTree{nil}
}

func (k *KDTree) String() string {
	return fmt.Sprintf("[%s]", Print(k.root))
}

func Print(n *node) string {
	if n != nil && (n.Left != nil || n.Right != nil) {
		return fmt.Sprintf("[%s %s %s]", Print(n.Left), n.Point, Print(n.Right))
	}
	if n == nil {
		return fmt.Sprint( nil)
	}
	return fmt.Sprintf("%s", n.Point)
}

func (k *KDTree) Add(p Point) {
	if k.root != nil {
		k.root.Add(p,0)
	} else {
		k.root = &node{p, nil, nil}
	}
}

func (k *KDTree) Remove(p Point) {
	if k.root != nil {
		if !k.root.Remove(p, 0, nil) {
			k.root = nil
		}
	}
}

func (n *node) Add(p Point, d int) {
	if p.Dimension(d) < n.Point.Dimension(d) {
		if n.Left == nil {
			n.Left = &node{p, nil, nil}
		} else {
			n.Left.Add(p, (d+1)%n.Dimensions())
		}
	} else {
		if n.Right == nil {
			n.Right = &node{p, nil, nil}
		} else {
			n.Right.Add(p, (d+1)%n.Dimensions())
		}
	}
}

func (n *node) Find(p Point, d int, parent *node) (*node, *node, int) {
	d = d % n.Dimensions()
	if equals(n, p) {
		return n, parent, d
	}
	if p.Dimension(d) < n.Dimension(d) {
		if n.Left == nil {
			return nil, parent, d
		}
		return n.Left.Find(p, d+1, n)
	}
	if n.Right == nil {
		return nil, parent, d
	}
	return n.Right.Find(p, d+1, n)
}

// returns false if node to be deleted is root
func (n *node) Remove(p Point, d int, parent *node) bool {
	cd := d % n.Dimensions()
	found, parent, foundD := n.Find(p, cd, parent)

	if found == nil {
		return true
	}

	if found.Right != nil {
		sub := found.Right.FindMin(foundD, foundD+1)
		found.Point = sub.Point
		return found.Right.Remove(sub.Point, foundD+1, found)
	}

	if found.Left != nil {
		sub := found.Left.FindMin(foundD, foundD+1)
		found.Point = sub.Point
		return found.Left.Remove(sub.Point, foundD+1, found)
	}

	if parent != nil {
		if parent.Left == found {
			parent.Left = nil
		} else if parent.Right == found {
			parent.Right = nil
		}
		return true
	}
	return false
}

func (n *node) FindMin(d, depth int) *node {
	cd := depth % n.Dimensions()
	if d == cd {
		if n.Left == nil {
			return n
		}
		return n.Left.FindMin(d, depth+1)
	}
	var left, right *node
	if n.Left != nil {
		left = n.Left.FindMin(d, depth+1)
	}
	if n.Right != nil {
		right = n.Right.FindMin(d, depth+1)
	}

	return n.Min(left, d).Min(right, d)
}

func (n *node) Min(nn *node, d int) *node {
	if nn == nil || n.Dimension(d) < nn.Dimension(d) {
		return n
	}
	return nn
}