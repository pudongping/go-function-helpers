package main

import (
	"github.com/pudongping/go-function-helpers/console"
)

func main() {
	s := "hello world"
	console.Success(s)
	console.Danger(s)
	console.Warning(s)
	console.Info(s)
}
