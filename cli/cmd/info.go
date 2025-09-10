package cmd

import (
	"log"
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

		log.Printf("Response: %s ; %s ; %s", resp.ApiVersion, resp.CoreVersion, resp.Build)
	},
}
