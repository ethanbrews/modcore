package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Microsoft/go-winio"
)

// getSocketPath returns the appropriate socket path for the current OS
func getSocketPath() string {
	// Get the directory where the executable is located
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	if runtime.GOOS == "windows" {
		// Windows named pipe
		return `\\.\pipe\modcore-socket`
	} else {
		// Unix socket for Linux/macOS
		return filepath.Join(exeDir, "modcore.sock")
	}
}

// createListener creates a listener appropriate for the current OS
func createListener() (net.Listener, error) {
	socketPath := getSocketPath()

	if runtime.GOOS == "windows" {
		// Windows named pipe
		log.Printf("Creating Windows named pipe: %s", socketPath)
		return winio.ListenPipe(socketPath, nil)
	} else {
		// Unix socket
		log.Printf("Creating Unix socket: %s", socketPath)

		// Remove existing socket file if it exists
		if _, err := os.Stat(socketPath); err == nil {
			if err := os.Remove(socketPath); err != nil {
				return nil, err
			}
		}

		return net.Listen("unix", socketPath)
	}
}

// cleanupSocket cleans up socket resources on shutdown
func cleanupSocket() {
	if runtime.GOOS != "windows" {
		socketPath := getSocketPath()
		if err := os.Remove(socketPath); err != nil {
			log.Printf("Warning: Failed to remove socket file: %v", err)
		}
	}
}
