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

// SnapshotsFilesetsGetv2Reader is a Reader for the SnapshotsFilesetsGetv2 structure.
type SnapshotsFilesetsGetv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotsFilesetsGetv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotsFilesetsGetv2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSnapshotsFilesetsGetv2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSnapshotsFilesetsGetv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSnapshotsFilesetsGetv2OK creates a SnapshotsFilesetsGetv2OK with default headers values
func NewSnapshotsFilesetsGetv2OK() *SnapshotsFilesetsGetv2OK {
	return &SnapshotsFilesetsGetv2OK{}
}

/*SnapshotsFilesetsGetv2OK handles this case with default header values.

successful operation
*/
type SnapshotsFilesetsGetv2OK struct {
	Payload *models.SnapshotInlineResponse200
}

func (o *SnapshotsFilesetsGetv2OK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/snapshots][%d] snapshotsFilesetsGetv2OK  %+v", 200, o.Payload)
}

func (o *SnapshotsFilesetsGetv2OK) GetPayload() *models.SnapshotInlineResponse200 {
	return o.Payload
}

func (o *SnapshotsFilesetsGetv2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotsFilesetsGetv2BadRequest creates a SnapshotsFilesetsGetv2BadRequest with default headers values
func NewSnapshotsFilesetsGetv2BadRequest() *SnapshotsFilesetsGetv2BadRequest {
	return &SnapshotsFilesetsGetv2BadRequest{}
}

/*SnapshotsFilesetsGetv2BadRequest handles this case with default header values.

Invalid request
*/
type SnapshotsFilesetsGetv2BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *SnapshotsFilesetsGetv2BadRequest) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/snapshots][%d] snapshotsFilesetsGetv2BadRequest  %+v", 400, o.Payload)
}

func (o *SnapshotsFilesetsGetv2BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *SnapshotsFilesetsGetv2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotsFilesetsGetv2InternalServerError creates a SnapshotsFilesetsGetv2InternalServerError with default headers values
func NewSnapshotsFilesetsGetv2InternalServerError() *SnapshotsFilesetsGetv2InternalServerError {
	return &SnapshotsFilesetsGetv2InternalServerError{}
}

/*SnapshotsFilesetsGetv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type SnapshotsFilesetsGetv2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *SnapshotsFilesetsGetv2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/filesystems/{filesystemName}/filesets/{filesetName}/snapshots][%d] snapshotsFilesetsGetv2InternalServerError  %+v", 500, o.Payload)
}

func (o *SnapshotsFilesetsGetv2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *SnapshotsFilesetsGetv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}