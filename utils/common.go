package utils

import "github.com/jutimi/grpc-service/common"

func FormatErrorResponse(errCode int32, errMessage string) *common.ErrorResponse {
	return &common.ErrorResponse{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}
}
