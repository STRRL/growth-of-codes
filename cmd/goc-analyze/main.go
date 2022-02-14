package main

import (
	"log"

	"github.com/STRRL/growth-of-codes/pkg/analyze/command"
)

func main() {
	rootCommand, err := command.NewRootCommand()
	if err != nil {
		log.Fatalln(err)
	}
	err = rootCommand.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
