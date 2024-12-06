package todo_test

import (
	"io/ioutil"
	"os"
	"testing"
	"UmairAhmedImran/todo"
	"fmt"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.\n", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.\n", taskName, l[0].Task)
	}
	if l[0].Done {
		fmt.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		fmt.Errorf("Completed task  should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
	"New york1",
	"New york2",
	"New york3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Ecpected %q, got %q instead.", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q got %q instead", tasks[2] ,l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New York"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.\n", taskName, l1[0].Task)
	}
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating Temp file: %s", err)
	}
	defer os.Remove(tf.Name())
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list to file: %s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}