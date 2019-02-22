package localmod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - localmod v0.1.0"
	fmt.Println(test)
	return test
}
