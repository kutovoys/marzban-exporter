# Marzban Metrics Exporter

Marzban Metrics Exporter is an application designed to collect and export metrics from the [Marzban VPN management panel](https://github.com/Gozargah/Marzban). This exporter enables monitoring of various aspects of the VPN service, such as node status, traffic, system metrics, and user information, making the data available for the Prometheus monitoring system.

## Features

- **Node Status Monitoring**: Tracks the connection status of VPN nodes, indicating whether each node is connected or not.
- **Traffic Metrics**: Provides detailed metrics on uplink and downlink traffic for each VPN node.
- **System Health Monitoring**: Reports on system health indicators such as total and used memory, CPU cores, CPU usage, total number of users, active users, and bandwidth metrics.
- **User Activity Tracking**: Monitors user activity, including data limits, used traffic, lifetime used traffic, and online status within the last 2 minutes.
- **Version and Start Time Information**: Includes core version information and whether the core service has started successfully.
- **Configurable via Environment Variables**: Allows customization and configuration through environment variables, making it easy to adjust settings such as metrics port, update interval, and more.
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

Below is a table of environment variables for configuring the exporter:

| Variable Name       | Required | Default Value              | Description                                                         |
| ------------------- | -------- | -------------------------- | ------------------------------------------------------------------- |
| `MARZBAN_BASE_URL`  | Yes      | `https://your-marzban-url` | URL of the Marzban management panel.                                |
| `MARZBAN_USERNAME`  | Yes      | `your-marzban-username`    | Username for the Marzban panel.                                     |
| `MARZBAN_PASSWORD`  | Yes      | `your-marzban-password`    | Password for the Marzban panel.                                     |
| `METRICS_PORT`      | No       | `9090`                     | Port for the metrics server.                                        |
| `METRICS_PROTECTED` | No       | `false`                    | Enable BasicAuth protection for metrics endpoint.                   |
| `METRICS_USERNAME`  | No       | `metricsUser`              | Username for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `METRICS_PASSWORD`  | No       | `MetricsVeryHardPassword`  | Password for BasicAuth, effective if `METRICS_PROTECTED` is `true`. |
| `UPDATE_INTERVAL`   | No       | `30`                       | Interval (in seconds) for metrics update.                           |
| `TIMEZONE`          | No       | `UTC`                      | Timezone for correct time display.                                  |
| `INACTIVITY_TIME`   | No       | `2`                        | Time (in minutes) to determine user activity.                       |

## How to Run

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

## Contribute

We welcome contributions to the Marzban Metrics Exporter! Whether you're looking to fix bugs, add new features, or improve documentation, your help is appreciated. Here's how you can contribute:

1. **Fork the Repository**: Start by forking the repository to your GitHub account.

2. **Create a New Branch**: Create a branch in your forked repository for your changes. It's a good practice to name your branch something descriptive about the feature or fix you're working on.

3. **Make Your Changes**: Implement your changes, fix, or feature in your branch. Make sure to keep your code clean and well-documented.

4. **Test Your Changes**: Ensure your changes don't break existing functionality by running any tests if available. Consider adding new tests if you're adding a new feature.

5. **Commit Your Changes**: Commit your changes with a clear and concise commit message, explaining what you've done and why.

6. **Submit a Pull Request**: Push your changes to your fork and submit a pull request to the main repository. Provide a description of your changes and any other relevant information for the reviewers.

7. **Respond to Feedback**: Once your pull request is reviewed, there might be feedback or questions from the repository maintainers. Be open to discuss your changes and make any necessary adjustments.

Your contributions are what make the open-source community such an amazing place to learn, inspire, and create. We look forward to your contributions to the Marzban Metrics Exporter project!

**Note**: If you're new to open-source contributions, the GitHub documentation on [Creating a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) is a great place to start.
