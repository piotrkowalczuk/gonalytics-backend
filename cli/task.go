package tasks

import (
	"os"

	"github.com/fatih/color"
)

// Job ...
type Job interface {
	DoIt() int
}

// Available ...
var Available = map[string]interface{}{
	"generator": new(Generator),
}

// Run ...
func Run(name string) {
	if task, exists := Available[name]; exists {
		if task, ok := task.(Job); ok {
			os.Exit(task.DoIt())
		} else {
			color.Red("These task dont implement expected methods")
		}
	}
}
