package avlogem

import (
	"encoding/json"
	"fmt"
	"time"
)

// NowFunc is a global function returning current time. Can be mocked in tests.
var NowFunc = func() time.Time { return time.Now() }

type LogEntry struct {
	Tag     string      `json:"tag,omitempty"`
	Time    string      `json:"time"`
	Content interface{} `json:"content,omitempty"`
}

func log(fields interface{}) {
	entry := LogEntry{
		Tag:     "avlogem",
		Time:    NowFunc().Format(time.RFC3339),
		Content: fields,
	}
	if bts, err := json.Marshal(entry); err == nil {
		fmt.Println(string(bts))
	}
}

// Log logs a value as a field with key "value"
func Log(value interface{}) {
	log(value)
}

// LogLine logs a message as a plain string (no fields), including caller file and line.
func LogLine(message string) {
	// Get caller info
	_, file, line, ok := getCallerInfo()
	var info string
	if ok {
		info = fmt.Sprintf("%s:%d called with: %s", file, line, message)
	} else {
		info = message
	}
	log(info)
}

// BunchItem allows building up fields and then logging them
type BunchItem struct {
	fields map[string]interface{}
}

func (b *BunchItem) Add(key string, value interface{}) *BunchItem {
	b.fields[key] = value
	return b
}

func (b *BunchItem) Log() {
	log(b.fields)
}

// For chaining: avlogem.Bunch().Add(...).Log()
func Bunch() *BunchItem {
	return &BunchItem{fields: make(map[string]interface{})}
}
