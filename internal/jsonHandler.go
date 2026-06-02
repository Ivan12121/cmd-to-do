package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const saveFile = "tasks.json"

type TaskStore struct {
	Tasks  []Task `json:"tasks"`
	nextId int
}

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
}

func loadTasks() (*TaskStore, error) {
	store := &TaskStore{
		Tasks: []Task{},
		nextId: 1,
	}

	if _, err := os.Stat(saveFile); os.IsNotExist(err) {
		return store, nil
	}

	data, err := os.ReadFile(saveFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		return nil, err
	}

	for _, task := range store.Tasks {
		if task.ID >= store.nextId {
			store.nextId = task.ID + 1
		}
	}
	return store, nil
}

func SaveTasks(store *TaskStore) error {
	data, err := json.MarshalIndent(store, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(saveFile, data, 0644)
}

func AddTask() {
	if len(os.Args) < 3 {
		fmt.Println("Ошибка: укажите название задачи")
        fmt.Println("Пример: todo add Купить молоко")
        return
	}

	title := os.Args[2]

	store, err := loadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
        return
	}

	task := Task{
		ID: store.nextId,
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	store.Tasks = append(store.Tasks, task)
	store.nextId++

	if err := SaveTasks(store); err != nil {
		fmt.Println("Ошибка сохранения:", err)
        return
    }
    
    fmt.Printf("Задача добавлена (ID: %d)\n", task.ID)
}

func ListTasks() {
	store, err := loadTasks()
	
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
        return
	}

	if len(store.Tasks) == 0 {
		fmt.Println("Нет задач. Добавьте первую: todo add <задача>")
        return
	}

	fmt.Println("Список задач:")
	fmt.Println("---")
	for _, task := range store.Tasks {
		status := "X"
		if task.Completed {
			status = "V"
		}
		fmt.Printf("%s [%d] %s\n", status, task.ID, task.Title)
	}
	fmt.Println("---")
}

func CompleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Ошибка: укажите ID задачи")
        return
	}

	var id int
	fmt.Sscanf(os.Args[2], "%d", &id)

	store, err := loadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
        return
	}

	found := false 
	for i := range store.Tasks {
		if store.Tasks[i].ID == id {
			store.Tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Задача с ID %d не найдена\n", id)
        return
	}

    if err := SaveTasks(store); err != nil {
        fmt.Println("Ошибка сохранения:", err)
        return
    }
    
    fmt.Printf("✓ Задача %d отмечена как выполненная\n", id)
}

func DeleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Ошибка: укажите ID задачи")
        return
	}

	var id int
	fmt.Sscanf(os.Args[2], "%d", &id)

	store, err := loadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
        return
	}

	newTasks := []Task{}
	found := false 

	for _, task := range store.Tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Printf("Задача с ID %d не найдена\n", id)
        return
    }
    
    store.Tasks = newTasks
    
    if err := SaveTasks(store); err != nil {
        fmt.Println("Ошибка сохранения:", err)
        return
    }
    
    fmt.Printf("Задача %d удалена\n", id)
}
