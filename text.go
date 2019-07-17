package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

type EvalType string

const (
	NUMBER   EvalType = "number"
	OPERATOR EvalType = "operator"
	OTHER    EvalType = "other"
)

type EvalNode struct {
	op          string
	left, right *EvalNode
}

func (e *EvalNode) getType() EvalType {
	switch {
	case isDigits(e.op):
		return NUMBER
	case isOperator(e.op):
		return OPERATOR
	default:
		return OTHER
	}
}

func (e *EvalNode) Value() int {
	if isDigits(e.op) {
		a, _ := strconv.Atoi(e.op)
		return a
	}
	switch e.op {
	case "+":
		return e.left.Value() + e.right.Value()

	case "-":
		return e.left.Value() - e.right.Value()

	case "*":
		return e.left.Value() * e.right.Value()
	case "/":
		return e.left.Value() / e.right.Value()
	}

	return 0
}

func (e *EvalNode) String() string {
	if e.left == nil && e.right == nil {
		return e.op
	}

	return fmt.Sprintf("%v(%v, %v)", e.op, e.left, e.right)
}

func isOperator(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true

	default:
		return false
	}
}

func isDigits(s string) bool {
	match, _ := regexp.MatchString("[0-9]+", s)
	return match
}

func pullOut(nodes []*EvalNode, ops []string) []*EvalNode {
	i := 0
	for i < len(nodes) {
		n := nodes[i]
		token := n.op

		if n.getType() != OPERATOR {
			i++
			continue
		}

		found := false
		for _, op := range ops {
			if token == op {
				found = true
				break
			}
		}
		if !found {
			i++
			continue
		}

		if n.left != nil && n.right != nil {
			i++
			continue
		}

		// Validate index.
		if i-1 < 0 || i+1 >= len(nodes) {
			panic("Invalid indexes!!!")
		}

		nodes[i].left = nodes[i-1]
		nodes[i].right = nodes[i+1]

		// Remove i-1 and i+1 elements.
		left := append(nodes[:i-1], nodes[i])
		nodes = append(left, nodes[i+2:]...)

	}

	return nodes
}

func parse(tokens []string) *EvalNode {
	if len(tokens) == 0 {
		return nil
	}

	// Map to EvalNodes:
	nodes := []*EvalNode{}
	for _, s := range tokens {
		nodes = append(nodes, &EvalNode{op: s})
	}

	// Parse parentheses first with stack.
	stack := []int{}
	i := 0
	for i < len(nodes) {
		n := nodes[i]
		if n.op == "(" {
			stack = append(stack, i)
			i++
			continue
		}

		if n.op == ")" {
			if len(stack) == 0 {
				panic("Invalid closing parentheses!")
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			end := i
			// Parse the insides:
			inside := parse(tokens[start+1 : end])

			lnodes := append(nodes[:start], inside)
			nodes = append(lnodes, nodes[end+1:]...)

			ltokens := append(tokens[:start], inside.op)
			tokens = append(ltokens, tokens[end+1:]...)

			i = start + 1
			continue
		}

		i++
	}

	// Parse oprations by priorites.
	priorites := [][]string{
		[]string{"*", "/"},
		[]string{"+", "-"},
	}
	for _, ops := range priorites {
		nodes = pullOut(nodes, ops)
	}

	if len(nodes) != 1 {
		panic("Not pull out all operations!")
	}

	return nodes[0]

}

func eval(code string) string {
	var s scanner.Scanner
	s.Init(strings.NewReader(code))
	s.Filename = "default"

	tokens := []string{}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}

	node := parse(tokens)

	fmt.Printf("code: %q\nnodes: %v\n\n", code, node)

	if node != nil {
		return fmt.Sprintf("[%v]\n", node.Value())
	}

	return ""
}

func main() {
	code := "2 * (3 * 4) * (3+1)"

	res := eval(code)
	fmt.Printf(res)

	// tokens := []string{"1", "+", "2", "*", "3", "+", "5"}
	// node := parse(tokens)

	// fmt.Printf("%q\n%v\n", tokens, node)

}
