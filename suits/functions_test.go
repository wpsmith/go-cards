package suits

import (
    "testing"
    "github.com/wpsmith/go-cards/cards-testing"
)

/** Test SliceIndex **/
type testSliceIndexStr struct {
    val   string
    slice []string
    index int
}

var (
    strSlice1 []string = []string{"C", "B", "K", "A"}
    strSlice2 []string = []string{"Z", "O", "W", "X"}
    testsSliceIndexStr = []testSliceIndexStr{
        {"C", strSlice1, 0},
        {"B", strSlice1, 1},
        {"K", strSlice1, 2},
        {"A", strSlice1, 3},
        {"Z", strSlice1, 4},
        {"", strSlice1, 4},

        {"Z", strSlice2, 0},
        {"O", strSlice2, 1},
        {"W", strSlice2, 2},
        {"X", strSlice2, 3},
        {"A", strSlice2, 4},
        {"", strSlice2, 4},
    }
)

func TestSliceIndexStr(t *testing.T) {
    for _, pair := range testsSliceIndexStr {
        v := SliceIndex(len(pair.slice), func(i int) bool {
            return pair.slice[i] == pair.val
        })
        if v != pair.index {
            pair.val = cards_testing.GetEmptyStringMaybe(pair.val)
            t.Errorf("For %s and %+v, expected %d, but got %d instead.", pair.val, pair.slice, pair.index, v)
        }
    }
}

type testSliceIndexInt struct {
    val   int
    slice []int
    index int
}

var (
    intSlice1 []int = []int{1, 3, 5, 7}
    intSlice2 []int = []int{2, 4, 6, 8}
    testsSliceIndexInt = []testSliceIndexInt{
        {1, intSlice1, 0},
        {3, intSlice1, 1},
        {5, intSlice1, 2},
        {7, intSlice1, 3},
        {9, intSlice1, 4},
        {0, intSlice1, 4},

        {2, intSlice2, 0},
        {4, intSlice2, 1},
        {6, intSlice2, 2},
        {8, intSlice2, 3},
        {10, intSlice2, 4},
        {0, intSlice2, 4},
    }
)

func TestSliceIndexInt(t *testing.T) {
    for _, pair := range testsSliceIndexInt {
        v := SliceIndex(len(pair.slice), func(i int) bool {
            return pair.slice[i] == pair.val
        })
        if v != pair.index {
            t.Errorf("For %d and %+v, expected %d, but got %d instead.", pair.val, pair.slice, pair.index, v)
        }
    }
}

/** Test SliceIndexFn **/
func TestSliceIndexFn(t *testing.T) {
    for _, pair := range testsSliceIndexStr {
        v := SliceIndex(len(pair.slice), SliceIndexFn(pair.slice, pair.val))
        if v != pair.index {
            pair.val = cards_testing.GetEmptyStringMaybe(pair.val)
            t.Errorf("For %s and %+v, expected %d, but got %d instead.", pair.val, pair.slice, pair.index, v)
        }
    }
}
