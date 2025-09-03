# 🍿 avlogem

A simple Go logging utility for structured and plain logs, with support for deterministic time in tests and caller info for plain logs.

## Features
- Structured logging with JSON output
- `Log(value interface{})` for logging any value
- `LogLine(message string)` logs a message with caller file and line
- `Bunch()` for building up fields and logging them
- Global `NowFunc` for mocking time in tests

## Usage

```go
import "github.com/mekariavre/avytta/avlogem"

// Log a value
avlogem.Log(map[string]interface{}{ "foo": "bar" })

// Log a stack message (includes caller file:line)
avlogem.LogLine("hello world")

// Log multiple fields
avlogem.Bunch().Add("a", 1).Add("b", "two").Log()
```

## Testing
Override `avlogem.NowFunc` in your tests to return a fixed time for deterministic output:

```go
avlogem.NowFunc = func() time.Time { return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) }
```

## License
MIT
