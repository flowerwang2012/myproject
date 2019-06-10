package test

import (
	"flag"
	"fmt"
)

var logDir = flag.String("log_dir", "./log_dir", "If non-empty, write log files in this directory")

func LogDirTest() {
	fmt.Println(*logDir)
}