package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceAliasMap(t *testing.T) {
	m, err := GetResourceAliasMap()

	assert.NoError(t, err)

	// HACK: since resource definitions are cluster-dependent, we test only particular aliases

	// singular
	assert.Equal(t, "po", m["pod"])
	assert.Equal(t, "rs", m["replicaset"])
	assert.Equal(t, "deploy", m["deployment"])
	assert.Equal(t, "svc", m["service"])

	// plural
	assert.Equal(t, "po", m["pods"])
	assert.Equal(t, "rs", m["replicasets"])
	assert.Equal(t, "deploy", m["deployments"])
	assert.Equal(t, "svc", m["services"])

	// resources that does not have alias
	assert.Equal(t, "", m["secret"])
}
