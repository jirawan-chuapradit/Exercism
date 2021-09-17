package change

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	for _, tc := range testCases {
		actual := Solution(tc.N, tc.A, tc.B)
		assert.Equal(t, tc.expectedOutput, actual)

	}
}

func BenchmarkChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solution(tc.N, tc.A, tc.B)
		}
	}
}
