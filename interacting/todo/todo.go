package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents a single todo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of todo items
type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}
	*l = append(*l, t)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		// adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
func (l *List) Print(verbose bool, hideComplete bool) {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
			if hideComplete {
				continue
			}
		}
		// adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		if verbose {
			if t.Done {
				formatted += fmt.Sprintf(" Completed At: %s\n", t.CompletedAt.Format("Mon Jan 2 15:04:05"))
			} else {
				formatted += fmt.Sprintf(" Created At: %s\n", t.CreatedAt.Format("Mon Jan 2 15:04:05"))
			}
		}
	}
	print(formatted)
}

// Complete method marks a todo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	// Adjust the index to account for 0-indexing
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	// Adjust the index to account for 0-indexing

	//this is the way to delete an item from a slice,
	//it takes the slice and appends the first part of the slice
	//to the second part of the slice, effectively deleting the item
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method saves the todo List as JSON to a file
func (l *List) Save(filename string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

// Get method loads a todo List from a JSON file
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
