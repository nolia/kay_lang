package parser

import (
	"reflect"
	"testing"

	"github.com/nolia/kay/lexer"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		tokens []lexer.Token
		want   *Node
	}{
		{
			name: "1 + 2",
			tokens: []lexer.Token{
				lexer.Token{Value: "1", Type: lexer.NUMBER},
				lexer.Token{Value: "+", Type: lexer.OPERATOR},
				lexer.Token{Value: "2", Type: lexer.NUMBER},
			},
			want: &Node{
				Type:  lexer.OPERATOR,
				Op:    "+",
				Left:  &Node{Type: lexer.NUMBER, Op: "1"},
				Right: &Node{Type: lexer.NUMBER, Op: "2"},
			},
		},

		{
			name: "a - 3*b",
			tokens: []lexer.Token{
				lexer.Token{Value: "a", Type: lexer.IDENTIFIER},
				lexer.Token{Value: "-", Type: lexer.OPERATOR},
				lexer.Token{Value: "3", Type: lexer.NUMBER},
				lexer.Token{Value: "*", Type: lexer.OPERATOR},
				lexer.Token{Value: "b", Type: lexer.IDENTIFIER},
			},
			want: &Node{
				Type: lexer.OPERATOR,
				Op:   "-",
				Left: &Node{Type: lexer.IDENTIFIER, Op: "a"},
				Right: &Node{
					Type:  lexer.OPERATOR,
					Op:    "*",
					Left:  &Node{Type: lexer.NUMBER, Op: "3"},
					Right: &Node{Type: lexer.IDENTIFIER, Op: "b"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
