package parser

import "github.com/nolia/kay/lexer"

type Node struct {
	Type        lexer.TokenType
	Op          string
	Left, Right *Node
}

func toNode(t lexer.Token) *Node {
	return &Node{
		Op:   t.Value,
		Type: t.Type,
	}
}

func pullOut(nodes []*Node, ops []string) []*Node {
	i := 0
	for i < len(nodes) {
		n := nodes[i]
		token := n.Op

		if n.Type != lexer.OPERATOR {
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

		if n.Left != nil && n.Right != nil {
			i++
			continue
		}

		// Validate index.
		if i-1 < 0 || i+1 >= len(nodes) {
			panic("Invalid indexes!!!")
		}

		nodes[i].Left = nodes[i-1]
		nodes[i].Right = nodes[i+1]

		// Remove i-1 and i+1 elements.
		left := append(nodes[:i-1], nodes[i])
		nodes = append(left, nodes[i+2:]...)

	}

	return nodes
}

func Parse(tokens []lexer.Token) *Node {
	nodes := []*Node{}

	// Map to nodes:
	for _, t := range tokens {
		nodes = append(nodes, toNode(t))
	}

	// Parse parentheses first with stack.
	stack := []int{}
	i := 0

	for i < len(nodes) {
		n := nodes[i]
		if n.Type == lexer.OPEN_PAR {
			stack = append(stack, i)
			i++
			continue
		}

		if n.Type == lexer.OPEN_PAR {
			if len(stack) == 0 {
				panic("Invalid closing parentheses!")
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			end := i
			// Parse the insides:
			inside := Parse(tokens[start+1 : end])

			lnodes := append(nodes[:start], inside)
			nodes = append(lnodes, nodes[end+1:]...)

			ltokens := append(tokens[:start], lexer.Token{
				Value: inside.Op,
				Type:  inside.Type,
			})
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
