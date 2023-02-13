package errs

import (
	common "github.com/axzed/project-common"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ConvertToGrpcError 将错误转换为grpc的错误
func ConvertToGrpcError(err *BError) error {
	return status.Error(codes.Code(int32(err.Code)), err.Msg)
}

// ParseGrpcError 从grpc的错误中解析出错误码和错误信息
func ParseGrpcError(err error) (code common.BusinessCode, msg string) {
	// 从grpc的错误中解析出错误码和错误信息
	fromError, _ := status.FromError(err)
	// 返回错误码和错误信息
	return common.BusinessCode(fromError.Code()), fromError.Message()
}
