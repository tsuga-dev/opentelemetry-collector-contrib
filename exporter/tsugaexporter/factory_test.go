// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tsugaexporter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

func TestNewFactory(t *testing.T) {
	f := NewFactory()
	require.NotNil(t, f)
	assert.Equal(t, component.MustNewType("tsuga"), f.Type())
}

func TestCreateDefaultConfig(t *testing.T) {
	cfg := NewFactory().CreateDefaultConfig()
	require.NotNil(t, cfg)
	tsugaCfg, ok := cfg.(*Config)
	require.True(t, ok)
	assert.Equal(t, 30*time.Second, tsugaCfg.Timeout)
}

func TestBuildOTLPConfig(t *testing.T) {
	cfg := &Config{
		ClusterID:   "my-cluster",
		APIKey:      configopaque.String("secret"),
		Timeout:     15 * time.Second,
		QueueConfig: configoptional.Some(exporterhelper.NewDefaultQueueConfig()),
		RetryConfig: configretry.NewDefaultBackOffConfig(),
	}

	otlpCfg := buildOTLPConfig(cfg)

	expectedEndpoint := "https://intake.my-cluster.tsuga.com/api/v1/otlp"
	assert.Equal(t, expectedEndpoint, otlpCfg.TracesEndpoint)
	assert.Equal(t, expectedEndpoint, otlpCfg.MetricsEndpoint)
	assert.Equal(t, expectedEndpoint, otlpCfg.LogsEndpoint)
	assert.Equal(t, 15*time.Second, otlpCfg.ClientConfig.Timeout)

	// Headers is configopaque.MapList with Get(name string) (val String, ok bool)
	authVal, ok := otlpCfg.ClientConfig.Headers.Get("Authorization")
	require.True(t, ok, "Authorization header should be set")
	assert.Equal(t, configopaque.String("Bearer secret"), authVal)
}
