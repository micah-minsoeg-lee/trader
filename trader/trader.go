package trader

import (
	"context"
	"errors"

	"github.com/micah-minsoeg-lee/trader/api"
	"github.com/micah-minsoeg-lee/trader/config"
)

const Key_error = "error"

type Trader struct {
	api    *api.Api
	upper  uint
	lower  uint
	ticker string
}

func NewTrader(cfg *config.Config) (*Trader, error) {
	if api, err := api.NewApi(cfg.Url(), cfg.Key()); err != nil {
		return nil, err
	} else {
		return &Trader{
			api:    api,
			upper:  cfg.UpperLine(),
			lower:  cfg.LowerLine(),
			ticker: cfg.Ticker(),
		}, nil
	}
}

func (t *Trader) Run(ctx context.Context) error {
	go runTrading(ctx, t)

	<-ctx.Done()
	if err := ctx.Value(Key_error); err != nil {
		return errors.New("invalid signal")
	} else {
		return err.(error)
	}
}
