package tasksService

import "fmt"

var ErrTaskNoFound = fmt.Errorf("task not found")
var ErrInvalidInput = fmt.Errorf("task has no title")
