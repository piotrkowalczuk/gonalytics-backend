package cli

import (
	"net/http"

	"github.com/piotrkowalczuk/gonalytics-backend/api/routers"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	"github.com/spf13/cobra"
)

// NewRunAPICommand ...
func NewRunAPICommand() *cobra.Command {
	var environment string

	runAPICommand := &cobra.Command{
		Use:   "api",
		Short: "Runes API server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			services.InitConfig(lib.APIConfigConsumer, environment)
			services.InitLogger(services.APIConfig.Logger)
			services.InitCassandra(services.APIConfig.Cassandra)
			services.InitRepositoryManager(services.Cassandra)

			defer services.Cassandra.Close()

			services.Logger.Info("Server successfully started on ", services.APIConfig.Server.GetAddress())
			http.ListenAndServe(
				services.APIConfig.Server.GetAddress(),
				routers.GetRouterV1(),
			)
		},
	}

	runAPICommand.Flags().StringVarP(
		&environment,
		"environment",
		"E",
		"development",
		"Environment at which the server will be started",
	)

	return runAPICommand
}
