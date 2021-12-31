package service

import (
	"context"
	"github.com/realotz/mstore/pkg/cron"
)

func NewCronService()*CronService{
	return &CronService{}
}

func HandlerError(f func(ctx context.Context) error) func() {
	return func() {
		err := f(context.Background())
		_ = err
	}
}

type CronService struct {

}

func (s *CronService) Register(manager *cron.CronManager) {
	manager.Register(
		//cron.NewJob("CreateSettlement", "0 0 1 1 * *", HandlerError(s.CreateSettlement)),
	)
}
