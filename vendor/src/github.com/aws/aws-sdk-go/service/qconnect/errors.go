// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package qconnect

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request could not be processed because of conflict in the current state
	// of the resource. For example, if you're using a Create API (such as CreateAssistant)
	// that accepts name, a conflicting resource (usually with the same name) is
	// being created or mutated.
	ErrCodeConflictException = "ConflictException"

	// ErrCodePreconditionFailedException for service response error code
	// "PreconditionFailedException".
	//
	// The provided revisionId does not match, indicating the content has been modified
	// since it was last read.
	ErrCodePreconditionFailedException = "PreconditionFailedException"

	// ErrCodeRequestTimeoutException for service response error code
	// "RequestTimeoutException".
	//
	// The request reached the service more than 15 minutes after the date stamp
	// on the request or more than 15 minutes after the request expiration date
	// (such as for pre-signed URLs), or the date stamp on the request is more than
	// 15 minutes in the future.
	ErrCodeRequestTimeoutException = "RequestTimeoutException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You've exceeded your service quota. To perform the requested action, remove
	// some of the relevant resources, or use service quotas to request a service
	// quota increase.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// Amazon Q in Connect throws this exception if you have too many tags in your
	// tag set.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by a service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"PreconditionFailedException":   newErrorPreconditionFailedException,
	"RequestTimeoutException":       newErrorRequestTimeoutException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"TooManyTagsException":          newErrorTooManyTagsException,
	"ValidationException":           newErrorValidationException,
}
