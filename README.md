# weather MCP server for JP Region

Go で実装したシンプルな日本向け天気予報取得MCPサーバです。

`get_forecast` ツールを提供し、緯度経度を受け取って Open-Meteo から天気予報を取得し、日本語で整形したテキストを返します。

## Features

- Go + `github.com/modelcontextprotocol/go-sdk` で実装
- `stdio` と `Streamable HTTP` の両方に対応
- `get_coordinates` ツールで地点名から座標を取得
- `get_forecast` ツールを提供
- Open-Meteo の forecast API を利用
- Open-Meteo の geocoding API を利用
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

デフォルトでは `stdio transport` で起動します。

MCP ホストから起動して使う場合は、そのまま `stdio` モードで利用できます。

### Run with HTTP transport

Kubernetes やコンテナ環境で使う場合は、HTTP モードで起動できます。

```bash
MCP_TRANSPORT=http \
MCP_HTTP_HOST=0.0.0.0 \
MCP_HTTP_PORT=8080 \
MCP_HTTP_PATH=/mcp \
go run ./cmd/mcp_server
```

利用可能なエンドポイント:

- MCP endpoint: `http://127.0.0.1:8080/mcp`
- Health check: `http://127.0.0.1:8080/healthz`

## Available Tool

### `get_coordinates`

地点名をもとに、日本国内の候補地点の緯度経度を取得します。

入力:

```json
{
  "query": "東京"
}
```

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
| `MCP_GEOCODING_API_BASE_URL` | `https://geocoding-api.open-meteo.com/v1` | geocoding API のベース URL |
| `MCP_USER_AGENT` | `weather-app/0.1.0` | HTTP リクエスト時の User-Agent |
| `MCP_TRANSPORT` | `stdio` | `stdio` または `http` |
| `MCP_HTTP_HOST` | `0.0.0.0` | HTTP モード時の listen host |
| `MCP_HTTP_PORT` | `8080` | HTTP モード時の listen port |
| `MCP_HTTP_PATH` | `/mcp` | HTTP モード時の MCP endpoint path |
| `PORT` | unset | `MCP_HTTP_PORT` より優先される HTTP listen port |

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

## Kubernetes / Container Usage

`k8s` 上で使う場合は、HTTP モードにして `Service` から到達できるようにするのが扱いやすいです。

例:

```bash
MCP_TRANSPORT=http
MCP_HTTP_HOST=0.0.0.0
MCP_HTTP_PORT=8080
MCP_HTTP_PATH=/mcp
```
