# 3X-UI Metrics Exporter

[![GitHub Release](https://img.shields.io/github/v/release/hteppl/3x-ui-exporter?style=flat&color=blue)](https://github.com/kutovoys/marzban-exporter/releases/latest)
[![GitHub License](https://img.shields.io/github/license/kutovoys/marzban-exporter?color=greeen)](https://github.com/kutovoys/marzban-exporter/blob/main/LICENSE)

3X-UI Metrics Exporter is an application designed to collect and export metrics from
the [3X-UI Web Panel](https://github.com/MHSanaei/3x-ui). This exporter enables monitoring of various aspects
of the VPN service, such as node status, traffic, system metrics, and user information, making the data available for
the Prometheus monitoring system.

## Features

- **Online users**: Tracks the total number of online users.
- **Clients up/down**: Tracks total uploaded/downloaded bytes per client.
- **Inbounds up/down**: Tracks total uploaded/downloaded bytes per inbound.
- **XRay version**: Provides XRay version used by 3X-UI.
- **System status**: Tracks metrics about 3X-UI panel resources usage.

## Metrics

Below is a table of the metrics provided by 3X-UI Metrics Exporter.

### Users

Users metrics, such as online:

| Name                      | Description                   |
|---------------------------|-------------------------------|
| `x_ui_total_online_users` | Total number of online users. |

### Clients

Clients metrics (params: `id`, `email`):

| Name                     | Description                        |
|--------------------------|------------------------------------|
| `x_ui_client_up_bytes`   | Total uploaded bytes per client.   |
| `x_ui_client_down_bytes` | Total downloaded bytes per client. |

### Inbounds

Inbounds metrics (params: `id`, `remark`):

| Name                      | Description                         |
|---------------------------|-------------------------------------|
| `x_ui_inbound_up_bytes`   | Total uploaded bytes per inbound.   |
| `x_ui_inbound_down_bytes` | Total downloaded bytes per inbound. |

### System

System metrics (`version` param for `x_ui_xray_version`):

| Name                 | Description                 |
|----------------------|-----------------------------|
| `x_ui_xray_version`  | XRay version used by 3X-UI. |
| `x_ui_panel_threads` | 3X-UI panel threads.        |
| `x_ui_panel_memory`  | 3X-UI panel memory usage.   |
| `x_ui_panel_uptime`  | 3X-UI panel uptime.         |

## Configuration

3X-UI Metrics Exporter can be configured using environment variables or command-line arguments. When both are
provided, command-line arguments take precedence.

Below is a table of configuration options:

| Variable Name       | Command-Line Argument | Required | Default Value              | Description                                                         |
|---------------------|-----------------------|----------|----------------------------|---------------------------------------------------------------------|
| `PANEL_BASE_URL`    | `--panel-base-url`    | Yes      | `https://<your-panel-url>` | URL of the 3X-UI management panel.                                  |
| `PANEL_USERNAME`    | `--panel-username`    | Yes      | `<your-panel-username>`    | Username for the 3X-UI panel.                                       |
| `PANEL_PASSWORD`    | `--panel-password`    | Yes      | `<your-panel-password>`    | Password for the 3X-UI panel.                                       |
| `METRICS_PORT`      | `--metrics-port`      | No       | `9090`                     | Port for the metrics server.                                        |
| `METRICS_PROTECTED` | `--metrics-protected` | No       | `false`                    | Enable BasicAuth protection for metrics endpoint.                   |
| `METRICS_USERNAME`  | `--metrics-username`  | No       | `metricsUser`              | Username for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `METRICS_PASSWORD`  | `--metrics-password`  | No       | `MetricsVeryHardPassword`  | Password for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `UPDATE_INTERVAL`   | `--update-interval`   | No       | `30`                       | Interval (in seconds) for metrics update.                           |
| `TIMEZONE`          | `--timezone`          | No       | `UTC`                      | Timezone for correct time display.                                  |

## Usage

### CLI

```bash
/x-ui-exporter --panel-base-url=<your-panel-url> --panel-username=<your-panel-username> --panel-password=<your-panel-password>
```

### Docker

```bash
indev
```

### Docker Compose

```bash
indev
```

### Integration with Prometheus

To collect metrics with Prometheus, add the exporter to your prometheus.yml configuration file:

```yaml
scrape_configs:
  - job_name: "x-ui_exporter"
    static_configs:
      - targets: [ "<exporter-ip>:9090" ]
```

Ensure to replace `<your-panel-url>`, `<your-panel-username>`, `<your-panel-password>`, and `<exporter-ip>`
with your actual information.

## TODO

- ⏳ Implement more useful metrics.

## Contribute

Contributions to 3X-UI Metrics Exporter are warmly welcomed. Whether it's bug fixes, new features, or documentation
improvements, your input helps make this project better. Here's a quick guide to contributing:

1. **Fork & Branch**: Fork this repository and create a branch for your work.
2. **Implement Changes**: Work on your feature or fix, keeping code clean and well-documented.
3. **Test**: Ensure your changes maintain or improve current functionality, adding tests for new features.
4. **Commit & PR**: Commit your changes with clear messages, then open a pull request detailing your work.
5. **Feedback**: Be prepared to engage with feedback and further refine your contribution.

Happy contributing! If you're new to this, GitHub's guide
on [Creating a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
is an excellent resource.
