package main

import (
	"os"

	"github.com/AdonyasG/go-projects/grocery-list-organizer/interfaces/cli"
	"github.com/AdonyasG/go-projects/grocery-list-organizer/usecases"
)

func main() {
	itemUseCase := usecases.NewItmeUseCase()
	cliHandler := cli.NewCLIHandler(itemUseCase)

	cliHandler.Start(os.Stdin, os.Stdout)
}