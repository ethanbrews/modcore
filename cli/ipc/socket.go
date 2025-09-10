package ipc

import (
	"context"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/natefinch/npipe"
)

func GetSocketPath() string {
	// Get the directory where the executable is located
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	if runtime.GOOS == "windows" {
		// Windows named pipe
		return `passthrough:///\\.\pipe\modcore-socket`
	} else {
		// Unix socket for Linux/macOS
		return filepath.Join(exeDir, "modcore.sock")
	}
}

func SocketDialer() func(ctx context.Context, addr string) (net.Conn, error) {
	return func(ctx context.Context, addr string) (net.Conn, error) {
		log.Println("Dialing ", addr)
		timeout := 5 * time.Second
		if deadline, ok := ctx.Deadline(); ok {
			timeout = time.Until(deadline)
		}

		if runtime.GOOS == "windows" {
			// Use npipe for Windows
			return npipe.DialTimeout(addr, timeout)
		}
		// Use Unix domain sockets for everything else
		return net.DialTimeout("unix", addr, timeout)
	}
}
