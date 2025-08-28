package daemon_tasks

import "fmt"

type TaskManager struct {
	taskMap map[string]TaskInterface
}

func (t *TaskManager) RegisterTask(name string, task TaskInterface) error {
	if _, ok := t.taskMap[name]; ok {
		return fmt.Errorf("task %s already exists", name)
	}
	t.taskMap[name] = task
	return nil
}

func (t *TaskManager) Start() error {
	for name, task := range t.taskMap {
		if err := task.Start(); err != nil {
			return fmt.Errorf("task %s start failed: %v", name, err)
		}
	}
	return nil
}
