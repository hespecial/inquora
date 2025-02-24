package xcode

import (
	"context"
	"errors"
	"fmt"
	e "github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"inquora/pkg/xcode/types"
	"strconv"
)

var _ XCode = (*Status)(nil)

type Status struct {
	sts *types.Status
}

func Error(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func Errorf(code Code, format string, args ...interface{}) *Status {
	code.msg = fmt.Sprintf(format, args...)
	return Error(code)
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}

	return s.sts.Message
}

func (s *Status) Details() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}
	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		// 使用 anypb.UnmarshalNew 动态解包 Any 类型
		detail, err := d.UnmarshalNew()
		if err != nil {
			details = append(details, err)
			continue
		}
		details = append(details, detail)
	}

	return details
}

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

func FromError(err error) *status.Status {
	err = e.Cause(err)
	var code XCode
	if errors.As(err, &code) {
		grpcStatus, err := gRPCStatusFromXCode(code)
		if err == nil {
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	switch {
	case errors.Is(err, context.Canceled):
		grpcStatus, _ = gRPCStatusFromXCode(Canceled)
	case errors.Is(err, context.DeadlineExceeded):
		grpcStatus, _ = gRPCStatusFromXCode(Deadline)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}

func gRPCStatusFromXCode(code XCode) (*status.Status, error) {
	var sts *Status
	switch v := code.(type) {
	case nil:
		return status.New(codes.Unknown, "nil error"), nil
	case *Status:
		sts = v
	case Code:
		sts = FromCode(v)
	default:
		sts = Error(Code{code.Code(), code.Message()})
		for _, detail := range code.Details() {
			if msg, ok := detail.(proto.Message); ok {
				_, _ = sts.WithDetails(msg)
			}
		}
	}

	grpcCode := mapXCodeToGRPCCode(sts.Code())
	stas := status.New(grpcCode, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.Proto())
}

func FromCode(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}

func (s *Status) Proto() *types.Status {
	return s.sts
}

func GrpcStatusToXCode(gstatus *status.Status) XCode {
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			return FromProto(pb)
		}
	}

	return toXCode(gstatus)
}

func FromProto(pbMsg proto.Message) XCode {
	msg, ok := pbMsg.(*types.Status)
	if ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Code{code: int(msg.Code)}
		}
		return &Status{sts: msg}
	}

	return Errorf(ServerErr, "invalid proto message get %v", pbMsg)
}

func toXCode(grpcStatus *status.Status) Code {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return Deadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Internal:
		return ServerErr
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

func mapXCodeToGRPCCode(code int) codes.Code {
	switch code {
	case OK.Code():
		return codes.OK
	case RequestErr.Code():
		return codes.InvalidArgument
	case LimitExceed.Code():
		return codes.ResourceExhausted
	case Unauthorized.Code():
		return codes.Unauthenticated
	case MethodNotAllowed.Code():
		return codes.Unimplemented
	case Deadline.Code():
		return codes.DeadlineExceeded
	case AccessDenied.Code():
		return codes.PermissionDenied
	case ServiceUnavailable.Code():
		return codes.Unavailable
	case NotFound.Code():
		return codes.NotFound
	case ServerErr.Code():
		return codes.Internal
	default:
		return codes.Unknown
	}
}
