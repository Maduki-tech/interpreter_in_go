package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/maduki-tech/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s this is the monkey lanugage enjoy\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
