package suits

// Get the index of an element from slice
// without requiring sort order
func SliceIndex(limitLen int, predicate func(i int) bool) int {
    for i := 0; i < limitLen; i++ {
        if predicate(i) {
            return i
        }
    }
    return limitLen
}

// Function generator for predicate function use
// with sliceIndex for []string
func SliceIndexFn(slice []string, value interface{}) func(int) bool {
    return func(i int) bool {
        return slice[i] == value
    }
}
