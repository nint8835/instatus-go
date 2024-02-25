package instatus_go

import (
	"fmt"
)

const BaseUrl = "https://api.instatus.com"

type Error struct {
	Details struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Details.Code, e.Details.Message)
}
