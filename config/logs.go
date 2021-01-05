package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// LogFormatter struct
type LogFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// Format func
func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", timestamp, f.LevelDesc[entry.Level], entry.Message)), nil
}
