package main

import (
	"github.com/STRRL/growth-of-codes/pkg/persistent/command"
	"log"
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
