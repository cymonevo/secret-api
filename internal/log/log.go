package log

import (
	"fmt"
	"os"
)

func setFormat(format string, err bool) string {
	if err {
		format = fmt.Sprint(format, errline)
	}
	return fmt.Sprint(timeline, traceline(), tagline, format, newline)
}

func Print(format string) {
	print(format)
}

func Printf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	printf(setFormat(format, false), args...)
}

func Info(format string) {
	info(format)
}

func Infof(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	infof(setFormat(format, false), args...)
}

func Debug(format string) {
	debug(format)
}

func Debugf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	debugf(setFormat(format, false), args...)
}

func DebugDetail(format string, args ...interface{}) {
	debugf(setFormat(format, true), args...)
}

func Warn(format string) {
	warn(format)
}

func Warnf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	warnf(setFormat(format, false), args...)
}

func WarnDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	warnf(setFormat(format, true), args...)
}

func Error(format string) {
	error(format)
}

func Errorf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	errorf(setFormat(format, false), args...)
}

func ErrorDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	errorf(setFormat(format, true), args...)
}

func Fatal(format string) {
	fatal(format)
	os.Exit(1)
}

func Fatalf(tag, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	fatalf(setFormat(format, false), args...)
	os.Exit(1)
}

func FatalDetail(tag string, format string, args ...interface{}) {
	args = append([]interface{}{tag}, args...)
	fatalf(setFormat(format, true), args...)
	os.Exit(1)
}
