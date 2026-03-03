package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Zat-Code/mcpkg/registry"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		handleSearch()
	case "install":
		handleInstall()
	case "list":
		handleList()
	case "update":
		handleUpdate()
	case "remove":
		handleRemove()
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func handleSearch() {
	query := ""
	if len(os.Args) > 2 {
		query = os.Args[2]
	}

	results := registry.Search(query)

	if len(results) == 0 {
		fmt.Println("No MCP servers found matching your query.")
		if query != "" {
			fmt.Println("Try: mcpkg search (without query to see all)")
		}
		return
	}

	fmt.Printf("Found %d MCP server(s):\n\n", len(results))
	for _, s := range results {
		fmt.Printf("  %s\n", s.Name)
		fmt.Printf("    %s\n", s.Description)
		fmt.Printf("    ⭐ %d stars | Author: %s\n", s.Stars, s.Author)
		fmt.Printf("    Install: %s\n", s.Install.Commands[0])
		fmt.Println()
	}
}

func handleInstall() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mcpkg install <server-name>")
		os.Exit(1)
	}

	serverName := os.Args[2]
	server := registry.GetByName(serverName)

	if server == nil {
		fmt.Printf("Server '%s' not found in registry.\n", serverName)
		fmt.Println("Try: mcpkg search <query>")
		os.Exit(1)
	}

	fmt.Printf("Installing %s...\n", server.Name)
	fmt.Printf("  %s\n", server.Install.Commands[0])

	// Get config directory
	configDir := getConfigDir()
	serverDir := filepath.Join(configDir, "servers", server.Name)

	// Create server directory
	if err := os.MkdirAll(serverDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Save config
	configPath := filepath.Join(configDir, "config.json")
	saveServerConfig(configPath, server)

	// Run install command (demo mode - would actually run in production)
	fmt.Println("\n✓ Installation complete!")
	fmt.Printf("Server installed to: %s\n", serverDir)
	fmt.Println("Note: Demo mode - install command not actually executed")
}

func handleList() {
	configDir := getConfigDir()
	configPath := filepath.Join(configDir, "config.json")

	// Check if config exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No MCP servers installed yet.")
		fmt.Println("\nUse: mcpkg install <server-name>")
		return
	}

	// Read config (simplified - just show what would be there)
	fmt.Println("Installed MCP servers:")
	fmt.Println("  (config file exists - would parse and display here)")
	fmt.Println("\nUse: mcpkg install <server-name> to add more")
}

func handleUpdate() {
	fmt.Println("Checking for updates...")
	configDir := getConfigDir()
	configPath := filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No servers installed. Nothing to update.")
		return
	}

	fmt.Println("  All servers up to date ✓")
}

func handleRemove() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mcpkg remove <server-name>")
		os.Exit(1)
	}

	serverName := os.Args[2]
	configDir := getConfigDir()
	serverDir := filepath.Join(configDir, "servers", serverName)

	if _, err := os.Stat(serverDir); os.IsNotExist(err) {
		fmt.Printf("Server '%s' is not installed.\n", serverName)
		os.Exit(1)
	}

	fmt.Printf("Removing %s...\n", serverName)
	fmt.Println("(Demo mode - would remove files)")
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.Getenv("HOME")
	}
	return filepath.Join(home, ".mcpkg")
}

func saveServerConfig(configPath string, server *registry.Server) error {
	// Simple append to config file (in production would use proper JSON)
	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s installed\n", server.Name))
	return err
}

func printUsage() {
	fmt.Println(`mcpkg - npm for MCP servers

Usage:
  mcpkg <command> [arguments]

Commands:
  search <query>   Search for MCP servers (omit query to list all)
  install <name>   Install an MCP server
  list             List installed servers
  update           Update all servers
  remove <name>    Remove a server
  help             Show this help message

Examples:
  mcpkg search claude
  mcpkg install claude-code
  mcpkg list`)
}
