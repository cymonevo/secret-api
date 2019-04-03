package log

import (
	"fmt"
	"os"
)

func setFormat(format string, tag bool, err bool) string {
	if tag {
		format = fmt.Sprint(tagline, format)
	}
	if err {
		format = fmt.Sprint(format, errline)
	}
	return fmt.Sprint(format, newline)
}

func setLine(format string) string {
	return fmt.Sprint(format, newline)
}

func Print(format string) {
	print(format)
}

func Printf(format string, args ...interface{}) {
	printf(setFormat(format, false, false), args)
}

func PrintDetail(flag string, format string, args ...interface{}) {
	printf(setFormat(format, true, false), flag, args)
}

func Info(format string) {
	info(format)
}

func Infof(format string, args ...interface{}) {
	infof(setFormat(format, false, false), args)
}

func InfoDetail(flag string, format string, args ...interface{}) {
	infof(setFormat(format, true, false), flag, args)
}

func Debug(format string) {
	debug(format)
}

func Debugf(format string, args ...interface{}) {
	debugf(setFormat(format, false, false), args)
}

func DebugDetail(flag string, format string, args ...interface{}) {
	debugf(setFormat(format, true, false), flag, args)
}

func Warn(format string) {
	warn(format)
}

func Warnf(format string, args ...interface{}) {
	warnf(setFormat(format, false, true), args)
}

func WarnDetail(flag string, format string, args ...interface{}) {
	warnf(setFormat(format, true, true), flag, args)
}

func Error(format string) {
	error(format)
}

func Errorf(format string, args ...interface{}) {
	errorf(setFormat(format, false, true), args)
}

func ErrorDetail(flag string, format string, args ...interface{}) {
	errorf(setFormat(format, true, true), flag, args)
}

func Fatal(format string) {
	fatal(format)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	fatalf(setFormat(format, false, true), args)
	os.Exit(1)
}

func FatalDetail(flag string, format string, args ...interface{}) {
	fatalf(setFormat(format, true, true), flag, args)
	os.Exit(1)
}
