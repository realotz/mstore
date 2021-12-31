package server

import (
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/internal/service"
	"github.com/realotz/mstore/pkg/cron"
)

// 定时任务
func NewCronServer(c *conf.Data, service *service.CronService) *cron.CronManager {
	manager := cron.New()
	service.Register(manager)
	return manager
}
