package components

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/morzzz007/gh-actions-dashboard/data"
	"github.com/rivo/tview"
)

func getHeadSHA(WorkflowRun data.WorkflowRun) string {
	return fmt.Sprintf("%s…%s", WorkflowRun.HeadSHA[0:4], WorkflowRun.HeadSHA[len(WorkflowRun.HeadSHA)-4:])
}

func getEvent(WorkflowRun data.WorkflowRun) string {
	switch WorkflowRun.Event {
	case "push":
		return "▲"
	case "schedule":
		return "⧗"
	case "workflow_dispatch":
		return "↻"
	default:
		return WorkflowRun.Event
	}
}

func getConclusion(WorkflowRun data.WorkflowRun) string {
	switch WorkflowRun.Conclusion {
	case "success":
		return "✓"
	case "cancelled":
		return "✗"
	case "failure":
		return "✗"
	default:
		return WorkflowRun.Conclusion
	}
}

func getConclusionTableCell(WorkflowRun data.WorkflowRun) *tview.TableCell {
	conclusionText := WorkflowRun.Conclusion
	tCellColor := tcell.ColorWhite

	switch WorkflowRun.Conclusion {
	case "success":
		conclusionText = " ✓"
		tCellColor = tcell.ColorGreen
	case "cancelled":
		conclusionText = " ✗"
		tCellColor = tcell.ColorYellow
	case "failure":
		conclusionText = " ✗"
		tCellColor = tcell.ColorRed
	}

	return tview.NewTableCell(conclusionText).SetTextColor(tCellColor).SetMaxWidth(2)
}

func CreateTable(app *tview.Application, workflows []data.WorkflowRun) *tview.Table {
	table := tview.NewTable()

	table.SetBorder(true)
	table.SetTitle(" Workflows ")

	table.SetCell(0, 0, tview.NewTableCell("").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignRight).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("REPOSITORY").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft).SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft).SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("WORKFLOW").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft).SetSelectable(false))
	table.SetCell(0, 4, tview.NewTableCell("COMMIT").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft).SetSelectable(false))
	table.SetCell(0, 5, tview.NewTableCell("STATUS").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignRight).SetSelectable(false))
	table.SetCell(0, 6, tview.NewTableCell("NUM").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignRight).SetSelectable(false))
	table.SetCell(0, 7, tview.NewTableCell("SHA").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignRight).SetSelectable(false))
	table.SetCell(0, 8, tview.NewTableCell("CREATED").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignRight).SetSelectable(false))

	for k, WorkflowRun := range workflows {
		col := k + 1
		table.SetCell(col, 0, getConclusionTableCell(WorkflowRun).SetAlign(tview.AlignRight))
		table.SetCell(col, 1, tview.NewTableCell(WorkflowRun.Repository.Name).SetAlign(tview.AlignLeft))
		table.SetCell(col, 2, tview.NewTableCell(getEvent(WorkflowRun)).SetAlign(tview.AlignLeft))
		table.SetCell(col, 3, tview.NewTableCell(WorkflowRun.Name).SetAlign(tview.AlignLeft).SetMaxWidth(15))
		table.SetCell(col, 4, tview.NewTableCell(strings.Split(WorkflowRun.HeadCommit.Message, "\n")[0]).SetMaxWidth(50).SetAlign(tview.AlignLeft))
		table.SetCell(col, 5, tview.NewTableCell(WorkflowRun.Status).SetAlign(tview.AlignRight))
		table.SetCell(col, 6, tview.NewTableCell(fmt.Sprintf("%d", WorkflowRun.RunNumber)).SetAlign(tview.AlignRight))
		table.SetCell(col, 7, tview.NewTableCell(getHeadSHA(WorkflowRun)).SetAlign(tview.AlignRight))
		table.SetCell(col, 8, tview.NewTableCell(WorkflowRun.CreatedAt.Format("01/02 15:04")).SetAlign(tview.AlignRight))
	}

	table.Select(0, 0).SetFixed(1, 0).SetSelectable(true, false).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	})

	return table
}
