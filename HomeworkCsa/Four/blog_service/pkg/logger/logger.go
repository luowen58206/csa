package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level =  iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)


type Logger struct {

	newLogger *log.Logger
	ctx context.Context
	fields Fields
	callers []string
}

func NewLogger(w io.Writer, prefix string, flag int ) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{
		newLogger: l,
	}
}

func (l *Logger) clone () *Logger {
	nl := *l
	return &nl
}

// WithFields 设置日志公共字段
func (l *Logger)WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger  {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s",file,line,f.Name())}
	}
	return ll
}

// WithCallersFrames 设置当前的整个调用栈信息
func (l *Logger)WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File,
			frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

//定义应用日志的Level 和 fields 的具体类型、并且分为了六个日志等级、便于在不同的场景记录不同级别的日志
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}


//	·编写日志内容的格式化和日志输出动作
func (l *Logger) JSONFormat (level Level, message string) map[string]interface{}{
	data := make(Fields, len(l.fields) + 4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}
func (l *Logger) Output(level Level, message string)  {
	body, _ := json.Marshal(l.JSONFormat(level,message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Print(content)
	case LevelPanic:
		l.newLogger.Print(content)
	}
}

// 	·根据先前定义的日志分级、编写对应的日志输出外部方法
func (l *Logger)Info(v ...interface{})  {
	l.Output(LevelInfo,fmt.Sprint(v...))
}
func (l *Logger)Infof(format string,v ...interface{})  {
	l.Output(LevelInfo,fmt.Sprintf(format,v...))
}
func (l *Logger)Fatal(v ...interface{})  {
	l.Output(LevelFatal,fmt.Sprint(v...))
}
func (l *Logger)Fatalf(format string,v ...interface{})  {
	l.Output(LevelFatal,fmt.Sprintf(format,v...))
}