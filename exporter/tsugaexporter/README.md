# Tsuga Exporter

| Status                   |                       |
| ------------------------ | --------------------- |
| Stability                | [development]         |
| Supported pipeline types | traces, metrics, logs |
| Distributions            | [contrib]             |

Exports telemetry data to [Tsuga](https://tsuga.com) via OTLP/HTTP.

## Configuration

| Field              | Type     | Default | Description                                        |
| ------------------ | -------- | ------- | -------------------------------------------------- |
| `cluster_id`       | string   | (none)  | **Required.** Your Tsuga cluster identifier.       |
| `api_key`          | string   | (none)  | **Required.** Your Tsuga API key (bearer token).   |
| `timeout`          | duration | `30s`   | HTTP request timeout.                              |
| `sending_queue`    | object   | enabled | [Exporter queue settings][queue].                  |
| `retry_on_failure` | object   | enabled | [Retry settings][retry].                           |

[queue]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#configuration
[retry]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/exporterhelper/README.md#configuration

The exporter sends all signals to:

```
https://intake.<cluster_id>.tsuga.com/api/v1/otlp
```

with `Authorization: Bearer <api_key>` header.

## Example

```yaml
exporters:
  tsuga:
    cluster_id: us-east-1
    api_key: ${env:TSUGA_API_KEY}
    timeout: 30s
    sending_queue:
      enabled: true
      num_consumers: 10
      queue_size: 1000
    retry_on_failure:
      enabled: true
      initial_interval: 5s
      max_interval: 30s
      max_elapsed_time: 300s

service:
  pipelines:
    traces:
      exporters: [tsuga]
    metrics:
      exporters: [tsuga]
    logs:
      exporters: [tsuga]
```
