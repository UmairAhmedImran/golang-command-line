package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task		string
	Done		bool
	CreatedAt	time.Time
	CompletedAt	time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:		task,
		Done:		false,
		CreatedAt:	time.Now(),
		CompletedAt:	time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	fmt.Print(ls)
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)	
	}
	ls[i - 1].Done = true
	ls[i - 1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)	
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Save(filename string) error {
	ls, err := json.Marshal(l)	
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, ls, 0644)
}

func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil	
	}
	return json.Unmarshal(file, l)
}

func (l *List) TodoList() *List {
    newList :=  &List{}
    for _, item := range *l {
        if !item.Done {
            *newList = append(*newList, item)
        }
    }
    return newList
}


func (l *List) CompletedList() *List {
    newList :=  &List{}
    for _, item := range *l {
        if item.Done {
            *newList = append(*newList, item)
        }
    }
    return newList
}


func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task) 
	}
	return formatted
}
