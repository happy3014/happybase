package tasks

import (
	"context"
	"github.com/happy3014/happybase/log"
	"go.uber.org/zap"
	"time"
)

const HelloTaskName = "hello_task"

type HelloTask struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger

	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewHelloTask() *HelloTask {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &HelloTask{
		logger:      log.Logger().With(zap.String("task", HelloTaskName)),
		sugarLogger: log.SugarLogger().With(zap.String("task", HelloTaskName)),
		ctx:         ctx,
		cancelFunc:  cancelFunc,
	}
}

func (h *HelloTask) Init() error {
	h.sugarLogger.Info("hello task init")
	return nil
}

func (h *HelloTask) Start() error {
	h.sugarLogger.Info("hello task start")
	go h.run()
	return nil
}

func (h *HelloTask) Stop() error {
	h.sugarLogger.Info("hello task stop")
	h.cancelFunc()
	return nil
}

func (h *HelloTask) run() {
	h.sugarLogger.Info("hello task run")
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-h.ctx.Done():
			h.sugarLogger.Info("hello task end")
		case <-ticker.C:
			h.sugarLogger.Info("hello task tick")
		}
	}
}
