package replacemod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - replacemod v0.1.2b"
	fmt.Println(test)
	return test
}
