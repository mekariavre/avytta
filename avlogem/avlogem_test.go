package avlogem_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/mekariavre/avytta/avlogem"
	"github.com/stretchr/testify/assert"
)

type SampleData struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

var (
	NowFunc = func() (t time.Time) { t, _ = time.Parse(time.RFC3339, "2023-01-01T00:00:00Z"); return }
	Sample  = SampleData{Foo: "foo", Bar: 42}
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestLog(t *testing.T) {
	avlogem.NowFunc = NowFunc
	output := captureStdout(func() {
		avlogem.Log(map[string]interface{}{"foo": "bar"})
	})

	expected := `{"tag":"avlogem","time":"2023-01-01T00:00:00Z","data":{"foo":"bar"}}` + "\n"
	assert.Equal(t, expected, output)
}

func TestLogLine(t *testing.T) {
	avlogem.NowFunc = NowFunc
	output := captureStdout(func() {
		avlogem.LogLine("hello world")
	})

	expected := `{"tag":"avlogem","time":"2023-01-01T00:00:00Z","data":"hello world"}` + "\n"
	assert.Equal(t, expected, output)
}

func TestBunch(t *testing.T) {
	avlogem.NowFunc = NowFunc
	output := captureStdout(func() {
		avlogem.Bunch().
			Add("okay", 1).
			Add("testing data", "two").
			Add("sample", Sample).
			Add("sample.pointer", &Sample).
			Add("nullish", nil).
			Log()
	})

	expected := `{"tag":"avlogem","time":"2023-01-01T00:00:00Z","data":{"nullish":null,"okay":1,"sample":{"foo":"foo","bar":42},"sample.pointer":{"foo":"foo","bar":42},"testing data":"two"}}` + "\n"
	assert.Equal(t, expected, output)
}
