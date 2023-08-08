// Exercise 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
// that reveals the sequence of values in the tree.

package treesort

import (
	"bytes"
	"fmt"
	"strconv"
)

func (t *tree) String() string {
	var buf bytes.Buffer

	if t == nil {
		return ""
	}
	queue := make([]*tree, 1)
	queue[0] = t
	for len(queue) > 0 {
		node := queue[0]
		buf.WriteString(strconv.Itoa(node.value))
		buf.WriteByte(' ')
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}

	return buf.String()
}

func (t *tree) StringAlt() string {
	order := make([]int, 0)
	order = appendValues(order, t)
	if len(order) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v", order)
	return b.String()
}

// !+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
