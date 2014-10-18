package cli

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/services"
	"github.com/piotrkowalczuk/gonalytics-backend/tracker/routers"
	"github.com/spf13/cobra"
)

// NewRunTrackerCommand ...
func NewRunTrackerCommand() *cobra.Command {
	var address string

	runTrackerCommand := &cobra.Command{
		Use:   "tracker",
		Short: "Runes tracker server",
		Long:  "Tracker is a thin layer between HTTP world and queue.",
		Run: func(cmd *cobra.Command, args []string) {
			services.InitLogger()
			cassandra := services.InitCassandra("gonalytics", []string{"127.0.0.1"})
			services.InitRepositoryManager(cassandra)

			defer cassandra.Close()

			beego.AddNamespace(routers.GetNamespaceV1())
			beego.Run(address)
		},
	}

	runTrackerCommand.Flags().StringVarP(
		&address,
		"address",
		"A",
		"127.0.0.1:8000",
		"Address and IP at which the server will be started",
	)

	return runTrackerCommand
}
