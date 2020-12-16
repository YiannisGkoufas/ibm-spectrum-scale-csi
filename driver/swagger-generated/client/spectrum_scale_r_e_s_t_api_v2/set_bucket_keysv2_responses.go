// Code generated by go-swagger; DO NOT EDIT.

package spectrum_scale_r_e_s_t_api_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"example.com/m/v2/models"
)

// SetBucketKeysv2Reader is a Reader for the SetBucketKeysv2 structure.
type SetBucketKeysv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetBucketKeysv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSetBucketKeysv2Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetBucketKeysv2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetBucketKeysv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetBucketKeysv2Accepted creates a SetBucketKeysv2Accepted with default headers values
func NewSetBucketKeysv2Accepted() *SetBucketKeysv2Accepted {
	return &SetBucketKeysv2Accepted{}
}

/*SetBucketKeysv2Accepted handles this case with default header values.

successful operation
*/
type SetBucketKeysv2Accepted struct {
	Payload *models.AsyncRequestResponse
}

func (o *SetBucketKeysv2Accepted) Error() string {
	return fmt.Sprintf("[PUT /scalemgmt/v2/bucket/keys][%d] setBucketKeysv2Accepted  %+v", 202, o.Payload)
}

func (o *SetBucketKeysv2Accepted) GetPayload() *models.AsyncRequestResponse {
	return o.Payload
}

func (o *SetBucketKeysv2Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetBucketKeysv2BadRequest creates a SetBucketKeysv2BadRequest with default headers values
func NewSetBucketKeysv2BadRequest() *SetBucketKeysv2BadRequest {
	return &SetBucketKeysv2BadRequest{}
}

/*SetBucketKeysv2BadRequest handles this case with default header values.

Invalid request
*/
type SetBucketKeysv2BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *SetBucketKeysv2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /scalemgmt/v2/bucket/keys][%d] setBucketKeysv2BadRequest  %+v", 400, o.Payload)
}

func (o *SetBucketKeysv2BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *SetBucketKeysv2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetBucketKeysv2InternalServerError creates a SetBucketKeysv2InternalServerError with default headers values
func NewSetBucketKeysv2InternalServerError() *SetBucketKeysv2InternalServerError {
	return &SetBucketKeysv2InternalServerError{}
}

/*SetBucketKeysv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type SetBucketKeysv2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *SetBucketKeysv2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /scalemgmt/v2/bucket/keys][%d] setBucketKeysv2InternalServerError  %+v", 500, o.Payload)
}

func (o *SetBucketKeysv2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *SetBucketKeysv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}