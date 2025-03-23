package errors

import "errors"

var ErrEmptyTasksList = errors.New("empty tasks list")

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task already exists")

var ErrInvalidTask = errors.New("invalid task")

var ErrInvalidStatus = errors.New("invalid status")

var ErrInvalidTaskID = errors.New("invalid task id")

var ErrInvalidRequest = errors.New("invalid request")
