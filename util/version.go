package util

import (
	"runtime"
	"time"
)

var (
	Version   = "1.0"
	GitCommit = "n/a"
	BuildDate = time.Now().Format("01/02/06")
	GoVersion = runtime.Version()
	Author    = "Evi1cg"
)
