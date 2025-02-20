package xcode

import (
	"context"
	"errors"
	e "github.com/pkg/errors"
)

func CodeFromError(err error) XCode {
	err = e.Cause(err)
	var code XCode
	if errors.As(err, &code) {
		return code
	}

	switch {
	case errors.Is(err, context.Canceled):
		return Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return Deadline
	}

	return ServerErr
}
