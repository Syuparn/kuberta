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
			name:     "delegate to kubectl when long resource name is not specified",
			args:     []string{"create", "ns", "default", "--dry-run=client"},
			expected: "namespace/default created (dry run)\n",
		},
		{
			name:     "raise an error when long resource name is used",
			args:     []string{"create", "namespaces", "default", "--dry-run=client"},
			expected: "ERROR: too long! should be `kubectl create ns default --dry-run=client`\n",
		},
		{
			name:     "delegate to kubectl when long long option is not specified",
			args:     []string{"run", "nginx", "-n", "default", "--image", "nginx", "--dry-run=client"},
			expected: "pod/nginx created (dry run)\n",
		},
		{
			name:     "raise an error when long option is used",
			args:     []string{"run", "nginx", "--namespace", "default", "--image", "nginx", "--dry-run=client"},
			expected: "ERROR: too long! should be `kubectl run nginx -n default --image nginx --dry-run=client`\n",
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
