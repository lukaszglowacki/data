package main

import (
	"flag"
	"time"

	"github.com/lukaszglowacki/data/pkg/worker"
	"github.com/lukaszglowacki/data/pkg/workerJob"

	"github.com/lukaszglowacki/data/pkg/job"
	systemjob "github.com/lukaszglowacki/data/pkg/systemJob"

	"github.com/lukaszglowacki/data/pkg/repository"
	"github.com/lukaszglowacki/data/pkg/util/cfg"
	"github.com/lukaszglowacki/data/pkg/util/db"
	"github.com/lukaszglowacki/data/pkg/util/log"
)

func parseFlags() {
	flag.Parse()
}

func main() {
	log.Info(`Application starting...`)
	parseFlags()

	log.Info(`Initialize worker.yml config file`)
	cfg, err := cfg.Init(`worker`)
	failOnError(err)

	jobs := []job.Job{}
	// Read jobs from config file
	log.Info(`Read jobs from config file`)

	wj, err := workerjob.Unmarshal(cfg)
	failOnError(err)

	jobs = append(jobs, wj...)
	jobs = append(jobs, systemjob.NewJob(1, "check time", func() (string, error) {
		return time.Now().String(), nil
	}))

	jobs = append(jobs, systemjob.NewJob(1, "hostname", func() (string, error) {
		return cfg.GetString("SERVICE.HOSTNAME"), nil
	}))

	// Open database instance
	log.Info(`Open database instance`)
	db, err := db.GetFromConfig(cfg)
	failOnError(err)
	defer db.Close()

	repo := repository.NewProjection(db)

	cancelChan := make(chan bool)
	defer close(cancelChan)

	// Running worker jobs
	log.Info(`Running worker jobs`)
	worker.Run(repo, jobs, cancelChan)

	log.Info("Application stoping...")
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}
