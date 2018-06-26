package systemjob

import "github.com/lukaszglowacki/data/pkg/job"

type fn func() (string, error)

type Job struct {
	tick uint64
	task Task
}

func (j Job) GetTick() uint64   { return j.tick }
func (j Job) GetTask() job.Task { return j.task }

type Task struct {
	name string
	cmd  fn
}

func (t Task) GetName() string { return t.name }

func (t Task) Exec() (string, error) {
	return t.cmd()
}

func NewJob(t uint64, n string, f fn) Job {
	return Job{
		tick: t,
		task: NewTask(n, f),
	}
}

func NewTask(n string, f fn) Task {
	return Task{
		name: n,
		cmd:  f,
	}
}
