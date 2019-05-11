package treesort

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(value, root)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(value int, t *tree) *tree {
	if t == nil {
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add(value, t.left)
	} else {
		t.right = add(value, t.right)
	}
	return t
}
