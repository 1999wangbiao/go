// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package textract

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You aren't authorized to perform the action. Use the Amazon Resource Name
	// (ARN) of an authorized user or IAM role to perform the operation.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadDocumentException for service response error code
	// "BadDocumentException".
	//
	// Amazon Textract isn't able to read the document. For more information on
	// the document limits in Amazon Textract, see limits.
	ErrCodeBadDocumentException = "BadDocumentException"

	// ErrCodeDocumentTooLargeException for service response error code
	// "DocumentTooLargeException".
	//
	// The document can't be processed because it's too large. The maximum document
	// size for synchronous operations 10 MB. The maximum document size for asynchronous
	// operations is 500 MB for PDF files.
	ErrCodeDocumentTooLargeException = "DocumentTooLargeException"

	// ErrCodeHumanLoopQuotaExceededException for service response error code
	// "HumanLoopQuotaExceededException".
	//
	// Indicates you have exceeded the maximum number of active human in the loop
	// workflows available
	ErrCodeHumanLoopQuotaExceededException = "HumanLoopQuotaExceededException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// A ClientRequestToken input parameter was reused with an operation, but at
	// least one of the other input parameters is different from the previous call
	// to the operation.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Amazon Textract experienced a service issue. Try your call again.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidJobIdException for service response error code
	// "InvalidJobIdException".
	//
	// An invalid job identifier was passed to an asynchronous analysis operation.
	ErrCodeInvalidJobIdException = "InvalidJobIdException"

	// ErrCodeInvalidKMSKeyException for service response error code
	// "InvalidKMSKeyException".
	//
	// Indicates you do not have decrypt permissions with the KMS key entered, or
	// the KMS key was entered incorrectly.
	ErrCodeInvalidKMSKeyException = "InvalidKMSKeyException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// An input parameter violated a constraint. For example, in synchronous operations,
	// an InvalidParameterException exception occurs when neither of the S3Object
	// or Bytes values are supplied in the Document request parameter. Validate
	// your parameter before calling the API operation again.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidS3ObjectException for service response error code
	// "InvalidS3ObjectException".
	//
	// Amazon Textract is unable to access the S3 object that's specified in the
	// request. for more information, Configure Access to Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html)
	// For troubleshooting information, see Troubleshooting Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/troubleshooting.html)
	ErrCodeInvalidS3ObjectException = "InvalidS3ObjectException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// An Amazon Textract service limit was exceeded. For example, if you start
	// too many asynchronous jobs concurrently, calls to start operations (StartDocumentTextDetection,
	// for example) raise a LimitExceededException exception (HTTP status code:
	// 400) until the number of concurrently running jobs is below the Amazon Textract
	// service limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeProvisionedThroughputExceededException for service response error code
	// "ProvisionedThroughputExceededException".
	//
	// The number of requests exceeded your throughput limit. If you want to increase
	// this limit, contact Amazon Textract.
	ErrCodeProvisionedThroughputExceededException = "ProvisionedThroughputExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Amazon Textract is temporarily unable to process the request. Try your call
	// again.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnsupportedDocumentException for service response error code
	// "UnsupportedDocumentException".
	//
	// The format of the input document isn't supported. Documents for operations
	// can be in PNG, JPEG, PDF, or TIFF format.
	ErrCodeUnsupportedDocumentException = "UnsupportedDocumentException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                  newErrorAccessDeniedException,
	"BadDocumentException":                   newErrorBadDocumentException,
	"DocumentTooLargeException":              newErrorDocumentTooLargeException,
	"HumanLoopQuotaExceededException":        newErrorHumanLoopQuotaExceededException,
	"IdempotentParameterMismatchException":   newErrorIdempotentParameterMismatchException,
	"InternalServerError":                    newErrorInternalServerError,
	"InvalidJobIdException":                  newErrorInvalidJobIdException,
	"InvalidKMSKeyException":                 newErrorInvalidKMSKeyException,
	"InvalidParameterException":              newErrorInvalidParameterException,
	"InvalidS3ObjectException":               newErrorInvalidS3ObjectException,
	"LimitExceededException":                 newErrorLimitExceededException,
	"ProvisionedThroughputExceededException": newErrorProvisionedThroughputExceededException,
	"ThrottlingException":                    newErrorThrottlingException,
	"UnsupportedDocumentException":           newErrorUnsupportedDocumentException,
}