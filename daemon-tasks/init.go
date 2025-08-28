package daemon_tasks

import "github.com/happy3014/happybase/daemon-tasks/tasks"

var taskManager *TaskManager

func init() {
	taskManager = &TaskManager{taskMap: make(map[string]TaskInterface)}

	taskManager.RegisterTask(tasks.HelloTaskName, tasks.NewHelloTask())
}
