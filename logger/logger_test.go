package logger

import (
	"testing"
)

func TestMyLogger(t *testing.T) {

	MyLogger.InfoLog.Println("hello")
	MyLogger.OkLog.Println("hello")
	MyLogger.WarnLog.Println("hello")
	MyLogger.StatsLog.Println("hello")
	MyLogger.TraceLog.Println("hello")
	MyLogger.ErrorLog.Println("hello")
	MyLogger.OffLog.Println("hello")
	MyLogger.IgnoreLog.Println("hello")

}
