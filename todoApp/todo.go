package todoApp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task          string
	Done          bool
	CreatedTime   time.Time
	CompletedTime time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:          task,
		Done:          false,
		CreatedTime:   time.Now(),
		CompletedTime: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return errors.New("Error occurred index out of bounds")
	}
	ls[index].Done = true
	ls[index].CompletedTime = time.Now()
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}
func (t *Todos) Load(filename string) error {
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

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	fmt.Println("Todo List")
	for i, todo := range *t {
		fmt.Printf("%d. %s \n", i, todo.Task)
	}

}
