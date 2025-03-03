package agent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dedbee/Calcserv_Go/pkg/calculation"
)

type Task struct {
	ID            int
	Arg1          float64
	Arg2          float64
	Operation     string
	OperationTime time.Duration
}

func StartAgent() {
	for {
		task, err := getTask()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		result, err := calculation.ArithmeticOperation(task.Arg1, task.Arg2, task.Operation)
		if err != nil {
			continue
		}

		submitTaskResult(task.ID, result)
	}
}

func getTask() (*Task, error) {
	resp, err := http.Get("http://localhost/internal/task")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func submitTaskResult(id int, result float64) error {
	data := map[string]interface{}{
		"id":     id,
		"result": result,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = http.Post("http://localhost/internal/task", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	return nil
}
