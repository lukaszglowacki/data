package worker

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/lukaszglowacki/data/pkg/counter"
	"github.com/lukaszglowacki/data/pkg/job"
	"github.com/lukaszglowacki/data/pkg/util/log"
)

const timeFormat = "YYYY-MM-DD HH:SS"

type projection interface {
	Save(t time.Time, json string) error
}

// Run worker process
func Run(p projection, jobs []job.Job, done <-chan bool) {
	tickChan := counter.NewCounter(time.Second).C
	for {
		select {
		case t := <-tickChan:
			dt := time.Now()
			res := make(map[string]string)
			for _, wj := range jobs {
				processJob(t, wj, res)
			}

			jsonString, err := json.Marshal(res)
			if err != nil {
				log.Error(err)
			}

			log.Info(`Save data into projection`)
			if err = p.Save(dt, string(jsonString)); err != nil {
				log.Error(err)
			}
		case <-done:
			return
		}
	}
}

func processJob(t uint64, job job.Job, m map[string]string) {
	jn := job.GetTask().GetName()
	key := strings.Replace(jn, " ", "_", -1)
	val := ""
	var err error
	if t%job.GetTick() == 0 {
		log.Info("Execute job:", jn)
		if val, err = job.GetTask().Exec(); err != nil {
			log.Error(err)
		}
	}
	m[key] = val
}

func appendTime(dt time.Time, m map[string]string) {
	m["check_time"] = dt.Format(timeFormat)
}
