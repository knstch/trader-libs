package cron

import (
	"context"
	"github.com/knstch/trader-libs/log"
	"github.com/robfig/cron/v3"
)

type Cron interface {
	Run(ctx context.Context) error
	GetInterval() string
	GetName() string
}

func StartCronRunner(ctx context.Context, logger *log.Logger, crons ...Cron) error {
	c := cron.New()

	for _, cron := range crons {
		if err := cron.Run(ctx); err != nil {
			return err
		}

		if _, err := c.AddFunc(cron.GetInterval(), func() {
			err := cron.Run(ctx)
			logger.Error("error running cron", err, log.AddMessage("cron", cron.GetName()))
		}); err != nil {
			return err
		}
	}

	return nil
}
