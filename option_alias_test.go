package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOptionAliasMap(t *testing.T) {
	tests := []struct {
		name string
		args []string
		// HACK: since options are version-dependent, we test only particular aliases
		expected map[string]string
	}{
		{
			name: "get command-specific options",
			args: []string{"get", "po"},
			expected: map[string]string{
				"--output":   "-o",
				"--selector": "-l",
				"--watch":    "-w",
			},
		},
		{
			name: "get global options",
			args: []string{"get", "po"},
			expected: map[string]string{
				"--namespace": "-n",
			},
		},
	}

	for _, tt := range tests {
		tt := tt // pin
		t.Run(tt.name, func(t *testing.T) {
			m, err := GetOptionAliasMap(tt.args)

			assert.NoError(t, err)

			for key, expected := range tt.expected {
				assert.Equal(t, expected, m[key])
			}
		})
	}
}
