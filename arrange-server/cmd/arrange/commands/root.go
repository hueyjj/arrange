package commands

import (
	"github.com/spf13/cobra"
)

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "arrange",
	Short: "Streamline commission process",
	Long:  "Arrange reduces time spent organizing commission process, interacting and updating clients, and handling client issues.",
}

func init() {
}
