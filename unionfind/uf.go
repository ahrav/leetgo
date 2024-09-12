package unionfind

import "sort"

type UnionFind struct {
	parent map[string]string
	rank   map[string]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
}

func (uf *UnionFind) Find(email string) string {
	if uf.parent[email] != email {
		uf.parent[email] = uf.Find(uf.parent[email])
	}

	return uf.parent[email]
}

func (uf *UnionFind) Union(email1, email2 string) {
	rootX, rootY := uf.Find(email1), uf.Find(email2)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootY] > uf.rank[rootX] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// AccountsMerge - https://leetcode.com/problems/accounts-merge/
func AccountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind()
	emailToName := make(map[string]string)

	for _, account := range accounts {
		name := account[0]
		firstEmail := account[1]
		for _, email := range account[1:] {
			if _, ok := uf.parent[email]; !ok {
				uf.parent[email] = email
			}
			uf.Union(firstEmail, email)
			emailToName[email] = name
		}
	}

	rootToEmails := make(map[string][]string)
	for email := range uf.parent {
		root := uf.Find(email)
		rootToEmails[root] = append(rootToEmails[root], email)
	}

	result := make([][]string, 0, len(rootToEmails))
	for root, emails := range rootToEmails {
		sort.Strings(emails)
		name := emailToName[root]
		result = append(result, append([]string{name}, emails...))
	}

	return result
}
