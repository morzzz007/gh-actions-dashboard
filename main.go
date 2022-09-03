package main

import (
	"github.com/morzzz007/gh-actions-dashboard/config"
	"github.com/morzzz007/gh-actions-dashboard/data"
	"github.com/morzzz007/gh-actions-dashboard/ui"
	"github.com/rivo/tview"
)

func main() {
	config, err := config.ParseConfig()
	if err != nil {
		return
	}

	WorkflowRuns := data.GetWorkflowRuns(config.RepoPaths)

	app := tview.NewApplication()
	ui.Render(app, WorkflowRuns)
}
