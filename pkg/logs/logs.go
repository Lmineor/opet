package logs

import (
	"log"

	"github.com/golang/glog"
)

// GlogWriter serves as a bridge between the standard log package and glog package.
type GlogWriter struct{}

// Write implements the io.Writer interface
func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.InfoDepth(1, string(data))
	return len(data), nil
}

// InitLogs initializes logs the way we want for kubernetes.
func InitLogs() {
	log.SetOutput(GlogWriter{})
	log.SetFlags(0)
}

// FlushLogs flushes logs immediately
func FlushLogs() {
	glog.Flush()
}

// NewLogger creates a new log.Logger which sends logs to glog.Info.
func NewLogger(prefix string) *log.Logger {
	return log.New(GlogWriter{}, prefix, 0)
}