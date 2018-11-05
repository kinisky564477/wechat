package core

// Log 日志
type Log interface {
	Critical(string, ...interface{})
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

// DefaultLogger 默认日志实现
type DefaultLogger struct{}

// Critical Critical
func (t *DefaultLogger) Critical(f string, v ...interface{}) {}

// Error Error
func (t *DefaultLogger) Error(f string, v ...interface{}) {}

// Warning Warning
func (t *DefaultLogger) Warning(f string, v ...interface{}) {}

// Info Info
func (t *DefaultLogger) Info(f string, v ...interface{}) {}

// Debug Debug
func (t *DefaultLogger) Debug(f string, v ...interface{}) {}

var _ Log = &DefaultLogger{}
