package utils

import (
	"fmt"
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
