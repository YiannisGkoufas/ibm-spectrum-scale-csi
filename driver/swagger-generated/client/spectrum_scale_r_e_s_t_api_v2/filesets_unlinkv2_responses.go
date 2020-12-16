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

// FilesetsUnlinkv2Reader is a Reader for the FilesetsUnlinkv2 structure.
type FilesetsUnlinkv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilesetsUnlinkv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewFilesetsUnlinkv2Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFilesetsUnlinkv2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFilesetsUnlinkv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFilesetsUnlinkv2Accepted creates a FilesetsUnlinkv2Accepted with default headers values
func NewFilesetsUnlinkv2Accepted() *FilesetsUnlinkv2Accepted {
	return &FilesetsUnlinkv2Accepted{}
}

/*FilesetsUnlinkv2Accepted handles this case with default header values.

successful operation
*/
type FilesetsUnlinkv2Accepted struct {
	Payload *models.AsyncRequestResponse
}

func (o *FilesetsUnlinkv2Accepted) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/link][%d] filesetsUnlinkv2Accepted  %+v", 202, o.Payload)
}

func (o *FilesetsUnlinkv2Accepted) GetPayload() *models.AsyncRequestResponse {
	return o.Payload
}

func (o *FilesetsUnlinkv2Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesetsUnlinkv2BadRequest creates a FilesetsUnlinkv2BadRequest with default headers values
func NewFilesetsUnlinkv2BadRequest() *FilesetsUnlinkv2BadRequest {
	return &FilesetsUnlinkv2BadRequest{}
}

/*FilesetsUnlinkv2BadRequest handles this case with default header values.

Invalid request
*/
type FilesetsUnlinkv2BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *FilesetsUnlinkv2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/link][%d] filesetsUnlinkv2BadRequest  %+v", 400, o.Payload)
}

func (o *FilesetsUnlinkv2BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *FilesetsUnlinkv2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesetsUnlinkv2InternalServerError creates a FilesetsUnlinkv2InternalServerError with default headers values
func NewFilesetsUnlinkv2InternalServerError() *FilesetsUnlinkv2InternalServerError {
	return &FilesetsUnlinkv2InternalServerError{}
}

/*FilesetsUnlinkv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type FilesetsUnlinkv2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *FilesetsUnlinkv2InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/link][%d] filesetsUnlinkv2InternalServerError  %+v", 500, o.Payload)
}

func (o *FilesetsUnlinkv2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *FilesetsUnlinkv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}