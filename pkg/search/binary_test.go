package search_test

import (
	"fmt"
	"testing"

	"github.com/leorcvargas/grokking-algorithms/pkg/search"
)

func TestBinarySearch(t *testing.T) {
	list := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		list = append(list, i*10)
	}

	testCases := []struct {
		desc   string
		target int
	}{
		{target: 50},
		{target: 150},
		{target: 250},
		{target: 350},
		{target: 450},
		{target: 550},
		{target: 650},
		{target: 750},
		{target: 850},
		{target: 950},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("searching for %d", tC.target), func(t *testing.T) {
			position := search.BinarySearch(list, tC.target)

			if tC.target != list[position] {
				t.Errorf("expected %d, got %d", tC.target, list[position])
			}
		})
	}

	failTestCases := []struct {
		desc   string
		target int
	}{
		{target: -50},
		{target: 321},
		{target: 59},
		{target: 1},
		{target: 45},
		{target: 51},
		{target: 65},
		{target: 87},
		{target: 8421},
		{target: 95041234},
	}
	for _, tC := range failTestCases {
		t.Run(fmt.Sprintf("searching for %d", tC.target), func(t *testing.T) {
			position := search.BinarySearch(list, tC.target)

			if position != search.BinarySearchNotFound {
				t.Errorf("expected %d to not be found, got position %d", tC.target, position)
			}
		})
	}
}
