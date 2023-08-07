package service

import (
	"cron-job/conf"
	"cron-job/handler"
	"cron-job/model"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/robfig/cron/v3"
)

type App struct {
	cronjob    *cron.Cron
	configJobs []*model.JobConfig
}

func NewApp() *App {
	return &App{
		cronjob: cron.New(),
	}
}

func (a *App) registerJobs() error {
	//a.cronjob.AddFunc("@every 0h0m1s", func() { println("Every second") })
	for i, jobCfg := range a.configJobs {
		// fill a name if empty
		if jobCfg.Name == "" {
			jobCfg.Name = fmt.Sprintf("job-%d", i)
		}
		handler := handler.NewHandler(jobCfg)
		_, err := a.cronjob.AddJob(jobCfg.Spec, handler)
		if err != nil {
			return fmt.Errorf("register job error (name: %s): %v", jobCfg.Name, err)
		}
	}
	return nil
}

func (a *App) Run() error {
	config := conf.NewConfig()
	err := env.Parse(config)
	if err != nil {
		return errors.New("error while parsing extra setting: " + err.Error())
	}

	a.configJobs, err = config.LoadConfigJobs()
	if err != nil {
		return fmt.Errorf("load jobs error: %v", err)
	}

	if err = a.registerJobs(); err != nil {
		return fmt.Errorf("failed to register jobs: %v", err)
	}
	a.cronjob.Start()

	var forever chan struct{}
	<-forever

	return nil
}
