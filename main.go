package main

import (
	"fmt"

	anyName "github.com/johndelavega/test-mod-app/replacemod"
)

const _version = "v0.1.2"

func main() {

	fmt.Printf("test-mod-app %s\n", _version)
	fmt.Println("main.go: " + anyName.FuncTest())
}

// Version test export from main
func Version() string {
	return _version
}
