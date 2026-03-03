package registry

// Server represents an MCP server in the registry
type Server struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Stars       int      `json:"stars"`
	Author      string   `json:"author"`
	Tags        []string `json:"tags"`
	Install     Install  `json:"install"`
}

// Install describes how to install the server
type Install struct {
	Type     string   `json:"type"` // "npm", "pip", "binary", "go"
	Commands []string `json:"commands"`
}

// Registry contains all MCP servers
var Registry = []Server{
	{
		Name:        "claude-code",
		Description: "Claude Code - Local AI coding assistant with full context awareness",
		URL:         "https://github.com/anthropics/claude-code",
		Stars:       15420,
		Author:      "Anthropic",
		Tags:        []string{"coding", "ai", "assistant"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @anthropic-ai/claude-code",
			},
		},
	},
	{
		Name:        "filesystem",
		Description: "MCP filesystem server - Read/write local files",
		URL:         "https://github.com/modelcontextprotocol/server-filesystem",
		Stars:       892,
		Author:      "ModelContextProtocol",
		Tags:        []string{"filesystem", "storage"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-filesystem",
			},
		},
	},
	{
		Name:        "brave-search",
		Description: "Brave Search API integration for MCP",
		URL:         "https://github.com/modelcontextprotocol/server-brave-search",
		Stars:       456,
		Author:      "ModelContextProtocol",
		Tags:        []string{"search", "api"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-brave-search",
			},
		},
	},
	{
		Name:        "github",
		Description: "GitHub API integration - Issues, PRs, repos management",
		URL:         "https://github.com/modelcontextprotocol/server-github",
		Stars:       1234,
		Author:      "ModelContextProtocol",
		Tags:        []string{"github", "api", "devtools"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-github",
			},
		},
	},
	{
		Name:        "slack",
		Description: "Slack integration - Send messages, list channels",
		URL:         "https://github.com/modelcontextprotocol/server-slack",
		Stars:       567,
		Author:      "ModelContextProtocol",
		Tags:        []string{"slack", "messaging"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-slack",
			},
		},
	},
	{
		Name:        "postgres",
		Description: "PostgreSQL database operations",
		URL:         "https://github.com/modelcontextprotocol/server-postgres",
		Stars:       789,
		Author:      "ModelContextProtocol",
		Tags:        []string{"database", "postgres"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-postgres",
			},
		},
	},
	{
		Name:        "memory",
		Description: "Knowledge graph / vector memory for AI agents",
		URL:         "https://github.com/modelcontextprotocol/server-memory",
		Stars:       2100,
		Author:      "ModelContextProtocol",
		Tags:        []string{"memory", "knowledge", "vector"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-memory",
			},
		},
	},
	{
		Name:        "puppeteer",
		Description: "Browser automation via Puppeteer",
		URL:         "https://github.com/modelcontextprotocol/server-puppeteer",
		Stars:       678,
		Author:      "ModelContextProtocol",
		Tags:        []string{"browser", "automation", "puppeteer"},
		Install: Install{
			Type: "npm",
			Commands: []string{
				"npm install -g @modelcontextprotocol/server-puppeteer",
			},
		},
	},
}

// Search returns servers matching the query
func Search(query string) []Server {
	query = lowercase(query)
	var results []Server
	for _, s := range Registry {
		if contains(lowercase(s.Name), query) ||
			contains(lowercase(s.Description), query) ||
			containsAny(s.Tags, query) {
			results = append(results, s)
		}
	}
	return results
}

func lowercase(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		result[i] = c
	}
	return string(result)
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func containsAny(tags []string, query string) bool {
	for _, tag := range tags {
		if contains(lowercase(tag), query) {
			return true
		}
	}
	return false
}

// GetByName returns a server by name
func GetByName(name string) *Server {
	for i := range Registry {
		if Registry[i].Name == name {
			return &Registry[i]
		}
	}
	return nil
}
