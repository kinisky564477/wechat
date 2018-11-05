package core

import (
	"time"
)

// API API 接口信息
type API struct {
	Method       string
	URI          string
	ResponseType string
}

// Kernel 默认核心类
type Kernel struct {
	schedules map[string]ScheduleTask
}

const (
	// ResponseJSON 返回类型 json
	ResponseJSON = "json"

	// ResponseXML 返回类型 xml
	ResponseXML = "xml"
)

var log Log

// NewKernel 初始化
func NewKernel() *Kernel {
	return &Kernel{
		schedules: make(map[string]ScheduleTask),
	}
}

// SetScheduleTask 设置定时任务
func (t *Kernel) SetScheduleTask(id string, schedule ScheduleTask) {
	t.schedules[id] = schedule
}

// SetTask 设置任务
func (t *Kernel) SetTask(id string, task Task) {
	schedule, ok := t.schedules[id]
	if !ok {
		schedule = NewDefaultScheduleServer()
		t.schedules[id] = schedule
	}

	taskServer, _ := schedule.(*DefaultScheduleServer)
	taskServer.SetTask(task)
}

// StartTask 启动
func (t *Kernel) StartTask(id string, delay ...time.Duration) {
	if schedule, ok := t.schedules[id]; ok {
		schedule.Start(delay...)
	}
}

// SetLogInst 设置全局日志实例
func SetLogInst(l Log) {
	log = l
}

func init() {
	log = &DefaultLogger{}
}
