package job

type Job interface {
	GetTick() uint64
	GetTask() Task
}

type Task interface {
	Exec() (string, error)
	GetName() string
}
