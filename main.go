package main

import (
	"context"
	"fmt"

	"github.com/micah-minsoeg-lee/trader/app"
)

func main() {
	if app, err := app.NewApp(); err != nil {
		fmt.Println(err)
		return
	} else {
		app.Run(context.Background())
	}
}
