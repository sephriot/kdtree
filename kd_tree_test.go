package kdtree

import (
	"github.com/sephriot/kdtree/point2"
	"testing"
)

func testTree() *KDTree {
	tree := New()
	tree.Add(point2.Point2{})
	tree.Add(point2.Point2{X: 1})
	tree.Add(point2.Point2{X: -1})
	tree.Add(point2.Point2{Y: 2})
	tree.Add(point2.Point2{Y: -2})
	tree.Add(point2.Point2{X: -1, Y: -1})
	tree.Add(point2.Point2{X: -1, Y: 1})
	return tree
}

func TestKDTree_Add(t *testing.T) {
	tree := testTree()
	if !equals(tree.root, point2.Point2{}) {
		t.Fail()
	}

	if !equals(tree.root.Left, point2.Point2{X:-1}) {
		t.Fail()
	}

	if !equals(tree.root.Right, point2.Point2{X:1}) {
		t.Fail()
	}

	if !equals(tree.root.Right.Right, point2.Point2{Y:2}) {
		t.Fail()
	}

	if !equals(tree.root.Right.Left, point2.Point2{Y:-2}) {
		t.Fail()
	}

	if !equals(tree.root.Left.Left, point2.Point2{X:-1, Y:-1}) {
		t.Fail()
	}

	if !equals(tree.root.Left.Right, point2.Point2{X:-1, Y:1}) {
		t.Fail()
	}


}

func TestKDTree_String(t *testing.T) {
	tree := testTree()
	_ = tree.String()
	tree.Remove(point2.Point2{X: -1, Y: 1})
	_ = tree.String()
}

func TestNode_Min(t *testing.T) {
	n1 := &node{point2.Point2{Y: 1}, nil, nil}
	n2 := &node{point2.Point2{X: 1}, nil, nil}
	if n1.Min(n2, 0) == n2 {
		t.Log("Incorrect value for dimension 0")
		t.Fail()
	}

	if n1.Min(n2, 1) == n1 {
		t.Log("Incorrect value for dimension 1")
		t.Fail()
	}
}

func TestNode_FindMin(t *testing.T) {
	tree := testTree()

	if tree.root.FindMin(0, 0).Dimension(0) != -1 {
		t.Log("Incorrect value for dimension 0")
		t.Fail()
	}

	if tree.root.FindMin(1, 0).Dimension(1) != -2 {
		t.Log("Incorrect value for dimension 1")
		t.Fail()
	}
}

func TestNode_Find(t *testing.T) {
	tree := testTree()

	ret, parent, dim := tree.root.Find(point2.Point2{}, 0, nil)
	if ret != tree.root || parent != nil || dim != 0 {
		t.Log("Incorrect value for (0,0)")
		t.Fail()
	}

	ret, parent, dim = tree.root.Find(point2.Point2{X: -1}, 0, nil)
	if ret != tree.root.Left || parent != tree.root || dim != 1 {
		t.Log("Incorrect value for (-1,0)")
		t.Fail()
	}

	ret, parent, dim = tree.root.Find(point2.Point2{X: 1}, 0, nil)
	if ret != tree.root.Right || parent != tree.root || dim != 1 {
		t.Log("Incorrect value for (1,0)")
		t.Fail()
	}

	ret, parent, dim = tree.root.Find(point2.Point2{Y: 2}, 0, nil)
	if ret != tree.root.Right.Right || parent != tree.root.Right || dim != 0 {
		t.Log("Incorrect value for (0,2)")
		t.Log("Expected", tree.root.Right.Right, tree.root.Right, 0)
		t.Log("Received", ret, parent, dim)
		t.Fail()
	}

	ret, parent, dim = tree.root.Find(point2.Point2{Y: -2}, 0, nil)
	if ret != tree.root.Right.Left || parent != tree.root.Right || dim != 0 {
		t.Log("Incorrect value for (0,-2)")
		t.Log("Expected", tree.root.Right.Left, tree.root.Right, 0)
		t.Log("Received", ret, parent, dim)
		t.Fail()
	}

	ret, parent, dim = tree.root.Right.Right.Find(point2.Point2{Y: 2}, 0, tree.root.Right)
	if ret != tree.root.Right.Right || parent != tree.root.Right || dim != 0 {
		t.Log("Incorrect value for (0,-2)")
		t.Log("Expected", tree.root.Right.Right, tree.root.Right, 0)
		t.Log("Received", ret, parent, dim)
		t.Fail()
	}

	ret, _, _ = tree.root.Find(point2.Point2{X: 10, Y: 10}, 0, nil)
	if ret != nil {
		t.Log("Incorrect value for (10,10)")
		t.Fail()
	}
}

func TestKDTree_Remove(t *testing.T) {
	tree := New()
	tree.Add(point2.Point2{})
	tree.Add(point2.Point2{X: 1})
	tree.Add(point2.Point2{X: -1})
	tree.Remove(point2.Point2{X: 1})
	if tree.root.Right != nil {
		t.Fail()
	}
	tree.Remove(point2.Point2{X: -1})
	if tree.root.Left != nil {
		t.Fail()
	}
	tree.Remove(point2.Point2{})
	if tree.root != nil {
		t.Fail()
	}

	tree = testTree()
	tree.Remove(point2.Point2{})
	if !equals(tree.root, point2.Point2{Y: 2}) {
		t.Fail()
	}

	tree = testTree()
	tree.Remove(point2.Point2{X: 1})
	if !equals(tree.root.Right, point2.Point2{Y: 2}) {
		t.Fail()
	}

	tree = testTree()
	tree.Remove(point2.Point2{X: 1})
	if !equals(tree.root.Right, point2.Point2{Y: 2}) {
		t.Fail()
	}

	tree.Remove(point2.Point2{})
	if !equals(tree.root, point2.Point2{Y: -2}) {
		t.Fail()
	}

	tree = testTree()
	tree.Remove(point2.Point2{X: -1})
	if !equals(tree.root.Left, point2.Point2{X:-1, Y:1}) {
		t.Fail()
	}

	if tree.root.Left.Right != nil {
		t.Fail()
	}

	tree.Remove(point2.Point2{X: -1, Y: -1})
	if tree.root.Left.Left != nil {
		t.Fail()
	}

	tree.Remove(point2.Point2{X: -1})
	if tree.root.Left.Left != nil {
		t.Fail()
	}

	tree = testTree()
	tree.Remove(point2.Point2{X:1})
	tree.Remove(point2.Point2{Y: 2})
	tree.Remove(point2.Point2{Y: -2})
	tree.Remove(point2.Point2{X:-1})
	if tree.root.Right != nil {
		t.Fail()
	}
	if !equals(tree.root.Left, point2.Point2{X:-1,Y:1}) {
		t.Fail()
	}
}
