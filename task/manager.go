package task

import (
	"github.com/goroutine/template/api"
	"github.com/goroutine/template/log"
	"github.com/robfig/cron/v3"
)

type TaskManager struct {
	tasks []api.Task
	debug bool
}

func NewTaskManager(debug ...bool) *TaskManager {
	return &TaskManager{
		debug: len(debug) > 0 && debug[0],
	}
}

func (t *TaskManager) AddTasks(task ...api.Task) {
	t.tasks = append(t.tasks, task...)
}

func (t *TaskManager) RunTasks() {
	c := cron.New(cron.WithSeconds())

	for _, task := range t.tasks {
		taskk := task
		if _, err := c.AddFunc(taskk.CronString(), func() {
			if t.debug {
				log.Logger.Debug("Running task: ", taskk.Name())
			}
			taskk.Run()
		}); err != nil {
			log.Logger.Error("Could not add auto voice task to cron: ", err)
			return
		}
	}

	c.Start()
}
