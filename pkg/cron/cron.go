package cron

import (
	"context"
	"github.com/robfig/cron/v3"
)

type Job struct {
	spec   string
	name   string
	handle func()
	id     cron.EntryID
}

func NewJob(name, spec string, handle func()) *Job {
	return &Job{
		name:   name,
		spec:   spec,
		handle: handle,
	}
}

type CronManager struct {
	c      *cron.Cron
	ctx    context.Context
	cancel func()
	jobs   map[cron.EntryID]*Job
}

func New() *CronManager {
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	ctx, cancel := context.WithCancel(context.Background())
	return &CronManager{
		c:      c,
		jobs:   make(map[cron.EntryID]*Job),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (c *CronManager) Register(jobs ...*Job) *CronManager {
	for _, v := range jobs {
		id, err := c.c.AddFunc(v.spec, v.handle)
		if err != nil {
			panic(err)
		}
		v.id = id
		c.jobs[id] = v
	}
	return c
}

func (c *CronManager) Start(context.Context) error {
	c.c.Start()
	return nil
}

func (c *CronManager) Stop(ctx context.Context) error {
	c.cancel()
	c.c.Stop()
	return nil
}
