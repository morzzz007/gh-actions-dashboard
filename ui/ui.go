package ui

import (
	"github.com/morzzz007/gh-actions-dashboard/data"
	"github.com/morzzz007/gh-actions-dashboard/ui/components"
	"github.com/rivo/tview"
)

func Render(app *tview.Application, WorkflowRuns []data.WorkflowRun) {
	newPrimitive := func(text string, align int) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(align).
			SetText(text)
	}

	grid := tview.NewGrid().
		SetRows(1, 0).
		SetColumns(30, 0, 30).
		AddItem(newPrimitive("Workflow Dashboard", tview.AlignLeft), 0, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("<esc> Exit", tview.AlignRight), 0, 2, 1, 1, 0, 0, false).
		AddItem(components.CreateTable(app, WorkflowRuns), 1, 0, 1, 3, 0, 0, true)

	grid.SetBorderPadding(0, 0, 1, 1)
	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
