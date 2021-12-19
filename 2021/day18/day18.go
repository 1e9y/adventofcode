package day18

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Node struct {
	value       int
	leaf        bool
	left, right *Node
	parent      *Node
	depth       int
}

var numberRe = regexp.MustCompile(`^\[[0-9,\[\]]+\]$`)

func newNode(input string, parent *Node, depth int) (node *Node) {
	if n, err := strconv.Atoi(input); err == nil {
		return &Node{
			value:  n,
			leaf:   true,
			parent: parent,
			depth:  depth,
		}
	}

	if !numberRe.MatchString(input) {
		panic(fmt.Errorf("bad input: not a snailfish number: %s", input))
	}

	if input[1] != '[' {
		i := strings.Index(input, ",")
		node = &Node{
			parent: parent,
			depth:  depth,
		}

		node.left = newNode(input[1:i], node, depth+1)
		node.right = newNode(input[i+1:len(input)-1], node, depth+1)
		return node
	}

	b := 0
	for i := 1; i < len(input)-1; i++ {
		switch input[i] {
		case '[':
			b++
		case ']':
			b--
		}
		if b == 0 {
			node = &Node{
				parent: parent,
				depth:  depth,
			}

			node.left = newNode(input[1:i+1], node, depth+1)
			node.right = newNode(input[i+2:len(input)-1], node, depth+1)
			return node
		}
	}

	panic(fmt.Errorf("bad input: %s", input))
}

func newNumberFromInput(input string) *Node {
	return newNode(input, nil, 0)
}

func (node Node) String() string {
	if node.leaf {
		return fmt.Sprintf("%d", node.value)
	}
	return fmt.Sprintf("[%s,%s]", node.left, node.right)
}

func (node *Node) findDepth(depth int) *Node {
	if !node.leaf && node.depth == depth {
		return node
	}

	if node.left != nil {
		if n := node.left.findDepth(depth); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.findDepth(depth); n != nil {
			return n
		}
	}

	return nil
}

func (node *Node) findValueGreaterThan(n int) *Node {
	if node.leaf && node.value >= n {
		return node
	}

	if node.left != nil {
		if n := node.left.findValueGreaterThan(n); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.findValueGreaterThan(n); n != nil {
			return n
		}
	}

	return nil
}

func (node *Node) findLeft() *Node {
	var base *Node
	current := node
	parent := current.parent
	for parent != nil {
		if parent.right == current {
			base = parent
			break
		}
		current = parent
		parent = parent.parent
	}

	if base != nil {
		base = base.left
		for !base.leaf {
			base = base.right
		}
	}

	return base
}

func (node *Node) findRight() *Node {
	var base *Node
	current := node
	parent := current.parent
	for parent != nil {
		if parent.left == current {
			base = parent
			break
		}
		current = parent
		parent = parent.parent
	}

	if base != nil {
		base = base.right
		for !base.leaf {
			base = base.left
		}
	}

	return base
}

func (node *Node) explode() bool {
	n := node.findDepth(4)

	if n == nil {
		return false
	}

	leftNode := n.findLeft()
	if leftNode != nil {
		leftNode.value += n.left.value
	}

	rightNode := n.findRight()
	if rightNode != nil {
		rightNode.value += n.right.value
	}

	*n = Node{
		value:  0,
		leaf:   true,
		parent: n.parent,
		depth:  n.depth,
	}

	return true
}

func (node Node) split() bool {
	n := node.findValueGreaterThan(10)

	if n == nil {
		return false
	}

	leftValue := n.value / 2
	rightValue := leftValue + n.value%2

	*n = Node{
		leaf:   false,
		parent: n.parent,
		depth:  n.depth,
	}

	n.left = &Node{
		value:  leftValue,
		leaf:   true,
		parent: n,
		depth:  n.depth + 1,
	}

	n.right = &Node{
		value:  rightValue,
		leaf:   true,
		parent: n,
		depth:  n.depth + 1,
	}

	return true
}

func (node *Node) reduce() {
	exploded := true
	splitted := true
	for exploded || splitted {
		for exploded {
			exploded = node.explode()
		}
		splitted = node.split()
		exploded = node.explode()
	}
}

func (node *Node) deepen() {
	node.depth++
	if node.left != nil {
		node.left.deepen()
	}
	if node.right != nil {
		node.right.deepen()
	}
}

func add(a, b *Node) (node *Node) {
	node = &Node{
		left:   a,
		right:  b,
		leaf:   false,
		parent: nil,
		depth:  0,
	}
	node.left.parent = node
	node.right.parent = node
	node.left.deepen()
	node.right.deepen()
	node.reduce()
	return
}

func (node *Node) magnitude() int {
	if node.leaf {
		return node.value
	}
	return node.left.magnitude()*3 + node.right.magnitude()*2
}

func homework(input <-chan string) int {
	number := newNumberFromInput(<-input)
	for line := range input {
		number = add(number, newNumberFromInput(line))
	}
	return number.magnitude()
}

func secondHomework(input []string) (max int) {
	var number *Node
	for i := range input {
		for j := range input {
			if i == j {
				continue
			}
			number = add(newNumberFromInput(input[i]), newNumberFromInput(input[j]))
			max = util.MaxInt(max, number.magnitude())
		}
	}
	return
}

func A(input *challenge.Challenge) int {
	return homework(input.Lines())
}

func B(input *challenge.Challenge) int {
	return secondHomework(input.LineSlice())
}
