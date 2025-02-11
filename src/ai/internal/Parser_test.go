package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		errNil bool
	}{
		{
			name:   "Base Test L0",
			input:  "L0",
			errNil: true,
		},
		{
			name:   "Base Test L1",
			input:  "L1",
			errNil: true,
		},
		{
			name:   "Base Test L0",
			input:  "L0",
			errNil: true,
		},
		{
			name:   "Base Test L1",
			input:  "L1",
			errNil: true,
		},
		{
			name:   "Base Test L2",
			input:  "L2",
			errNil: true,
		},
		{
			name:   "Base Test L3",
			input:  "L3",
			errNil: true,
		},
		{
			name:   "Base Test L4",
			input:  "L4",
			errNil: true,
		},
		{
			name:   "Base Test EvaluableExpression",
			input:  "L3(1+(2*(3+4)-5))",
			errNil: true,
		},
		{
			name:   "Test L3 with argument",
			input:  "L3(3, L1)",
			errNil: true,
		},
		{
			name:   "Test L4 with L3 argument",
			input:  "L4(L3(3, L1))",
			errNil: true,
		},
		{
			name:   "Error Test (Invalid Model)",
			input:  "L7",
			errNil: false,
		},
		{
			name:   "Error Test (Missing Parentheses)",
			input:  "L3(3",
			errNil: false,
		},
		{
			name:   "Error Test (Mismatched Parentheses)",
			input:  "L3(3))",
			errNil: false,
		},
		{
			name:   "Complex Recursive Test (L3 inside L4)",
			input:  "L4(L3(10, L1))",
			errNil: true,
		},
		{
			name:   "Complex Recursive Test (L4 inside L3 inside L4)",
			input:  "L4(L3(1,L4(L3(1,L4(L3(1,L4(L3())))))))",
			errNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewInterface(tt.input)
			if tt.errNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
