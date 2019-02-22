package localmod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - localmod v0.0.2"
	fmt.Println(test)
	return test
}
