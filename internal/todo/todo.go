package todo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*t = append(*t, todo)
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (t *Todos) Delete(index int) error {
	td := *t
	if err := td.validateIndex(index); err != nil {
		return err
	}

	*t = append(td[:index], td[index+1:]...)

	return nil
}

func (t *Todos) Toggle(index int) error {
	td := *t
	if err := td.validateIndex(index); err != nil {
		return err
	}

	isCompleted := td[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		td[index].CompletedAt = &completionTime
	}

	td[index].Completed = !isCompleted

	return nil
}

func (t *Todos) Edit(index int, title string) error {
	td := *t
	if err := td.validateIndex(index); err != nil {
		return err
	}
	td[index].Title = title

	return nil
}

func (t *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")

	for index, td := range *t {
		completed := "❌"
		completedAt := ""

		if td.Completed {
			completed = "✅"
			if td.CompletedAt != nil {
				completedAt = td.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), td.Title, completed, td.CreatedAt.Format(time.RFC1123), completedAt)

	}
	table.Render()
}
