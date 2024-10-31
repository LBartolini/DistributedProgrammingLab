package gotodo

import (
	"io"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type TableRender struct {
	writer io.Writer
}

func NewTableRender(writer io.Writer) *TableRender {
	return &TableRender{writer}
}

func (render *TableRender) Render(todos *Todos) error {
	tbl := table.New(render.writer)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	tbl.SetAlignment(table.AlignCenter, table.AlignLeft, table.AlignCenter, table.AlignCenter, table.AlignCenter)

	ts := *todos
	for index, t := range ts.todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Local().Format(time.RFC1123)
			}
		}

		tbl.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Local().Format(time.RFC1123), completedAt)
	}
	tbl.Render()
	return nil
}
