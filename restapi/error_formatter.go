package restapi

import "github.com/gifff/xenelectronic/models"

func formatError(code int64, err error) *models.Error {
	errMsg := err.Error()
	return &models.Error{
		Code:    code,
		Message: &errMsg,
	}
}
