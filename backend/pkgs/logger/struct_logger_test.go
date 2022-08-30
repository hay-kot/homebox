package logger

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lastWrite = []byte{}

type testLogRecorder struct {
	t *testing.T
}

func (tlr testLogRecorder) Write(p []byte) (n int, err error) {
	lastWrite = p
	return len(p), nil
}

type logEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Props   *Props `json:"properties"`
}

func (lr *logEntry) Unmarshal(t *testing.T, jbytes []byte) {
	err := json.Unmarshal(jbytes, lr)
	if err != nil {
		t.Error(err)
	}
}

func Test_LevelString(t *testing.T) {
	assert.Equal(t, "DEBUG", LevelDebug.String())
	assert.Equal(t, "INFO", LevelInfo.String())
	assert.Equal(t, "ERROR", LevelError.String())
	assert.Equal(t, "FATAL", LevelFatal.String())
	assert.Equal(t, "", LevelOff.String())
}

func Test_NewLogger(t *testing.T) {
	logRecorder := testLogRecorder{t: t}

	logger := New(logRecorder, LevelInfo)
	assert.NotNil(t, logger)
}

func getTestLogger(t *testing.T, level Level) *Logger {
	logRecorder := testLogRecorder{t: t}

	logger := New(logRecorder, level)
	assert.NotNil(t, logger)

	return logger
}

func checkLastEntry(t *testing.T, level Level, message string, props *Props) {
	entry := &logEntry{}
	entry.Unmarshal(t, lastWrite)

	assert.Equal(t, level.String(), entry.Level)
	assert.Equal(t, message, entry.Message)
	assert.Equal(t, props, entry.Props)

}

func Test_LoggerDebug(t *testing.T) {
	lgr := getTestLogger(t, LevelDebug)

	lgr.Debug("Test Debug", Props{"Hello": "World"})
	checkLastEntry(t, LevelDebug, "Test Debug", &Props{"Hello": "World"})

	lastWrite = []byte{}
}

func Test_LoggerInfo(t *testing.T) {
	lgr := getTestLogger(t, LevelInfo)

	lgr.Info("Test Info", Props{"Hello": "World"})
	checkLastEntry(t, LevelInfo, "Test Info", &Props{"Hello": "World"})
	lastWrite = []byte{}

}

func Test_LoggerError(t *testing.T) {
	lgr := getTestLogger(t, LevelError)

	myerror := errors.New("Test Error")

	lgr.Error(myerror, Props{"Hello": "World"})
	checkLastEntry(t, LevelError, "Test Error", &Props{"Hello": "World"})
	lastWrite = []byte{}

}

func Test_LoggerLevelScale(t *testing.T) {
	lgr := getTestLogger(t, LevelInfo)
	lastWrite = []byte{}
	lgr.Debug("Test Debug", Props{"Hello": "World"})

	assert.Equal(t, []byte{}, lastWrite)

	lgr = getTestLogger(t, LevelError)
	lastWrite = []byte{}
	lgr.Info("Test Debug", Props{"Hello": "World"})
	lgr.Debug("Test Debug", Props{"Hello": "World"})

	assert.Equal(t, []byte{}, lastWrite)

	lgr = getTestLogger(t, LevelFatal)

	lgr.Info("Test Debug", Props{"Hello": "World"})
	lgr.Debug("Test Debug", Props{"Hello": "World"})
	lgr.Error(errors.New("Test Error"), Props{"Hello": "World"})

	assert.Equal(t, []byte{}, lastWrite)
}
