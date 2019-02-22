package localmod

import (
	"fmt"
)

// FuncTest test for local module import reference
func FuncTest() string {
	const test = "FuncTest() - localmod"
	fmt.Println(test)
	return test
}
