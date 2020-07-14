package rest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractVersionAndCompare(t *testing.T) {
	type test struct {
		input           string
		requiredMinimum string
		expectedIsMet   bool
		expectedErr     error
	}

	tests := []test{
		test{
			input:           "notavalidversiontag",
			requiredMinimum: "1.2.3",
			expectedIsMet:   false,
			expectedErr:     fmt.Errorf("unexpected input version tag: notavalidversiontag"),
		},
		test{
			input:           "abc-v0.1.2",
			requiredMinimum: "invalidrequired",
			expectedIsMet:   false,
			expectedErr:     fmt.Errorf("unexpected threshold version tag: invalidrequired"),
		},

		// valid matches

		// exact match
		test{
			input:           "abc-v0.1.2",
			requiredMinimum: "0.1.2",
			expectedIsMet:   true,
			expectedErr:     nil,
		},

		// every digit bigger
		test{
			input:           "abc-v1.2.3",
			requiredMinimum: "0.1.2",
			expectedIsMet:   true,
			expectedErr:     nil,
		},

		// only major bigger
		test{
			input:           "abc-v1.0.0",
			requiredMinimum: "0.1.2",
			expectedIsMet:   true,
			expectedErr:     nil,
		},

		// only minor bigger
		test{
			input:           "abc-v0.2.0",
			requiredMinimum: "0.1.2",
			expectedIsMet:   true,
			expectedErr:     nil,
		},

		// only patch bigger
		test{
			input:           "abc-v0.1.2",
			requiredMinimum: "0.1.2",
			expectedIsMet:   true,
			expectedErr:     nil,
		},

		// invalid requirements

		// only patch smaller
		test{
			input:           "abc-v0.1.1",
			requiredMinimum: "0.1.2",
			expectedIsMet:   false,
			expectedErr:     nil,
		},

		// only minor smaller
		test{
			input:           "abc-v0.0.9",
			requiredMinimum: "0.1.2",
			expectedIsMet:   false,
			expectedErr:     nil,
		},

		// only major smaller
		test{
			input:           "abc-v0.9.9",
			requiredMinimum: "1.1.2",
			expectedIsMet:   false,
			expectedErr:     nil,
		},
	}

	for _, test := range tests {
		ok, err := extractVersionAndCompare(test.input, test.requiredMinimum)
		assert.Equal(t, test.expectedIsMet, ok)
		assert.Equal(t, test.expectedErr, err)
	}
}