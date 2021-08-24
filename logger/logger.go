package logger

import (
	"log"
	"os"
)

var (
	STDCOLOR = StdColor{
		"\033[97;32m",
		"\033[90;47m",
		"\033[90;33m",
		"\033[97;31m",
		"\033[97;34m",
		"\033[97;35m",
		"\033[97;36m",
		"\033[0m",
	}
)

type myLogger struct {
	InfoLog   *log.Logger
	OkLog     *log.Logger
	WarnLog   *log.Logger
	StatsLog  *log.Logger
	TraceLog  *log.Logger
	NormalLog *log.Logger
	ErrorLog  *log.Logger
	OffLog    *log.Logger
	IgnoreLog *log.Logger
}

type StdColor struct {
	Green   string
	White   string
	Yellow  string
	Red     string
	Blue    string
	Magenta string
	Cyan    string
	Reset   string
}

var MyLogger = myLogger{
	InfoLog:   log.New(os.Stdout, STDCOLOR.Reset+"INFO ", log.Flags()),
	OkLog:     log.New(os.Stdout, STDCOLOR.White+"OK ", log.Flags()),
	WarnLog:   log.New(os.Stdout, STDCOLOR.Yellow+"WARN ", log.Flags()),
	StatsLog:  log.New(os.Stdout, STDCOLOR.Blue+"STATS ", log.Flags()),
	TraceLog:  log.New(os.Stdout, STDCOLOR.Cyan+"TRACE ", log.Flags()),
	NormalLog: log.New(os.Stdout, STDCOLOR.Cyan+"", log.Flags()),
	ErrorLog:  log.New(os.Stderr, STDCOLOR.Red+"ERROR ", log.Flags()),
	OffLog:    log.New(os.Stdout, STDCOLOR.Green+"OFF ", log.Flags()),
	IgnoreLog: log.New(os.Stdout, STDCOLOR.Magenta+"IGNORE ", log.Flags()),
}
