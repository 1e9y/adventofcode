package day18

import (
	"fmt"
	"regexp"
	"strconv"

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

var plainNumberRe = regexp.MustCompile(`^\[(\d+),(\d+)\]$`)
var singleNestLestNumberRe = regexp.MustCompile(`^\[(\[\d+,\d+\]),(\d+)\]$`)
var singleNestRightNumberRe = regexp.MustCompile(`^\[(\d+),(\[\d+,\d+\])\]$`)
var numberRe = regexp.MustCompile(`^\[[0-9,\[\]]+\]$`)

func newNumberFromInput(input string) *Node {
	return newNode(input, nil, 0)
}

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
		panic(fmt.Errorf("bad input: not a number: %s", input))
	}

	if matches := plainNumberRe.FindStringSubmatch(input); matches != nil {
		if len(matches) != 3 {
			panic(fmt.Errorf("bad input: not a plain number:%s", input))
		}
		node = &Node{
			parent: parent,
			depth:  depth,
		}
		node.left = newNode(matches[1], node, depth+1)
		node.right = newNode(matches[2], node, depth+1)
		return node
	}
	if matches := singleNestLestNumberRe.FindStringSubmatch(input); matches != nil {
		if len(matches) != 3 {
			panic(fmt.Errorf("bad input: not a plain number:%s", input))
		}
		node = &Node{
			parent: parent,
			depth:  depth,
		}
		node.left = newNode(matches[1], node, depth+1)
		node.right = newNode(matches[2], node, depth+1)
		return node
	}
	if matches := singleNestRightNumberRe.FindStringSubmatch(input); matches != nil {
		if len(matches) != 3 {
			panic(fmt.Errorf("bad input: not a plain number:%s", input))
		}
		node = &Node{
			parent: parent,
			depth:  depth,
		}
		node.left = newNode(matches[1], node, depth+1)
		node.right = newNode(matches[2], node, depth+1)
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

func (node *Node) find(depth int) *Node {
	if !node.leaf && node.depth == depth {
		return node
	}

	if node.left != nil {
		if n := node.left.find(depth); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.find(depth); n != nil {
			return n
		}
	}

	return nil
}

func (node *Node) reduce() {
	canExpl := true
	canSplit := true
	// fmt.Println(node)
	for canExpl || canSplit {
		for canExpl {
			canExpl = (*node).explode()
		}
		// if canExpl {
		// 	fmt.Println("exploded:")
		// 	fmt.Println(node)
		// }

		canSplit = (*node).split()
		// if canSplit {
		// 	fmt.Println("splitted:")
		// 	fmt.Println(node)
		// }
		canExpl = (*node).explode()
		// canSplit = (*node).split()
	}
}

func (node *Node) explode() bool {
	n := node.find(4)
	// fmt.Printf("%s \n", n)

	if n == nil {
		return false
	}
	// find left
	// =========
	var root *Node
	cur := n
	par := cur.parent
	// fmt.Println(cur, par)
	for par != nil {
		// fmt.Println(">>")
		if par.right == cur {
			root = par
			break
		}

		cur = par
		par = par.parent
	}

	// fmt.Println("Parent L")
	// fmt.Println(root)
	if root != nil {
		root = root.left
		for !root.leaf {
			root = root.right
		}
		root.value += n.left.value
	}

	// find right
	// =========
	root = nil
	cur = n
	par = cur.parent
	// fmt.Println(cur, par)
	for par != nil {
		// fmt.Println(">>")
		if par.left == cur {
			root = par
			break
		}

		cur = par
		par = par.parent
	}
	// fmt.Println("Parent R")
	// fmt.Println(root)
	if root != nil {
		root = root.right
		for !root.leaf {
			root = root.left
		}
		root.value += n.right.value
	}
	// fmt.Println(root)

	// zero
	*n = Node{
		value:  0,
		leaf:   true,
		parent: n.parent,
		depth:  n.depth,
	}

	return true
}

func (node *Node) findGte(n int) *Node {
	if node.leaf && node.value >= n {
		return node
	}

	if node.left != nil {
		if n := node.left.findGte(n); n != nil {
			return n
		}
	}

	if node.right != nil {
		if n := node.right.findGte(n); n != nil {
			return n
		}
	}

	return nil
}

func (node Node) split() bool {
	n := node.findGte(10)
	if n == nil {
		return false
	}
	l := n.value / 2
	r := l + n.value%2
	*n = Node{
		leaf:   false,
		parent: n.parent,
		depth:  n.depth,
	}
	n.left = &Node{
		leaf:   true,
		value:  l,
		parent: n,
		depth:  n.depth + 1,
	}
	n.right = &Node{
		leaf:   true,
		value:  r,
		parent: n,
		depth:  n.depth + 1,
	}
	return true
}

func (node Node) String() string {
	if node.leaf {
		return fmt.Sprintf("%d", node.value)
	}
	return fmt.Sprintf("[%s,%s]", node.left, node.right)
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
	// fmt.Println(a)
	// fmt.Println("+")
	// fmt.Println(b)

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

	// fmt.Println("====")
	// fmt.Println(node)
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
