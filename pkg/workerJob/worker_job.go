package workerjob

import (
	"os/exec"

	"github.com/lukaszglowacki/data/pkg/job"
)

type unmarshaler interface {
	Unmarshal(interface{}) error
}

type Job struct {
	tick uint64
	task Task
}

func (j Job) GetTick() uint64   { return j.tick }
func (j Job) GetTask() job.Task { return j.task }

type Task struct {
	name string
	cmd  string
}

func (t Task) GetName() string { return t.name }

func (t Task) Exec() (string, error) {
	r, err := exec.Command("sh", "-c", t.cmd).Output()
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func NewJob(t uint64, n, c string) Job {
	return Job{
		tick: t,
		task: NewTask(n, c),
	}
}

func NewTask(n, c string) Task {
	return Task{
		name: n,
		cmd:  c,
	}
}

type worker struct {
	Jobs []job.Job
}

func Unmarshal(u unmarshaler) ([]job.Job, error) {
	return readWorkersFromConfig(u)
}

func readWorkersFromConfig(u unmarshaler) ([]job.Job, error) {
	type (
		intjb struct {
			Name    string
			Tick    uint64
			Command string
		}
		intwrk struct {
			Jobs []intjb
		}
	)
	intw := intwrk{}
	w := worker{}
	err := u.Unmarshal(&intw)
	for _, j := range intw.Jobs {
		w.Jobs = append(w.Jobs, NewJob(j.Tick, j.Name, j.Command))
	}
	return w.Jobs, err
}
