module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tsugaexporter

go 1.25.0

require (
	github.com/stretchr/testify v1.11.1
	go.opentelemetry.io/collector/component v1.52.1-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/component/componenttest v0.146.2-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/config/configopaque v1.52.1-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/config/configoptional v1.52.1-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/config/configretry v1.52.1-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/exporter v1.52.1-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/exporter/exporterhelper v0.146.2-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/exporter/exportertest v0.146.2-0.20260219223409-66996adfaaf7
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.146.2-0.20260219223409-66996adfaaf7
)
