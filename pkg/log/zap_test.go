package log

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name string
		args args
	}{
		{"zap_logger", args{level: InfoLevel}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := os.Stdout
			logger := New(writer, tt.args.level)
			defer logger.Sync()

			logger.Infof("Info")
			logger.Error("ERROR")
			logger.Warn("Warn")
			logger.Debug("Debug")
		})
	}
}
