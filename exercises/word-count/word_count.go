// +build !example

package wordcount

const testVersion = 2

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency
