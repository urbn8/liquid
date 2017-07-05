package expression

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var parseErrorTests = []struct{ in, expected string }{
	{"a syntax error", "parse error"},
}

func TestExpressionParseErrors(t *testing.T) {
	for i, test := range parseErrorTests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			expr, err := Parse(test.in)
			require.Nilf(t, expr, test.in)
			require.Errorf(t, err, test.in, test.in)
			require.Containsf(t, err.Error(), test.expected, test.in)
		})
	}
}