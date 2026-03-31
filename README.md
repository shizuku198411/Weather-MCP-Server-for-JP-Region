# weather MCP server for JP Region

Go で実装したシンプルな日本向け天気予報取得MCPサーバです。

`get_forecast` ツールを提供し、緯度経度を受け取って Open-Meteo から天気予報を取得し、日本語で整形したテキストを返します。

## Features

- Go + `github.com/modelcontextprotocol/go-sdk` で実装
- stdio transport で動作
- `get_forecast` ツールを提供
- Open-Meteo の forecast API を利用
- 日本語の予報文に整形して返却

返却例:

```text
2026-03-31:
天気: 霧雨
気温: 13.0-18.0°C
降水量: 1.5 mm
降水確率: 45%
最大風速: 17.3 km/h
```

## Requirements

- Go 1.25+

## Run locally

依存関係を取得したあと、次のコマンドで起動できます。

```bash
go run ./cmd/mcp_server
```

このサーバは stdio transport を使うため、MCP ホストから起動して使う想定です。

## Available Tool

### `get_forecast`

指定した緯度経度の 5 日分の天気予報を取得します。

入力:

```json
{
  "latitude": 35.681236,
  "longitude": 139.767125
}
```

## Configuration

環境変数で設定を上書きできます。

| Name | Default | Description |
| --- | --- | --- |
| `MCP_SERVER_NAME` | `weather` | MCP サーバ名 |
| `MCP_SERVER_VERSION` | `0.1.0` | MCP サーバのバージョン |
| `MCP_API_BASE_URL` | `https://api.open-meteo.com/v1` | forecast API のベース URL |
| `MCP_USER_AGENT` | `weather-app/0.1.0` | HTTP リクエスト時の User-Agent |

## Codex VS Code Extension example

`~/.codex/config.toml` に次のような設定を追加すると、Codex VS Code Extension から利用できます。

```toml
[mcp_servers.weather]
command = "go"
args = ["run", "./cmd/mcp_server"]
cwd = "/path/to/mcp_server"
```

今回の例では `go run ./cmd/mcp_server` を使っているため、`cwd` にはこのリポジトリのルートディレクトリを指定してください。

もしビルド済みバイナリを使う場合は、次のようにバイナリパスを直接指定しても構いません。

```toml
[mcp_servers.weather]
command = "/path/to/bin/mcp_server"
args = []
```
