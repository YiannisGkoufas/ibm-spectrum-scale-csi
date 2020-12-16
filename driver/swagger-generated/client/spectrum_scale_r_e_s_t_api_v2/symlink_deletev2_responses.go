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

// SymlinkDeletev2Reader is a Reader for the SymlinkDeletev2 structure.
type SymlinkDeletev2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SymlinkDeletev2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSymlinkDeletev2Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSymlinkDeletev2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSymlinkDeletev2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSymlinkDeletev2Accepted creates a SymlinkDeletev2Accepted with default headers values
func NewSymlinkDeletev2Accepted() *SymlinkDeletev2Accepted {
	return &SymlinkDeletev2Accepted{}
}

/*SymlinkDeletev2Accepted handles this case with default header values.

successful operation
*/
type SymlinkDeletev2Accepted struct {
	Payload *models.AsyncRequestResponse
}

func (o *SymlinkDeletev2Accepted) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/symlink/{path}][%d] symlinkDeletev2Accepted  %+v", 202, o.Payload)
}

func (o *SymlinkDeletev2Accepted) GetPayload() *models.AsyncRequestResponse {
	return o.Payload
}

func (o *SymlinkDeletev2Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSymlinkDeletev2BadRequest creates a SymlinkDeletev2BadRequest with default headers values
func NewSymlinkDeletev2BadRequest() *SymlinkDeletev2BadRequest {
	return &SymlinkDeletev2BadRequest{}
}

/*SymlinkDeletev2BadRequest handles this case with default header values.

Invalid file system or path
*/
type SymlinkDeletev2BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *SymlinkDeletev2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/symlink/{path}][%d] symlinkDeletev2BadRequest  %+v", 400, o.Payload)
}

func (o *SymlinkDeletev2BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *SymlinkDeletev2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSymlinkDeletev2InternalServerError creates a SymlinkDeletev2InternalServerError with default headers values
func NewSymlinkDeletev2InternalServerError() *SymlinkDeletev2InternalServerError {
	return &SymlinkDeletev2InternalServerError{}
}

/*SymlinkDeletev2InternalServerError handles this case with default header values.

An unexpected error occurred.
*/
type SymlinkDeletev2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *SymlinkDeletev2InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/symlink/{path}][%d] symlinkDeletev2InternalServerError  %+v", 500, o.Payload)
}

func (o *SymlinkDeletev2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *SymlinkDeletev2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}