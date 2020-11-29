package commands

import (
	"fmt"

	"github.com/alog-rs/bridge/internal/helpers"
	"github.com/alog-rs/bridge/server"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "bridge",
	Short: "A service to bridge the gap between JAGEX APIs and alog.rs",
	Run:   server.Initialize,
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("Bridge Service :: %s\n", helpers.Version)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Display version information for the bridge service",
	Run:   printVersion,
}

func init() {
	rootCommand.AddCommand(versionCommand)
}

// Execute initializes the bridge service command
func Execute() {
	rootCommand.Execute()
}
