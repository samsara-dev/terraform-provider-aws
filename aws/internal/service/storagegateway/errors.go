package storagegateway

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/storagegateway"
)

const (
	OperationErrCodeFileShareNotFound = "FileShareNotFound"
)

// OperationErrorCode returns the operation error code from the specified error:
//  * err is of type awserr.Error and represents a storagegateway.InternalServerError or storagegateway.InvalidGatewayRequestException
//  * Error_ is not nil
// See https://docs.aws.amazon.com/storagegateway/latest/userguide/AWSStorageGatewayAPI.html#APIErrorResponses for details.
func OperationErrorCode(err error) string {
	if inner := (*storagegateway.InternalServerError)(nil); errors.As(err, &inner) && inner.Error_ != nil {
		return aws.StringValue(inner.Error_.ErrorCode)
	}

	if inner := (*storagegateway.InvalidGatewayRequestException)(nil); errors.As(err, &inner) && inner.Error_ != nil {
		return aws.StringValue(inner.Error_.ErrorCode)
	}

	return ""
}
