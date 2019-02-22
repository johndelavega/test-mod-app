package replacemod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - replacemod v0.0.1"
	fmt.Println(test)
	return test
}
