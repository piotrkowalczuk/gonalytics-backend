package cli

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/lib"
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
			services.InitConfig(lib.TrackerConfigConsumer, environment)
			services.InitLogger(services.TrackerConfig.Logger)
			services.InitGeoIP(services.TrackerConfig.GeoIP)
			services.InitCassandra(services.TrackerConfig.Cassandra)
			services.InitKafkaClient("gonalytics-publisher", services.TrackerConfig.Kafka)
			services.InitKafkaPublisher(services.TrackerConfig.Kafka)
			services.InitRepositoryManager(services.Cassandra)

			defer services.GeoIP.Close()
			defer services.Cassandra.Close()
			defer services.KafkaClient.Close()

			services.Logger.Info("Server successfully started on ", services.TrackerConfig.Server.GetAddress())
			http.ListenAndServe(
				services.TrackerConfig.Server.GetAddress(),
				routers.GetRouterV1(),
			)
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
