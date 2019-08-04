package log

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAdapt := NewAdapterMock(ctrl)
	mockAdapt.EXPECT().Log(
		Fields{
			FieldKeyLevel:   LevelInfo,
			FieldKeyMessage: "test_msg",
		},
	)

	std = &Logger{
		adapter: mockAdapt,
	}
	Log(LevelInfo, "test_msg")
}

func TestConfigure(t *testing.T) {
	ctrl := gomock.NewController(t)
	adapter := NewAdapterMock(ctrl)

	// configure adapter
	Configure(adapter)

	assert.NotNil(t, std)
	assert.Equal(t, std.adapter, adapter)
}

func TestLogFunctions(t *testing.T) {
	cases := []struct {
		level   Level
		logFunc func(string, ...Arg)
	}{
		{LevelTrace, Trace},
		{LevelDebug, Debug},
		{LevelInfo, Info},
		{LevelWarn, Warn},
		{LevelError, Error},
		{LevelPanic, Panic},
		{LevelFatal, Fatal},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.level.String(), func(t *testing.T) {
			t.Parallel()
			testLogFunc(t, c.level, c.logFunc)
		})
	}
}

func testLogFunc(t *testing.T, level Level, logFunc func(string, ...Arg)) {
	const testMsg = "test_msg"
	adapter := NewAdapterMock(gomock.NewController(t))
	adapter.EXPECT().Log(Fields{
		FieldKeyLevel:   level,
		FieldKeyMessage: testMsg,
	})
	std = &Logger{
		adapter: adapter,
	}
	logFunc(testMsg)
}
