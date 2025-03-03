package orchestrator

import (
	"errors"
	"sync"
	"time"
)

type Expression struct {
	ID     int
	Expr   string
	Status string
	Result float64
}

type Task struct {
	ID            int
	Arg1          float64
	Arg2          float64
	Operation     string
	OperationTime time.Duration
}

var (
	expressions = make(map[int]*Expression)
	tasks       = make(map[int]*Task)
	mutex       = &sync.Mutex{}
	nextID      = 1
)

func AddExpression(expr string) (int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	id := nextID
	nextID++

	expressions[id] = &Expression{
		ID:     id,
		Expr:   expr,
		Status: "pending",
		Result: 0,
	}

	return id, nil
}

func GetExpressions() ([]*Expression, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var exprs []*Expression
	for _, expr := range expressions {
		exprs = append(exprs, expr)
	}

	return exprs, nil
}

func GetExpressionByID(id int) (*Expression, error) {
	mutex.Lock()
	defer mutex.Unlock()

	expr, exists := expressions[id]
	if !exists {
		return nil, errors.New("expression not found")
	}

	return expr, nil
}

func GetTask() (*Task, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, task := range tasks {
		return task, nil
	}

	return nil, errors.New("no tasks available")
}

func SubmitTaskResult(id int, result float64) error {
	mutex.Lock()
	defer mutex.Unlock()

	task, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	expr, exists := expressions[task.ID]
	if !exists {
		return errors.New("expression not found")
	}

	expr.Result = result
	expr.Status = "completed"

	delete(tasks, id)

	return nil
}
