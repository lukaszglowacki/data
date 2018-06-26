package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lukaszglowacki/data/pkg/repository"
	"github.com/lukaszglowacki/data/pkg/service"
	"github.com/lukaszglowacki/data/pkg/util/cfg"
	"github.com/lukaszglowacki/data/pkg/util/db"
	"github.com/lukaszglowacki/data/pkg/util/log"
	"github.com/sirupsen/logrus"
)

func parseFlags() {
	flag.Parse()
}

func main() {
	parseFlags()

	// Initialize worker.yml config file
	log.Info(`Initialize worker.yml config file`)
	cfg, err := cfg.Init(`service`)
	failOnError(err)
	log.Get().SetLevel(logrus.DebugLevel)

	log.Info(`Open database instance`)
	db, err := db.GetFromConfig(cfg)
	failOnError(err)
	defer db.Close()

	repo := repository.NewProjection(db)

	log.Info("Starting service")
	s := service.New()
	fmt.Println(cfg.GetInt("SERVICE.PORT"))
	s.Run(gin.Default(), cfg.GetInt("SERVICE.PORT"), repo)
	log.Info("Stopping service")
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}
