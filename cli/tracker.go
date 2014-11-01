package cli

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	"github.com/piotrkowalczuk/gonalytics-backend/tracker/routers"
	"github.com/spf13/cobra"
)

// NewRunTrackerCommand ...
func NewRunTrackerCommand() *cobra.Command {
	var environment string

	runTrackerCommand := &cobra.Command{
		Use:   "tracker",
		Short: "Runes tracker server",
		Long:  "Tracker is a thin layer between HTTP world and queue.",
		Run: func(cmd *cobra.Command, args []string) {
			services.InitConfig(environment)
			services.InitLogger(services.Config.Logger)
			services.InitGeoIP(services.Config.GeoIP)
			services.InitCassandra(services.Config.Cassandra)
			services.InitRabbitMQ(services.Config.RabbitMQ)
			services.InitRepositoryManager(services.Cassandra)

			defer services.GeoIP.Close()
			defer services.Cassandra.Close()
			defer services.RabbitMQ.Close()

			beego.AddNamespace(routers.GetNamespaceV1())
			beego.Run(services.Config.Server.GetAddress())
		},
	}

	runTrackerCommand.Flags().StringVarP(
		&environment,
		"environment",
		"E",
		"development",
		"Environment at which the server will be started",
	)

	return runTrackerCommand
}
