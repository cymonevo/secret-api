package log

import (
	"github.com/fatih/color"
)

var start = "[%s] "
var end = ", Err: %v"

var Print = color.New(color.FgRed, color.Bold).PrintfFunc()
var Info = color.New(color.FgRed, color.Bold).PrintfFunc()
var Warn = color.New(color.FgRed, color.Bold).PrintfFunc()
var Error = color.New(color.FgRed, color.Bold).PrintfFunc()
