package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	app := &cobra.Command{Use: "gonalytics"}

	app.AddCommand(NewRunTrackerCommand())
	app.AddCommand(NewRunAPICommand())
	app.AddCommand(NewRunActionsWorkerCommand())
	app.Execute()
}
