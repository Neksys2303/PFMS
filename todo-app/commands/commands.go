package commands

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Update      string
	CreatedAt   time.Time
}

type Todo []Task

func (t *Todo) Add(ID int, Description, Update string) {
	todo := Task{
		ID:          ID,
		Description: Description,
		Update:      Update,
		CreatedAt:   time.Now(),
	}
	*t = append(*t, todo)
}

func (t *Todo) UpdateStatus(ID int) error {
	ls := *t
	if ID <= 0 || ID > len(ls) {
		return errors.New("invalid index")
	}
	ls[ID-1].Update = "Complete"
	ls[ID].Update = "In progress"
	return nil
}

func (t *Todo) Delete(ID int) error {
	ls := *t
	if ID <= 0 || ID > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:ID-1], ls[ID:]...)
	return nil
}

func (t *Todo) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) Storage(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
