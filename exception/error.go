package exception

import (
	"fmt"
)

type GetError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     string `json:"err"`
}

func (g *GetError) Error() string {
	return fmt.Sprintf("status code : %d message: %s err: %s", g.Code, g.Message, g.Err)
}
