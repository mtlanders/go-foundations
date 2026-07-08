package main

import (
	"errors"
	"fmt"
)

//*******************************************************************

// Abandoning ErrTaskFailed, letting TaskError be the underlying type for clarity and avoiding redundance
//var ErrTaskFailed = errors.New("task failed")

//*******************************************************************

type TaskError struct {
	TaskName string
	Cause    error
}

//*******************************************************************

func (t TaskError) Error() string {
	if t.Cause == nil {
		return "task " + t.TaskName + ": no cause"
	}
	return "task " + t.TaskName + ": " + t.Cause.Error()
}

func (t TaskError) Unwrap() error {
	return t.Cause
}

//*******************************************************************

func runTask(name string, fn func() error) (retval error) {
	taskErr := TaskError{TaskName: name, Cause: nil}
	defer func() { // Panic handling
		if r := recover(); r != nil {
			taskErr.Cause = fmt.Errorf("runTask: (panic) %v", r)
			retval = taskErr
		}
	}()
	err := fn()
	if err != nil { // Regular error
		taskErr.Cause = fmt.Errorf("runTask: %w", err)
		retval = taskErr
		return retval
	}
	return nil // Success
}

//*******************************************************************

func main() {

	// Read from empty map
	err := runTask("TaskPlainError", func() error {
		return errors.New("this is a plain error")
	})
	if err != nil {
		var errType TaskError
		if errors.As(err, &errType) {
			fmt.Printf("*** TaskName: %s\n", errType.TaskName)
		}
	}

	// Explictly panic
	err = runTask("TaskManualPanic", func() error {
		panic("manually panicking")
	})
	if err != nil {
		fmt.Println(err)
	}

	err = runTask("TaskSuccess", func() error {
		return nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("(End) Task success!")
	}
}

//*******************************************************************
