package avlogem

import (
	"encoding/json"
	"fmt"
	"time"
)

type LogEntry struct {
	Tag  string      `json:"tag,omitempty"`
	Time string      `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

func log(fields interface{}) {
	entry := LogEntry{
		Tag:  "avlogem",
		Time: time.Now().Format(time.RFC3339),
		Data: fields,
	}
	if bts, err := json.Marshal(entry); err == nil {
		fmt.Println(string(bts))
	}
}

// Log logs a value as a field with key "value"
func Log(value interface{}) {
	log(value)
}

// LogLine logs a message as a plain string (no fields)
func LogLine(message string) {
	log(message)
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
