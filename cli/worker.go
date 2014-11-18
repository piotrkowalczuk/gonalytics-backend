package cli

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	"github.com/spf13/cobra"
)

// NewRunWorkerCommand ...
func NewRunWorkerCommand() *cobra.Command {
	var environment string

	runTrackerCommand := &cobra.Command{
		Use:   "tracker",
		Short: "Runes tracker server",
		Long:  "Tracker is a thin layer between HTTP world and queue.",
		Run: func(cmd *cobra.Command, args []string) {
			services.InitConfig(lib.APIConfigConsumer, environment)
			services.InitLogger(services.APIConfig.Logger)
			services.InitGeoIP(services.APIConfig.GeoIP)
			services.InitCassandra(services.APIConfig.Cassandra)
			services.InitRabbitMQ(services.APIConfig.RabbitMQ)
			services.InitRepositoryManager(services.Cassandra)

			defer services.GeoIP.Close()
			defer services.Cassandra.Close()
			defer services.RabbitMQ.Close()
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
