package cli

import "github.com/spf13/cobra"

func init() {
	app := &cobra.Command{Use: "gonalytics"}
	app.AddCommand(NewRunTrackerCommand())
	app.AddCommand(NewRunWorkerCommand())
	app.Execute()
}
