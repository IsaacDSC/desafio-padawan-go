package pkg

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {

	expectedSettup := slog.NewJSONHandler(os.Stderr, nil)
	expectedLogger := slog.New(expectedSettup)

	tests := []struct {
		name string
		want *slog.Logger
	}{
		{
			name: "Should be return logger",
			want: expectedLogger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
