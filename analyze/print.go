package analyze

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func printEndpointsTable(allEndpoints map[string]Endpoint, newmanReportPath string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "METHOD", "PATH", "RESULT", "SOURCE"})

	count := 0
	for _, details := range allEndpoints {
		count++
		method := details.Method
		source := ""
		if details.Result != "Not Covered" {
			source = newmanReportPath
		}
		t.AppendRow(table.Row{count, method, details.Path, details.Result, source})
	}

	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
