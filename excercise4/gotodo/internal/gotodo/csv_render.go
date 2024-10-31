package gotodo

import (
	"fmt"
	"io"
	"time"
)

type CSVRender struct {
	writer io.Writer
}

func NewCSVRender(writer io.Writer) *CSVRender {
	return &CSVRender{writer}
}

func (render *CSVRender) Render(todos *Todos) error {
	fmt.Printf("#,Title,Completed,Created At,Completed At\n")
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
		fmt.Printf("%d,%s,%s,%s,%s\n", index, t.Title, completed, t.CreatedAt.Local().Format(time.RFC1123), completedAt)
	}
	return nil
}
