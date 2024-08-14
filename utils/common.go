package utils

import (
	"time"

	"github.com/jutimi/grpc-service/common"
	"google.golang.org/grpc/metadata"
)

func FormatErrorResponse(errCode int32, errMessage string) *common.ErrorResponse {
	return &common.ErrorResponse{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}
}

type Metadata struct {
	ServiceName string
}

func GenerateMetaData(data Metadata) metadata.MD {
	return metadata.Pairs(
		"time", time.Now().Format("2006-01-02 15:04:05"),
		"service_name", data.ServiceName,
	)
}
