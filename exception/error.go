package exception

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
	Err     string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s %s", e.Code, e.Message, e.Err)
}
