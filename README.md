# mcpkg - npm pour MCP servers

Discover, install & manage Model Context Protocol servers.

## Installation

```bash
go install github.com/yourusername/mcpkg@latest
```

## Usage

```bash
# Discover MCP servers
mcpkg search <query>

# Install an MCP server
mcpkg install <server-name>

# List installed servers
mcpkg list

# Update servers
mcpkg update

# Remove a server
mcpkg remove <server-name>
```

## Why mcpkg?

MCP (Model Context Protocol) is exploding in 2025, but the ecosystem is fragmented across GitHub. There's no easy way to discover and install MCP servers.

mcpkg is like npm for MCP servers - a centralized registry and CLI to manage them all.

## Development

```bash
go build -o mcpkg ./cmd/mcpkg
./mcpkg --help
```

## License

MIT
