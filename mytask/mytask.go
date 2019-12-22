package mytask

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var (
	tasks  []*Task
	nextID int
)

func init() {
	tasks = []*Task{
		&Task{
			1, "Learn Redux", "The store, actions, and reducers, oh my!", "Unstarted",
		},
		&Task{
			2, "Peace on Earth",
			"No big deal.",
			"Unstarted",
		},
		&Task{
			3,
			"Create Facebook for dogs",
			"The hottest new social network",
			"Completed",
		},
	}
	nextID = 4
}

func GetAll() []*Task {
	fmt.Printf("task.len = %d\n", len(tasks))
	return tasks
}

func Add(t *Task) *Task {
	t.ID = nextID
	nextID += 1
	t.Status = "Unstarted"
	tasks = append(tasks, t)
	return t
}

func Edit(id int, t *Task) *Task {
	fmt.Printf("XXX Edit: %d, %+v\n", id, t)

	for _, v := range tasks {
		if v.ID == id {
			if t.Title != "" {
				v.Title = t.Title
			}
			if t.Description != "" {
				v.Description = t.Description
			}
			if t.Status != "" {
				v.Status = t.Status
			}
			return v
		}
		continue
	}
	return nil
}
