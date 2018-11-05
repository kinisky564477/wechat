package core

import (
	"time"
)

// ScheduleTask 定时任务
type ScheduleTask interface {
	// Start 启动服务
	Start(delay ...time.Duration)

	// Stop 停止服务
	Stop()
}

// Task 任务, 返回下次执行的时间间隔
type Task func() time.Duration

// DefaultScheduleServer 默认 access_token 获取服务
type DefaultScheduleServer struct {
	done chan bool
	task Task
}

// NewDefaultScheduleServer 构造默认 Token Server
func NewDefaultScheduleServer() *DefaultScheduleServer {
	return &DefaultScheduleServer{
		done: make(chan bool),
	}
}

// SetTask 设置定时任务
func (t *DefaultScheduleServer) SetTask(task Task) {
	t.task = task
}

// Stop 停止服务
func (t *DefaultScheduleServer) Stop() {
	t.done <- true
}

// Start 启动服务
func (t *DefaultScheduleServer) Start(delay ...time.Duration) {
	go t.start(delay...)
}

/**
* 定时任务
* 不依赖文件系统 或者 DB, 依据每次调用反馈的结果，进行下一次任务定时设置
 */

func (t *DefaultScheduleServer) start(delay ...time.Duration) {
	go func() {
		if delay != nil && len(delay) > 0 {
			time.Sleep(delay[0])
		}

		d := t.task() // 上次任务, 返回下一次的执行周期
		for {
			time.Sleep(d)
			d = t.task()
		}
	}()

	for {
		select {
		case done := <-t.done:
			if done {
				return
			}
		}
	}
}

var _ ScheduleTask = &DefaultScheduleServer{}
