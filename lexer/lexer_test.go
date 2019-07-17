package lexer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	// Setup
	f_debug = true

	tests := []struct {
		name string
		s    string
		want []Token
	}{
		// TODO: Add test cases.
		{
			name: "Number",
			s:    " 12345 ",
			want: []Token{Token{Type: NUMBER, Value: "12345"}},
		},
		{
			name: "whitespace",
			s: "    		",
			want: []Token{},
		},
		{
			name: "Identifier",
			s:    "a123",
			want: []Token{Token{Value: "a123", Type: IDENTIFIER}},
		},
		{
			name: "Identifier + Number",
			s:    "a123+987",
			want: []Token{
				Token{Type: IDENTIFIER, Value: "a123"},
				Token{Type: OPERATOR, Value: "+"},
				Token{Type: NUMBER, Value: "987"},
			},
		},

		{
			name: "Complex operation",
			s:    "a + 2*3 + b/4",
			want: []Token{
				Token{Type: IDENTIFIER, Value: "a"},
				Token{Type: OPERATOR, Value: "+"},
				Token{Type: NUMBER, Value: "2"},
				Token{Type: OPERATOR, Value: "*"},
				Token{Type: NUMBER, Value: "3"},
				Token{Type: OPERATOR, Value: "+"},
				Token{Type: IDENTIFIER, Value: "b"},
				Token{Type: OPERATOR, Value: "/"},
				Token{Type: NUMBER, Value: "4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
