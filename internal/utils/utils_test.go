package utils

import (
	"fmt"
	route "ova_route_api/internal/models"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestSplitSlice(t *testing.T) {
	resp, err := SplitSlice([]string{}, 5)
	assert.Assert(t, err == nil)
	assert.Assert(t, fmt.Sprint(resp) == "[]")

	resp, err = SplitSlice([]string{"A", "B", "C", "D", "E"}, 2)
	assert.Assert(t, err == nil)
	assert.Assert(t, fmt.Sprint(resp) == "[[A B] [C D] [E]]")

	resp, err = SplitSlice([]string{"A", "B", "C", "D", "E"}, 0)
	assert.Assert(t, err.Error() == "batchSize should be more than zero")
	assert.Assert(t, resp == nil)
}

func TestReverseKey(t *testing.T) {
	resp, err := ReverseKey(map[int]string{})
	assert.Assert(t, err == nil)
	assert.Assert(t, reflect.DeepEqual(resp, map[string]int{}))

	resp, err = ReverseKey(map[int]string{0: "A", 1: "B", 2: "C", 3: "D"})
	assert.Assert(t, err == nil)
	assert.Assert(t, reflect.DeepEqual(resp, map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}))

	resp, err = ReverseKey(map[int]string{0: "A", 1: "A", 2: "A", 3: "A"})
	assert.Assert(t, err.Error() == "key is duplicated")
	assert.Assert(t, resp == nil)
}

func TestFilterSlice(t *testing.T) {
	arr := []int{}
	resp, err := FilterSlice(arr, Ints)
	assert.Assert(t, fmt.Sprint(resp) == "[]")
	assert.Assert(t, err == nil)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resp, err = FilterSlice(arr, []int{})
	assert.Assert(t, fmt.Sprint(resp) == "[]")
	assert.Assert(t, err == nil)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resp, err = FilterSlice(arr, Ints)
	assert.Assert(t, fmt.Sprint(resp) == "[1 3 5 7]")
	assert.Assert(t, err == nil)
}

func TestSplitToBulks(t *testing.T) {
	// empty routes
	routes := []route.Route{}
	resp, err := SplitToBulks(routes, 2)
	assert.Assert(t, fmt.Sprint(resp) == "[]")
	assert.Assert(t, err == nil)

	// sample routes
	routes = []route.Route{
		{
			ID:        1,
			UserID:    1,
			RouteName: "Tets route",
			Length:    10.5,
		},
		{
			ID:        2,
			UserID:    1,
			RouteName: "Tets route 2",
			Length:    6,
		},
		{
			ID:        3,
			UserID:    2,
			RouteName: "Tets route 3",
			Length:    13,
		},
	}
	resp, err = SplitToBulks(routes, 2)
	assert.Assert(t, fmt.Sprint(resp) == "[[{1 1 Tets route 10.5} {2 1 Tets route 2 6}] [{3 2 Tets route 3 13}]]")
	assert.Assert(t, err == nil)

	// change base routes
	resp, err = SplitToBulks(routes, 2)

	//nolint
	routes = []route.Route{
		{
			ID:        4,
			UserID:    3,
			RouteName: "Tets route 4",
			Length:    11,
		},
	}
	assert.Assert(t, fmt.Sprint(resp) != "[[{4 3 Tets route 4 11}]]")
	assert.Assert(t, err == nil)
}

func TestConvertToView(t *testing.T) {
	// empty routes
	routes := []route.Route{}
	resp, err := ConvertToView(routes)
	assert.Assert(t, resp == nil)
	assert.Assert(t, err.Error() == "empty routes list")

	// sample routes
	routes = []route.Route{
		{
			ID:        1,
			UserID:    1,
			RouteName: "Tets route",
			Length:    10.5,
		},
		{
			ID:        2,
			UserID:    1,
			RouteName: "Tets route 2",
			Length:    6,
		},
		{
			ID:        3,
			UserID:    2,
			RouteName: "Tets route 3",
			Length:    13,
		},
	}

	expected := map[uint64]route.Route{
		0: {
			ID:        1,
			UserID:    1,
			RouteName: "Tets route",
			Length:    10.5,
		},
		1: {
			ID:        2,
			UserID:    1,
			RouteName: "Tets route 2",
			Length:    6,
		},
		2: {
			ID:        3,
			UserID:    2,
			RouteName: "Tets route 3",
			Length:    13,
		},
	}
	resp, err = ConvertToView(routes)
	assert.DeepEqual(t, resp, expected)
	assert.Assert(t, err == nil)
}
