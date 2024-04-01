package app

import (
	"context"

	"github.com/micah-minsoeg-lee/trader/app/flag"
	"github.com/micah-minsoeg-lee/trader/config"
	"github.com/micah-minsoeg-lee/trader/trader"
)

type App struct {
	trader *trader.Trader
}

func NewApp() (*App, error) {
	a := new(App)

	if cfg, err := config.NewConfig(flag.NewFlag().ConfigFlag()); err != nil {
		return nil, err
	} else if a.trader, err = trader.NewTrader(cfg); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.trader.Run(ctx)
}
