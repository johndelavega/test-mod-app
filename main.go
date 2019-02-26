package main

import (
	"fmt"

	golib "github.com/johndelavega/go-lib"
	"github.com/johndelavega/gomod"
	anyName "github.com/johndelavega/test-mod-app/replacemod"
)

const _version = "v0.1.6"

func main() {

	fmt.Printf("test-mod-app %s\n", _version)
	fmt.Println("main.go: " + anyName.FuncTest())
	fmt.Println("main.go: " + golib.FuncTest())
	fmt.Println(fmt.Sprintf("gomod.Gomod(): %s\n", gomod.Gomod()))

}

// Version test export from main
func Version() string {
	return _version
}

// to make this run from another app, say test:
// add to test/go.mod: require github.com/johndelavega/test-mod-app v0.1.3
// go mod tidy
// go.mod gets updated to: require github.com/johndelavega/test-mod-app/replacemod v0.0.0-20190222010842-aed14dd634ea

// john@penguin:~/egd/code/go/test$ go run .
// main.go:6:2: import "github.com/johndelavega/test-mod-app" is a program, not an importable package
