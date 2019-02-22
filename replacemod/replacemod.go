package replacemod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - replacemod"
	fmt.Println(test)
	return test
}
