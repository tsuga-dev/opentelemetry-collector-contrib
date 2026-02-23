// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tsugaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tsugaexporter"

import (
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for the Tsuga exporter.
type Config struct {
	// ClusterID is the Tsuga cluster identifier.
	// The target URL is derived as https://intake.<cluster_id>.tsuga.com/api/v1/otlp
	ClusterID string `mapstructure:"cluster_id"`

	// APIKey is the Tsuga API key. Sent as a Bearer token in the Authorization header.
	APIKey configopaque.String `mapstructure:"api_key"`

	// Timeout is the HTTP client timeout for each request.
	Timeout time.Duration `mapstructure:"timeout"`

	// QueueConfig configures the export queue.
	QueueConfig configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	// RetryConfig configures retry-on-failure behaviour.
	RetryConfig configretry.BackOffConfig `mapstructure:"retry_on_failure"`
}

func (c *Config) Validate() error {
	if c.ClusterID == "" {
		return errors.New("cluster_id is required")
	}
	if c.APIKey == "" {
		return errors.New("api_key is required")
	}
	return nil
}

func (c *Config) endpoint() string {
	return fmt.Sprintf("https://intake.%s.tsuga.com/api/v1/otlp", c.ClusterID)
}
