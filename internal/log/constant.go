package log

import (
	"fmt"
	"github.com/fatih/color"
	"path/filepath"
	"runtime"
	"time"
)

const (
	tagline = "[%s] "
	errline = ", Err: %v"
	newline = "\n"
)

var (
	timeline  = time.Now().Format("2006-01-02 15:04:05 ")
	traceline = func() string {
		pc, file, line, ok := runtime.Caller(3)
		if ok {
			file = filepath.Base(file)
			//TODO: Specify GOD mode to view function name below
			_ = filepath.Base(runtime.FuncForPC(pc).Name())
			return fmt.Sprintf("%s:%d ", file, line)
		}
		return "unknown"
	}

	white   = color.New(color.FgWhite)
	green   = color.New(color.FgGreen)
	blue    = color.New(color.FgBlue)
	yellow  = color.New(color.FgYellow)
	red     = color.New(color.FgRed)
	redBold = color.New(color.FgRed, color.Bold)

	print = white.PrintlnFunc()
	info  = green.PrintlnFunc()
	debug = blue.PrintlnFunc()
	warn  = yellow.PrintlnFunc()
	error = red.PrintlnFunc()
	fatal = redBold.PrintlnFunc()

	printf = white.PrintfFunc()
	infof  = green.PrintfFunc()
	debugf = blue.PrintfFunc()
	warnf  = yellow.PrintfFunc()
	errorf = red.PrintfFunc()
	fatalf = redBold.PrintfFunc()
)
