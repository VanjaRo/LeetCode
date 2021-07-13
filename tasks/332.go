package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Print(findItinerary(tickets))
}

func findItinerary(tickets [][]string) []string {
	graph := newGraph(tickets)

	ret := []string{}

	graph.dfs(&ret, "JFK")

	reverse(&ret)

	return ret

}

type Queue []string

type Graph map[string]*Queue

func newGraph(tickets [][]string) *Graph {
	g := make(Graph)

	// edges init
	ensureArray := func(node string) {
		if _, ok := g[node]; !ok {
			q := &Queue{}
			g[node] = q
		}
	}

	// edges add
	for _, tiket := range tickets {
		from, to := tiket[0], tiket[1]
		ensureArray(from)
		ensureArray(to)
		*g[from] = append(*g[from], to)
	}

	// edges sort
	for _, node := range g {
		sort.Strings(*node)
	}

	return &g
}

func (g *Graph) dfs(itinerary *[]string, from string) {
	q := (*g)[from]

	for len(*q) > 0 {
		child := q.deque()
		g.dfs(itinerary, child)
	}
	*itinerary = append(*itinerary, from)
}

func (q *Queue) deque() string {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func reverse(ss *[]string) {
	n := len(*ss)
	for i := 0; i < n/2; i++ {
		(*ss)[i], (*ss)[n-1-i] = (*ss)[n-1-i], (*ss)[i]
	}
}
