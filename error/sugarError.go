package error

import "fmt"

type SugarError struct {
	ResponseCode string
	Description  string
	StatusCode   int
}

func (pe *SugarError) Error() string {
	return fmt.Sprintf("response code: %s, error: %s, status code: %d", pe.ResponseCode, pe.Description, pe.StatusCode)
}
