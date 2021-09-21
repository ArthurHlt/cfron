package errors

import (
	"fmt"

	"github.com/orange-cloudfoundry/cfron/clients"
	"github.com/pivotal-cf/brokerapi/v8/domain/apiresponses"
)

func NewErrorResponse(msg string, statusCode int, loggerAction string) *apiresponses.FailureResponse {
	return apiresponses.NewFailureResponse(fmt.Errorf(msg), statusCode, loggerAction)
}

func NewErrorFromClient(errAPi clients.GenericOpenAPIError) error {
	if errAPi.Error() == "" {
		return nil
	}
	return fmt.Errorf("%s: %s", errAPi.Error(), string(errAPi.Body()))
}

type ErrWithStatusCode struct {
	Msg        string
	StatusCode int
}

func (e ErrWithStatusCode) Error() string {
	return e.Msg
}

func NewErrorWithStatusCode(msg string, statusCode int) *ErrWithStatusCode {
	return &ErrWithStatusCode{
		Msg:        msg,
		StatusCode: statusCode,
	}
}
