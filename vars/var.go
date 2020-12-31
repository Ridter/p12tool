package vars

import (
	"p12tool/util"
	"sync"
)

var (
	Threads = 100
	DebugMode	bool
	Cert		string
	Pass		string
	File		string
	OutFile		string
	Logger		util.Logger
	CrackedPassword string
	Attempts	int
	ResultsLock	sync.RWMutex
)