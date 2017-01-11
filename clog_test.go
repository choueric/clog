package clog_test

import (
	"log/syslog"
	"os"
	"testing"

	"github.com/choueric/clog"
)

func Test_FileLogger(t *testing.T) {
	logFile, err := os.Create("log")
	if err != nil {
		t.Error("create file:", err)
	}
	defer logFile.Close()
	logger := clog.New(logFile, "clog_test:", clog.Lcolor)
	logger.Printf("test")
}

func Test_FileLogger_2nd(t *testing.T) {
	logFile, err := os.Create("log2nd")
	if err != nil {
		t.Error("create file:", err)
	}
	defer logFile.Close()
	defer clog.SetOutput(os.Stderr) // set back to the default.
	clog.SetOutput(logFile)
	logger := clog.New(logFile, "clog_test2nd:", 0)
	logger.Printf("test")
}

func Test_Std(t *testing.T) {
	clog.SetFlags(clog.Ldate | clog.Ltime | clog.Lcolor | clog.Lshortfile)
	clog.SetPrefix("Test_Std: ")
	clog.Printf("string 1")
	clog.Warn("warning message")
	clog.Error("error message")
}

func Test_Syslog(t *testing.T) {
	// Configure logger to write to the syslog.
	sl, err := syslog.New(syslog.LOG_NOTICE, "clogTest")
	if err != nil {
		t.Error("new syslog:", err)
	}
	logger := clog.New(sl, "prefix:", clog.Lshortfile)

	// Now from anywhere else in your program, you can use this:
	logger.Print("Hello Logs!")
	logger.Error("error message!")
}
