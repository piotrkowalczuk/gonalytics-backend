package cli

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	"github.com/piotrkowalczuk/gonalytics-backend/worker"
	"github.com/spf13/cobra"
)

// NewRunActionsWorkerCommand ...
func NewRunActionsWorkerCommand() *cobra.Command {
	var environment string

	runActionsWorkerCommand := &cobra.Command{
		Use:   "actions-worker",
		Short: "Runes actions worker process",
		Long:  "Actions worker is a process that consume user actions messages.",
		Run: func(cmd *cobra.Command, args []string) {
			services.InitConfig(lib.ActionsWorkerConfigConsumer, environment)
			services.InitLogger(services.ActionsWorkerConfig.Logger)
			services.InitGeoIP(services.ActionsWorkerConfig.GeoIP)
			services.InitCassandra(services.ActionsWorkerConfig.Cassandra)
			services.InitKafkaClient("actions-worker-1", services.ActionsWorkerConfig.Kafka)
			services.InitRepositoryManager(services.Cassandra)

			defer services.GeoIP.Close()
			defer services.Cassandra.Close()

			actionsWorker := worker.ActionsWorker{
				KafkaClient:       services.KafkaClient,
				Config:            services.ActionsWorkerConfig,
				Logger:            services.Logger,
				GeoIP:             services.GeoIP,
				Cassandra:         services.Cassandra,
				RepositoryManager: services.RepositoryManager,
			}

			actionsWorker.Start()
		},
	}

	runActionsWorkerCommand.Flags().StringVarP(
		&environment,
		"environment",
		"E",
		"development",
		"Environment at which the server will be started",
	)

	return runActionsWorkerCommand
}
