// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tsugaexporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/configopaque"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantErr string
	}{
		{
			name: "valid",
			cfg: &Config{
				ClusterID: "my-cluster",
				APIKey:    "secret",
			},
		},
		{
			name:    "missing cluster_id",
			cfg:     &Config{APIKey: "secret"},
			wantErr: "cluster_id",
		},
		{
			name:    "missing api_key",
			cfg:     &Config{ClusterID: "my-cluster"},
			wantErr: "api_key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestConfigEndpoint(t *testing.T) {
	cfg := &Config{
		ClusterID: "us-east-1",
		APIKey:    configopaque.String("key"),
	}
	assert.Equal(t, "https://intake.us-east-1.tsuga.com/api/v1/otlp", cfg.endpoint())
}
