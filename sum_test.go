package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run usecase unit test
// func TestSum(t *testing.T) {
// 	t.Run("2 and 3, should return 5", func(t *testing.T) {
// 		result := Sum(2, 3)
// 		assert.Equal(t, 5, result)
// 	})

// 	t.Run("10 and 5, should return 15", func(t *testing.T) {
// 		result := Sum(10, 5)
// 		assert.Equal(t, 13, result)
// 	})
// }

// table driven unit test
func TestSum2(t *testing.T) {
	testCase := []struct {
		name string
		a int
		b int
		expected int
	}{
		{
			name: "2 and 3, should return 5",
			a: 3,
			b: 2,
			expected: 5,
		},
		{
			name: "10 and 5, should return 15",
			a: 10,
			b: 5,
			expected: 15,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			assert.Equal(t, tc.expected, result)
		})
	}
}