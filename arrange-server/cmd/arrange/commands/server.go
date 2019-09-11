package commands

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/hueyjj/arrange-server/app"
)

var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "Run the Arrange server",
	RunE:         serverCmdF,
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.RunE = serverCmdF
}

func serverCmdF(command *cobra.Command, args []string) error {
	return runServer()
}

func runServer() error {
	server := app.NewServer()
	defer server.Shutdown()

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	<-interruptChan

	return nil
}
