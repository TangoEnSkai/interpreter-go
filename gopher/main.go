package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/TangoEnSkai/interpreter-go/gopher/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Gopher programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
