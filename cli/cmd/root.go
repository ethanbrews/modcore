package cmd

import (
	"context"
	"log"
	"modcore/cli/ipc"
	pb "modcore/proto/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ContextValues struct {
	Connection           *pb.ModCoreClient
	ConnectionContext    context.Context
	CancellationFunction context.CancelFunc
	Verbose              bool
}

const defaultContextKey string = "modcore-ctx"

func withValues(ctx context.Context, c *pb.ModCoreClient, verbose bool) context.Context {
	return context.WithValue(ctx, defaultContextKey, ContextValues{
		Connection: c,
		Verbose:    verbose,
	})
}

func getValues(ctx context.Context) (*pb.ModCoreClient, bool) {
	values := ctx.Value(defaultContextKey).(ContextValues)
	return values.Connection, values.Verbose
}

func logConnectionFailed(err error) {
	log.Fatalf("Connection failed: %v", err)
}

var (
	verbose bool
	socket  string
)

var rootCmd = &cobra.Command{
	Use:   "modcore",
	Short: "modcore is a mod management CLI",
	Long:  "modcore lets you manage, install, and launch mods with ease.",

	// Runs before ANY command
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Handle global flags here
		if verbose {

		}

		// Example: validate socket before continuing
		if socket == "" {
			socket = ipc.GetSocketPath()
		}

		conn, err := grpc.NewClient(
			socket,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(ipc.SocketDialer()))
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {
				log.Fatalf("failed to close connection: %v", err)
			}
		}(conn)

		client := pb.NewModCoreClient(conn)

		cmd.SetContext(withValues(cmd.Context(), &client, verbose))

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add persistent (global) flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVar(&socket, "socket", "", "Path or URL to the modcore socket (required)")

	// Attach subcommands
	rootCmd.AddCommand(infoCmd)
	//rootCmd.AddCommand(importCmd)
	//rootCmd.AddCommand(listCmd)
	//rootCmd.AddCommand(launchCmd)
	//rootCmd.AddCommand(optionsCmd)
	//rootCmd.AddCommand(upgradeCmd)
	//rootCmd.AddCommand(installCmd)
}
