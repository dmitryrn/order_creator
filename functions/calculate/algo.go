package main

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Chain []int

func (c *Chain) WastedItems(orderedItems int) int {
	items := c.SumItems()
	if items < orderedItems {
		return 0
	}
	return items - orderedItems
}

// could optimize by using a running sum
func (c *Chain) SumItems() int {
	items := 0
	for _, p := range *c {
		items += p
	}
	return items
}

func (a *Chain) Cmp(ordered int, b Chain) int {
	wastedItemsA := a.WastedItems(ordered)
	wastedItemsB := b.WastedItems(ordered)

	totalPackagesA := len(*a)
	totalPackagesB := len(b)

	if wastedItemsA != wastedItemsB {
		return wastedItemsA - wastedItemsB
	}

	return totalPackagesA - totalPackagesB
}

// this is not efficient for comparison, could compare in a more efficient way
func (c Chain) String() string {
	sortedChain := make([]int, len(c))
	copy(sortedChain, c)

	sort.Ints(sortedChain)

	strs := make([]string, len(sortedChain))
	for i, num := range sortedChain {
		strs[i] = strconv.Itoa(num)
	}
	return "[" + strings.Join(strs, ",") + "]"
}

func compareAndReturnBest(ordered int, a Chain, b Chain) Chain {
	if a.SumItems() < ordered {
		return b
	} else if b.SumItems() < ordered {
		return a
	}

	if a.String() == b.String() {
		return a
	}

	if a.Cmp(ordered, b) < 0 {
		return a
	} else {
		return b
	}
}

func calculate(ordered int, availablePackages []int) []int {
	slices.SortFunc(availablePackages, func(a int, b int) int { return b - a })
	// fmt.Println(availablePackages)

	result := &Chain{}
	memo := map[int]int{}
	dfs(Chain{}, availablePackages, result, ordered, memo)

	return *result
}

// recursive function, could run into stack overflow, but possible to rewrite into a non-recursive function
func dfs(chain Chain, availablePackages []int, result *Chain, ordered int, memo map[int]int) {
	// fmt.Println(chain)
	sum := chain.SumItems()
	if count, ok := memo[sum]; ok && count <= len(chain) {
		return
	}
	memo[sum] = len(chain)
	if sum >= ordered {
		// fmt.Println("comparing", chain.String(), "with", result)
		*result = compareAndReturnBest(ordered, chain, *result)
		// fmt.Println("winner", (*result).String())
		return
	}
	for _, p := range availablePackages {
		newChain := make(Chain, len(chain)) // 1 1
		copy(newChain, chain)               // 1 1
		newChain = append(newChain, p)      // 1 1 1

		dfs(newChain, availablePackages, result, ordered, memo)
	}
}
