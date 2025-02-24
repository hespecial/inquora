package xcode

import (
	"inquora/pkg/xcode/types"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)
	//httpStatus := mapXCodeToHTTPStatus(code.Code())
	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
