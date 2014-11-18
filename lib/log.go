package lib

import (
	"bytes"
	"github.com/Sirupsen/logrus"
	"time"
)

type ConsoleFormatter struct {
}

func (cf *ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := bytes.NewBufferString("[ ")
	b.WriteString(entry.Time.Format(time.RFC3339))
	b.WriteString(" ] ")
	b.WriteString(entry.Message)
	b.WriteString("\n")

	return b.Bytes(), nil
}
