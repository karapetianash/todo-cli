package todo_test

import (
	"os"
	"testing"
	"todo"
)

// TestAdd tests the Add method of the List
func TestAdd(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if len(l) != 3 {
		t.Errorf("expected List length 3, got %d instead.", len(l))
	}

	for i := 0; i < 3; i++ {
		if l[i].TaskName != tasks[i] {
			t.Errorf("expected task %s, got %s instead.", tasks[i], l[i].TaskName)
		}
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)
	//n := len(l)

	// checking whether task added
	if l[0].TaskName != taskName {
		t.Errorf("expected task %s, got %s instead.", taskName, l[0].TaskName)
	}

	if l[0].Done {
		t.Errorf("task should not be completed.")
	}

	err := l.Complete(1)
	if err != nil {
		t.Error(err)
	}

	if !l[0].Done {
		t.Errorf("%s should be completed.", taskName)
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	// deleting the last element (with index 2)
	err := l.Delete(2)
	if err != nil {
		t.Error(err)
	}

	// checking whether any element is deleted
	if len(l) != 2 {
		t.Errorf("expected List length 2, got %d instead.", len(l))
	}

	if l[1].TaskName != tasks[1] {
		t.Errorf("expected %s, got %s instead.", tasks[1], l[1].TaskName)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
	}
	// TODO handle this error
	defer os.Remove(tempFile.Name())

	if err = l1.Save(tempFile.Name()); err != nil {
		t.Errorf("error saving List to file %s", err)
	}

	if err = l2.Get(tempFile.Name()); err != nil {
		t.Errorf("error geting List from file %s", err)
	}

	if l1[0].TaskName != l2[0].TaskName {
		t.Errorf("task %q should match %q", l1[0].TaskName, l2[0].TaskName)
	}
}
