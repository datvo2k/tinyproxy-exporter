# Tinyproxy Exporter

A Prometheus exporter for Tinyproxy metrics. This exporter collects statistics from Tinyproxy and exposes them in a format that can be scraped by Prometheus.

## Features

- Exposes Tinyproxy statistics as Prometheus metrics
- Monitors active connections, total requests, and other proxy metrics
- Easy to deploy and configure

## Installation

```bash
go get github.com/datvo2k/tinyproxy-exporter
```

## Configuration

The exporter requires the following environment variables:

- `TINYPROXY_HOST`: Tinyproxy host address (default: "localhost")
- `TINYPROXY_PORT`: Tinyproxy statistics port (default: 8888)
- `EXPORTER_PORT`: Port to expose Prometheus metrics (default: 9888)

## Usage

```bash
./tinyproxy-exporter
```

## Available Metrics

- `tinyproxy_up`: Shows if tinyproxy is up (1) or down (0)
- `tinyproxy_connections_active`: Number of active connections
- `tinyproxy_requests_total`: Total number of requests handled
- `tinyproxy_bytes_in_total`: Total bytes received
- `tinyproxy_bytes_out_total`: Total bytes sent

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License