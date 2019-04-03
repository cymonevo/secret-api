package log

import "fmt"

func Printf(format string, args ...interface{}) {
	Print(format, args)
}

func PrintDetail(flag string, format string, args ...interface{}) {
	Print(fmt.Sprint(start, format, end), flag, args)
}

func Infof(format string, args ...interface{}) {
	Info(format, args)
}

func InfoDetail(flag string, format string, args ...interface{}) {
	Info(fmt.Sprint(start, format, end), flag, args)
}

func Warnf(format string, args ...interface{}) {
	Warn(format, args)
}

func WarnDetail(flag string, format string, args ...interface{}) {
	Warn(fmt.Sprint(start, format, end), flag, args)
}

func Errorf(format string, args ...interface{}) {
	Error(format, args)
}

func ErrorDetail(flag string, format string, args ...interface{}) {
	Error(fmt.Sprint(start, format, end), flag, args)
}
