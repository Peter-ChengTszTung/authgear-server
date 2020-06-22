package executors

import (
	"context"

	"github.com/skygeario/skygear-server/pkg/log"
	"github.com/skygeario/skygear-server/pkg/task"
)

type RestoreTaskContext func(context.Context, *task.Context) context.Context

type InMemoryExecutor struct {
	Logger         *log.Logger
	RestoreContext RestoreTaskContext

	tasks map[string]task.Task
}

func NewInMemoryExecutor(loggerFactory *log.Factory, restoreContext RestoreTaskContext) *InMemoryExecutor {
	return &InMemoryExecutor{
		Logger:         loggerFactory.New("task-executor"),
		RestoreContext: restoreContext,
		tasks:          map[string]task.Task{},
	}
}

func (e *InMemoryExecutor) Register(name string, task task.Task) {
	e.tasks[name] = task
}

func (e *InMemoryExecutor) Submit(taskCtx *task.Context, spec task.Spec) {
	ctx := e.RestoreContext(context.Background(), taskCtx)
	task := e.tasks[spec.Name]

	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				e.Logger.WithFields(map[string]interface{}{
					"task_name": spec.Name,
					"error":     rec,
				}).Error("unexpected error occurred when running async task")
			}
		}()

		err := task.Run(ctx, spec.Param)
		if err != nil {
			e.Logger.WithFields(map[string]interface{}{
				"task_name": spec.Name,
				"error":     err,
			}).Error("error occurred when running async task")
		}
	}()
}