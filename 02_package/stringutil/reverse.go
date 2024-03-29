// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*
go build
	go build reverse.go reverseTwo.go
 	won't produce an output file.

go install
 	will place the package inside the pkg directory of the workspace.
*/
