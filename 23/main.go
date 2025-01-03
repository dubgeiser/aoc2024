package main

import (
	"aoc/lib/collections/set"
	"aoc/lib/input"
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func main() {
	// network maps computer to all the computers it is connected to.
	// We store each link, bidirectional
	network := map[string]*set.Set[string]{}
	input.Lines(func(s *bufio.Scanner) {
		c := strings.Split(s.Text(), "-")
		if _, ok := network[c[0]]; !ok {
			network[c[0]] = set.New[string]()
		}
		network[c[0]].Add(c[1])
		if _, ok := network[c[1]]; !ok {
			network[c[1]] = set.New[string]()
		}
		network[c[1]].Add(c[0])
	})

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
	fmt.Println(sets.Len())

	// --[ Part 2 ]--
	// Collect all the LANs and print out the longest.
	lans := set.New[string]()
	longest := ""
	for c := range network {
		mklans(c, []string{c}, lans, network)
	}
	for l := range lans.Values() {
		if len(l) > len(longest) {
			longest = l
		}
	}
	fmt.Println(longest)
}

func mklans(computer string, lanCurr []string, lans *set.Set[string], network map[string]*set.Set[string]) {
	// We need a comparable to store in sets.
	// Since we need to represent the longest LAN as a comma separated, ordered,
	// list of computers, we already pre-format.
	slices.Sort(lanCurr)
	currs := strings.Join(lanCurr, ",")

	// If we processed the current LAN already, this means that the current LAN
	// is the largest one we could build for the given computer, ie.
	// `allConnected` was false in the previous call and nothing has been added
	// to `lanCurr`.
	if lans.Contains(currs) {
		return
	}
	lans.Add(currs)

	// Recursive DFS the network to find the LAN (ie. all the computers that
	// have connections to each other).
	for cc := range network[computer].Values() {
		if slices.Contains(lanCurr, cc) {
			continue
		}
		allConnected := true
		for _, sc := range lanCurr {
			if !network[sc].Contains(cc) {
				allConnected = false
				break
			}
		}
		if allConnected {
			lanCurr = append(lanCurr, cc)
		}
		mklans(cc, lanCurr, lans, network)
	}
}
