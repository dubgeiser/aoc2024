package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/file"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println()

	// network maps computer to all the computers it is connected to.
	// We store each link, bidirectional
	network := map[string]*set.Set[string]{}
	for _, line := range strings.Split(file.Read(), "\n") {
		c := strings.Split(line, "-")
		if _, ok := network[c[0]]; !ok {
			network[c[0]] = set.New[string]()
		}
		network[c[0]].Add(c[1])
		if _, ok := network[c[1]]; !ok {
			network[c[1]] = set.New[string]()
		}
		network[c[1]].Add(c[0])
	}

	sets := set.New[string]()

	// for all computers c1
	for c1 := range network {
		// for all computers c2 connected to c1
		for c2 := range network[c1].Values() {
			// for all computers c3 connected to c2
			for c3 := range network[c2].Values() {
				// if c1 == c3, it's the original bi-directional link and not a
				// set of 3 different computers.
				// So c3 cannot be the first c1, but it does have to be connected
				// to it.
				if c1 != c3 && network[c3].Contains(c1) {
					if strings.HasPrefix(c1, "t") || strings.HasPrefix(c2, "t") || strings.HasPrefix(c3, "t") {
						// []string is not comparable, so we cannot make a set.Set
						// of it, but we need to sort our set so that we can
						// uniquely identify it in the set.  But slices.Sort() does
						// not sort arrays.
						set := []string{c1, c2, c3}
						slices.Sort(set)
						sets.Add(strings.Join(set, ""))
					}
				}
			}
		}
	}

	// 2332 too high
	fmt.Println(sets.Len())
}
