package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmp(t *testing.T) {
	type cmpTestCase struct {
		a        Chain
		b        Chain
		ordered  int
		expected Chain
	}

	testCases := []cmpTestCase{
		{
			a:        Chain{250},
			b:        Chain{500},
			ordered:  1,
			expected: Chain{250},
		},
		{
			a:        Chain{250},
			b:        Chain{500},
			ordered:  250,
			expected: Chain{250},
		},
		{
			a:        Chain{500},
			b:        Chain{250, 250},
			ordered:  251,
			expected: Chain{500},
		},
		{
			a:        Chain{500, 250},
			b:        Chain{1000, 250, 250, 250},
			ordered:  501,
			expected: Chain{500, 250},
		},
		{
			a:        Chain{5000, 5000, 2000, 250},
			b:        Chain{5000, 5000, 5000},
			ordered:  12001,
			expected: Chain{5000, 5000, 2000, 250},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := compareAndReturnBest(tc.ordered, tc.a, tc.b)
			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		ordered  int
		packages []int
		expected []int
	}{
		{name: "Order 1 item", ordered: 1, packages: []int{250, 500, 1000, 2000, 5000}, expected: []int{250}},
		{name: "Order 250 items", ordered: 250, packages: []int{250, 500, 1000, 2000, 5000}, expected: []int{250}},
		{name: "Order 251 items", ordered: 251, packages: []int{250, 500, 1000, 2000, 5000}, expected: []int{500}},
		{name: "Order 501 items", ordered: 501, packages: []int{250, 500, 1000, 2000, 5000}, expected: []int{500, 250}},
		{name: "Order 12001 items", ordered: 12001, packages: []int{250, 500, 1000, 2000, 5000}, expected: []int{5000, 5000, 2000, 250}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculate(test.ordered, test.packages)
			assert.ElementsMatch(t, test.expected, result)
			fmt.Println(test.expected, result)
		})
	}
}
