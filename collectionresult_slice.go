// Generated by: gen
// TypeWriter: slice
// Directive: +gen on CollectionResult

package main

// CollectionResultSlice is a slice of type CollectionResult. Use it where you would use []CollectionResult.
type CollectionResultSlice []CollectionResult

// Where returns a new CollectionResultSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv CollectionResultSlice) Where(fn func(CollectionResult) bool) (result CollectionResultSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
