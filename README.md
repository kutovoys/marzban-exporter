# Marzban Metrics Exporter

[![GitHub Release](https://img.shields.io/github/v/release/kutovoys/marzban-exporter?style=flat&color=blue)](https://github.com/kutovoys/marzban-exporter/releases/latest)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/kutovoys/marzban-exporter/build-publish.yml)](https://github.com/kutovoys/marzban-exporter/actions/workflows/build-publish.yml)
[![DockerHub](https://img.shields.io/badge/DockerHub-kutovoys%2Fmarzban--exporter-blue)](https://hub.docker.com/r/kutovoys/marzban-exporter/)
[![GitHub License](https://img.shields.io/github/license/kutovoys/marzban-exporter?color=greeen)](https://github.com/kutovoys/marzban-exporter/blob/main/LICENSE)

Marzban Metrics Exporter is an application designed to collect and export metrics from the [Marzban VPN management panel](https://github.com/Gozargah/Marzban). This exporter enables monitoring of various aspects of the VPN service, such as node status, traffic, system metrics, and user information, making the data available for the Prometheus monitoring system.

## Features

- **Node Status Monitoring**: Tracks the connection status of VPN nodes, indicating whether each node is connected or not.
- **Traffic Metrics**: Provides detailed metrics on uplink and downlink traffic for each VPN node.
- **System Health Monitoring**: Reports on system health indicators such as total and used memory, CPU cores, CPU usage, total number of users, active users, and bandwidth metrics.
- **User Activity Tracking**: Monitors user activity, including data limits, used traffic, lifetime used traffic, and online status within the last 2 minutes.
- **Version and Start Time Information**: Includes core version information and whether the core service has started successfully.
- **Configurable via Environment Variables and Command-line Arguments**: Allows customization and configuration through both environment variables and command-line arguments, making it easy to adjust settings.
- **Support for Multiple Architectures**: Docker images are available for multiple architectures, including AMD64 and ARM64, ensuring compatibility across various deployment environments.
- **Optional BasicAuth Protection**: Provides the option to secure the metrics endpoint with BasicAuth, adding an additional layer of security.
- **Integration with Prometheus**: Designed to integrate seamlessly with Prometheus, facilitating easy setup and configuration for monitoring Marzban VPN services.
- **Simplifies VPN Monitoring**: By providing a wide range of metrics, it simplifies the process of monitoring and managing VPN services, enhancing visibility into system performance and user activity.

## Metrics

Below is a table of the metrics provided by Marzban Metrics Exporter:

| Name                                     | Description                                             |
| ---------------------------------------- | ------------------------------------------------------- |
| `marzban_nodes_status`                   | Status of Marzban nodes (1 for connected, 0 otherwise). |
| `marzban_nodes_uplink`                   | Uplink traffic of Marzban nodes.                        |
| `marzban_nodes_downlink`                 | Downlink traffic of Marzban nodes.                      |
| `marzban_user_data_limit`                | Data limit of the user.                                 |
| `marzban_user_used_traffic`              | Used traffic of the user.                               |
| `marzban_user_lifetime_used_traffic`     | Lifetime used traffic of the user.                      |
| `marzban_user_expiration_date`           | User's subscription expiration date                     |
| `marzban_user_online`                    | Whether a user is online within the last 2 minutes.     |
| `marzban_core_started`                   | Indicates if the core is started (1) or not (0).        |
| `marzban_panel_mem_total`                | Total memory.                                           |
| `marzban_panel_mem_used`                 | Used memory.                                            |
| `marzban_panel_cpu_cores`                | Number of CPU cores.                                    |
| `marzban_panel_cpu_usage`                | CPU usage percentage.                                   |
| `marzban_panel_total_users`              | Total number of users.                                  |
| `marzban_panel_users_active`             | Number of active users.                                 |
| `marzban_all_incoming_bandwidth`         | Incoming bandwidth.                                     |
| `marzban_all_outgoing_bandwidth`         | Outgoing bandwidth.                                     |
| `marzban_panel_incoming_bandwidth_speed` | Incoming bandwidth speed.                               |
| `marzban_panel_outgoing_bandwidth_speed` | Outgoing bandwidth speed.                               |

## Configuration

Marzban Metrics Exporter can be configured using environment variables or command-line arguments. When both are provided, command-line arguments take precedence.

Below is a table of configuration options:

| Variable Name       | Command-Line Argument | Required | Default Value                | Description                                                         |
| ------------------- | --------------------- | -------- | ---------------------------- | ------------------------------------------------------------------- |
| `MARZBAN_BASE_URL`  | `--marzban-base-url`  | No       | `https://<your-marzban-url>` | URL of the Marzban management panel.                                |
| `MARZBAN_USERNAME`  | `--marzban-username`  | Yes      | `<your-marzban-username>`    | Username for the Marzban panel.                                     |
| `MARZBAN_PASSWORD`  | `--marzban-password`  | Yes      | `<your-marzban-password>`    | Password for the Marzban panel.                                     |
| `MARZBAN_SOCKET`    | `--marzban-socket`    | No       | `<your-marzban-socket-path>` | Path to Marzban Unix Domain Socket.                                 |
| `METRICS_PORT`      | `--metrics-port`      | No       | `9090`                       | Port for the metrics server.                                        |
| `METRICS_PROTECTED` | `--metrics-protected` | No       | `false`                      | Enable BasicAuth protection for metrics endpoint.                   |
| `METRICS_USERNAME`  | `--metrics-username`  | No       | `metricsUser`                | Username for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `METRICS_PASSWORD`  | `--metrics-password`  | No       | `MetricsVeryHardPassword`    | Password for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `UPDATE_INTERVAL`   | `--update-interval`   | No       | `30`                         | Interval (in seconds) for metrics update.                           |
| `TIMEZONE`          | `--timezone`          | No       | `UTC`                        | Timezone for correct time display.                                  |
| `INACTIVITY_TIME`   | `--inactivity-time`   | No       | `2`                          | Time (in minutes) to determine user activity.                       |

## Usage

### CLI

```bash
/marzban-exporter --marzban-base-url=<your-marzban-panel-url> --marzban-username=<your-marzban-username> --marzban-password=<your-marzban-password>
```

### Docker

```bash
docker run -d \
  -e MARZBAN_BASE_URL=<your-marzban-panel-url> \
  -e MARZBAN_USERNAME=<your-marzban-username> \
  -e MARZBAN_PASSWORD=<your-marzban-password> \
  -p 9090:9090 \
  kutovoys/marzban-exporter
```

### Docker Compose

```yaml
version: "3"
services:
  marzban-exporter:
    image: kutovoys/marzban-exporter
    environment:
      - MARZBAN_BASE_URL=<your-marzban-panel-url>
      - MARZBAN_USERNAME=<your-marzban-username>
      - MARZBAN_PASSWORD=<your-marzban-password>
    ports:
      - "9090:9090"
```

### Integration with Prometheus

To collect metrics with Prometheus, add the exporter to your prometheus.yml configuration file:

```yaml
scrape_configs:
  - job_name: "marzban_exporter"
    static_configs:
      - targets: ["<exporter-ip>:9090"]
```

Ensure to replace `<your-marzban-panel-url>`, `<your-marzban-username>`, `<your-marzban-password>`, and `<exporter-ip>` with your actual information.

## TODO

- ✅ Ensure all necessary environment variables are set and validate them at startup.
- ✅ Implement command line arguments for passing configuration options.
- ⏳ Create a Grafana dashboard tailored for the Marzban Metrics Exporter to visualize the collected metrics effectively.

## Contribute

Contributions to Marzban Metrics Exporter are warmly welcomed. Whether it's bug fixes, new features, or documentation improvements, your input helps make this project better. Here's a quick guide to contributing:

1. **Fork & Branch**: Fork this repository and create a branch for your work.
2. **Implement Changes**: Work on your feature or fix, keeping code clean and well-documented.
3. **Test**: Ensure your changes maintain or improve current functionality, adding tests for new features.
4. **Commit & PR**: Commit your changes with clear messages, then open a pull request detailing your work.
5. **Feedback**: Be prepared to engage with feedback and further refine your contribution.

Happy contributing! If you're new to this, GitHub's guide on [Creating a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) is an excellent resource.

## VPN Recommendation

For secure and reliable internet access, we recommend [BlancVPN](https://getblancvpn.com/?ref=exporter). Use promo code `TRYBLANCVPN` for 15% off your subscription.
