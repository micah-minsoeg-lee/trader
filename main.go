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
		if err := app.Run(context.Background()); err != nil {
			fmt.Println(err)
			return
		}
	}
}
