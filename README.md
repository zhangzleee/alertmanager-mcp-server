# Alertmanager MCP Server

This project provides a server that interacts with Alertmanager through the Model Context Protocol (MCP). It allows you to perform actions such as retrieving alerts and managing silences.

## Features

- **Get Alerts**: Retrieve a list of current alerts from Alertmanager.
- **Get Silences**: Fetch the list of active silences.
- **Set Silences**: Create new silences for specific alerts.

## Getting Started

To get the server up and running, follow these steps.

### Prerequisites

- Go 1.24.2 or higher
- A running instance of Alertmanager

### Installation & Running

1. **Clone the repository:**

    ```bash
    git clone https://github.com/zhangzleee/alertmanager-mcp-server.git
    cd alertmanager-mcp-server
    ```

2. **Build and run the server:**

    You can run the server using the following command. By default, it connects to an Alertmanager instance at `http://localhost:9093`.

    ```bash
    go run ./cmd/main.go server
    ```

## Usage

The server can be configured with the following command-line flags:

- `-host`: The host to listen on (default: `localhost`).
- `-port`: The port to listen on (default: `8000`).
- `-alertmanager-url`: The URL of the Alertmanager instance (default: `http://localhost:9093`).

### Examples

- **Run with default settings:**

  ```bash
  go run ./cmd/main.go server
  ```

- **Run on a different port and host:**

  ```bash
  go run ./cmd/main.go server -port 9000 -host 0.0.0.0
  ```

- **Connect to a different Alertmanager instance:**

  ```bash
  go run ./cmd/main.go server -alertmanager-url http://your-alertmanager-host:9093
  ```

## Tools

This server exposes the following tools via MCP:

### `get_alerts`

Retrieves a list of alerts from Alertmanager. You can filter alerts by their status.

**Parameters:**

- `active`: (string, optional) Filter by active alerts. Default: `false`.
- `silenced`: (string, optional) Filter by silenced alerts. Default: `false`.
- `inhibited`: (string, optional) Filter by inhibited alerts. Default: `false`.
- `unprocessed`: (string, optional) Filter by unprocessed alerts. Default: `false`.

**Example Usage:**

To get all active and silenced alerts:

```json
{
  "tool_name": "get_alerts",
  "parameters": {
    "active": "true",
    "silenced": "true"
  }
}
```

### `get_silences`

Retrieves a list of silences from Alertmanager.

**Parameters:**

- `active`: (string, optional) Filter by active silences. Default: `true`.

**Example Usage:**

To get all active silences:

```json
{
  "tool_name": "get_silences",
  "parameters": {
    "active": "true"
  }
}
```

### `set_silences`

Creates a new silence in Alertmanager.

**Parameters:**

- `comment`: (string, required) A comment to describe the silence.
- `createdBy`: (string, required) The creator of the silence.
- `startsAt`: (string, required) The start time of the silence in RFC3339 format (e.g., `2025-11-07T05:56:07.402Z`).
- `endsAt`: (string, required) The end time of the silence in RFC3339 format (e.g., `2025-11-07T06:56:07.402Z`).
- `matchers`: (array, required) A list of matchers to identify the alerts to be silenced. Each matcher requires `name`, `value`, `isRegex`, and `isEqual` fields.

**Example Usage:**

To set a silence for alerts with the label `severity` equal to `critical`:

```json
{
  "tool_name": "set_silences",
  "parameters": {
    "comment": "Silence critical alerts for maintenance",
    "createdBy": "admin",
    "startsAt": "2025-11-07T05:56:07.402Z",
    "endsAt": "2025-11-07T06:56:07.402Z",
    "matchers": [
      {
        "name": "severity",
        "value": "critical",
        "isRegex": false,
        "isEqual": true
      }
    ]
  }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
