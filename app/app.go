package app

import (
	"context"

	"github.com/micah-minsoeg-lee/trader/api/flag"

	"github.com/micah-minsoeg-lee/trader/config"
)

type App struct {
	flag   *flag.Flag
	config *config.Config
}

func NewApp() (*App, error) {
	a := &App{
		flag: flag.NewFlag(),
	}

	return nil, nil
}

func (a *App) Run(ctx context.Context) {

}
