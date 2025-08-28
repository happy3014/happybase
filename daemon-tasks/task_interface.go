package daemon_tasks

type TaskInterface interface {
	Init() error
	Start() error
	Stop() error
}
