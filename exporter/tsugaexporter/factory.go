// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package tsugaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tsugaexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tsugaexporter/internal/metadata"
)

var otlpHTTPFactory = otlphttpexporter.NewFactory()

// NewFactory creates a factory for the Tsuga exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTraces, metadata.TracesStability),
		exporter.WithMetrics(createMetrics, metadata.MetricsStability),
		exporter.WithLogs(createLogs, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Timeout:     30 * time.Second,
		QueueConfig: configoptional.Some(exporterhelper.NewDefaultQueueConfig()),
		RetryConfig: configretry.NewDefaultBackOffConfig(),
	}
}

func buildOTLPConfig(cfg *Config) *otlphttpexporter.Config {
	otlpCfg := otlpHTTPFactory.CreateDefaultConfig().(*otlphttpexporter.Config)

	endpoint := cfg.endpoint()

	// Override all signal endpoints to the same URL so the upstream exporter
	// does not append /v1/traces, /v1/metrics, /v1/logs suffixes.
	otlpCfg.TracesEndpoint = endpoint
	otlpCfg.MetricsEndpoint = endpoint
	otlpCfg.LogsEndpoint = endpoint

	// Inject bearer token auth using configopaque.MapList.Set
	otlpCfg.ClientConfig.Headers.Set(
		"Authorization",
		configopaque.String("Bearer "+string(cfg.APIKey)),
	)

	otlpCfg.ClientConfig.Timeout = cfg.Timeout
	otlpCfg.QueueConfig = cfg.QueueConfig
	otlpCfg.RetryConfig = cfg.RetryConfig

	return otlpCfg
}

// otlpSettings returns exporter settings with the component ID rewritten to use
// the otlp_http type so the upstream otlphttpexporter factory's type validation
// passes.
func otlpSettings(set exporter.Settings) exporter.Settings {
	set.ID = component.NewIDWithName(otlpHTTPFactory.Type(), set.ID.Name())
	return set
}

func createTraces(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	return otlpHTTPFactory.CreateTraces(ctx, otlpSettings(set), buildOTLPConfig(cfg.(*Config)))
}

func createMetrics(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	return otlpHTTPFactory.CreateMetrics(ctx, otlpSettings(set), buildOTLPConfig(cfg.(*Config)))
}

func createLogs(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	return otlpHTTPFactory.CreateLogs(ctx, otlpSettings(set), buildOTLPConfig(cfg.(*Config)))
}
