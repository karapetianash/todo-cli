package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	TaskName    string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Add creates a new item and appends it to the List
func (l *List) Add(name string) {
	t := item{
		TaskName:    name,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete marks an item as completed if it exists
func (l *List) Complete(i int) error {
	ls := *l

	if i < 0 || i > len(ls) {
		return fmt.Errorf("complete method: item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes item from the List if it exists
func (l *List) Delete(i int) error {
	ls := *l

	if i < 0 || i > len(ls) {
		return fmt.Errorf("delete method: item %d does not exist", i)
	}

	*l = append(ls[:i], ls[i+1:]...)
	return nil
}

// Save method encodes the List as JSON and saves it using provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

// Get method opens provided file name, decodes JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
