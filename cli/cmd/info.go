package cmd

import (
	"fmt"
	pb "modcore/proto/gen"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about modcore or the current setup",
	Run: func(cmd *cobra.Command, args []string) {
		conn, _ := getValues(cmd.Context())

		resp, err := (*conn).CoreInfo(cmd.Context(), &pb.Empty{})
		if err != nil {
			logConnectionFailed(err)
			return
		}
		fmt.Printf("modcore\nApi Version: %s\nCore Version: %s\nBuild: %s\n", resp.ApiVersion, resp.CoreVersion, resp.Build)
	},
}
