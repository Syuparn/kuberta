package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "show help when no args are passed",
			args:     []string{},
			expected: helpMessage,
		},
		{
			name:     "use long resource name",
			args:     []string{"get", "replicasets"},
			expected: "ERROR: too long! should be `kubectl get rs`\n",
		},
	}

	for _, tt := range tests {
		tt := tt // pin

		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Exec(tt.args, &buf)

			assert.Equal(t, tt.expected, buf.String())
		})
	}
}
