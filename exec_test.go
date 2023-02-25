package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expected       string
		expectedStdErr string
		expectedCode   int
	}{
		{
			name:           "show help when no args are passed",
			args:           []string{},
			expected:       helpMessage,
			expectedStdErr: "",
			expectedCode:   0,
		},
		{
			name:           "delegate to kubectl when long resource name is not specified",
			args:           []string{"create", "ns", "default", "--dry-run=client"},
			expected:       "namespace/default created (dry run)\n",
			expectedStdErr: "",
			expectedCode:   0,
		},
		{
			name:           "raise an error when long resource name is used",
			args:           []string{"create", "namespaces", "default", "--dry-run=client"},
			expected:       "",
			expectedStdErr: "ERROR: too long! should be `kubectl create ns default --dry-run=client`\n",
			expectedCode:   1,
		},
		{
			name:           "delegate to kubectl when long long option is not specified",
			args:           []string{"run", "nginx", "-n", "default", "--image", "nginx", "--dry-run=client"},
			expected:       "pod/nginx created (dry run)\n",
			expectedStdErr: "",
			expectedCode:   0,
		},
		{
			name:           "raise an error when long option is used",
			args:           []string{"run", "nginx", "--namespace", "default", "--image", "nginx", "--dry-run=client"},
			expected:       "",
			expectedStdErr: "ERROR: too long! should be `kubectl run nginx -n default --image nginx --dry-run=client`\n",
			expectedCode:   1,
		},
	}

	for _, tt := range tests {
		tt := tt // pin

		t.Run(tt.name, func(t *testing.T) {
			var stdOutBuf bytes.Buffer
			var stdErrBuf bytes.Buffer
			err, code := Exec(tt.args, &stdOutBuf, &stdErrBuf)

			assert.Equal(t, tt.expected, stdOutBuf.String())
			assert.Equal(t, tt.expectedStdErr, stdErrBuf.String())
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, code)
		})
	}
}
