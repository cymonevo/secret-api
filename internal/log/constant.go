package log

import (
	"github.com/fatih/color"
)

var tagline = "[%s] "
var errline = ", Err: %v"
var newline = "\n"

var white = color.New(color.FgWhite)
var green = color.New(color.FgGreen)
var blue = color.New(color.FgBlue)
var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var redBold = color.New(color.FgRed, color.Bold)

var print = white.PrintlnFunc()
var info = green.PrintlnFunc()
var debug = blue.PrintlnFunc()
var warn = yellow.PrintlnFunc()
var error = red.PrintlnFunc()
var fatal = redBold.PrintlnFunc()

var printf = white.PrintfFunc()
var infof = green.PrintfFunc()
var debugf = blue.PrintfFunc()
var warnf = yellow.PrintfFunc()
var errorf = red.PrintfFunc()
var fatalf = redBold.PrintfFunc()
