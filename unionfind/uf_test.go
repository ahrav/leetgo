package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]string
		expected [][]string
	}{
		{
			name: "Merges accounts with common emails",
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			name: "Merges accounts with no common emails",
			accounts: [][]string{
				{"John", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			expected: [][]string{
				{"John", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			name: "Merges single account",
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "johnsmith@mail.com"},
			},
		},
		{
			name:     "Merges empty accounts",
			accounts: [][]string{},
			expected: [][]string{},
		},
		{
			name: "Merges accounts with multiple names",
			accounts: [][]string{
				{"John", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com", "mary_newyork@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com", "mary_newyork@mail.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := AccountsMerge(tt.accounts)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkAccountsMerge(b *testing.B) {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	for i := 0; i < b.N; i++ {
		AccountsMerge(accounts)
	}
}
