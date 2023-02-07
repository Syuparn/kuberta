package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValiadateResourceNames(t *testing.T) {
	aliases := map[string]string{
		"pods":        "po",
		"replicasets": "rs",
		"deployments": "deploy",
	}

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "with shortNames",
			args: []string{"get", "po"},
		},
	}

	for _, tt := range tests {
		tt := tt // pin
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateResourceNames(tt.args, aliases)
			assert.NoError(t, err)
		})
	}
}

func TestValiadateResourceNamesError(t *testing.T) {
	aliases := map[string]string{
		"pods":        "po",
		"replicasets": "rs",
		"deployments": "deploy",
		"service":     "svc",
	}

	tests := []struct {
		name          string
		args          []string
		expectedError string
	}{
		{
			name:          "long resource name",
			args:          []string{"get", "pods"},
			expectedError: "too long! should be `kubectl get po`",
		},
		{
			name:          "long resource name before slash",
			args:          []string{"port-forward", "service/myservice", "8443:443"},
			expectedError: "too long! should be `kubectl port-forward svc/myservice 8443:443`",
		},
	}

	for _, tt := range tests {
		tt := tt // pin
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateResourceNames(tt.args, aliases)
			assert.Error(t, err)
			assert.EqualError(t, err, tt.expectedError)
		})
	}
}
