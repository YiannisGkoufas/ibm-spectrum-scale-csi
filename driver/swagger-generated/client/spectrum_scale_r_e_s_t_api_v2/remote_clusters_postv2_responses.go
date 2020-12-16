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

// RemoteClustersPostv2Reader is a Reader for the RemoteClustersPostv2 structure.
type RemoteClustersPostv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoteClustersPostv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRemoteClustersPostv2Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoteClustersPostv2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoteClustersPostv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoteClustersPostv2Accepted creates a RemoteClustersPostv2Accepted with default headers values
func NewRemoteClustersPostv2Accepted() *RemoteClustersPostv2Accepted {
	return &RemoteClustersPostv2Accepted{}
}

/*RemoteClustersPostv2Accepted handles this case with default header values.

successful operation
*/
type RemoteClustersPostv2Accepted struct {
	Payload *models.AsyncRequestResponse
}

func (o *RemoteClustersPostv2Accepted) Error() string {
	return fmt.Sprintf("[POST /scalemgmt/v2/remotemount/remoteclusters][%d] remoteClustersPostv2Accepted  %+v", 202, o.Payload)
}

func (o *RemoteClustersPostv2Accepted) GetPayload() *models.AsyncRequestResponse {
	return o.Payload
}

func (o *RemoteClustersPostv2Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoteClustersPostv2BadRequest creates a RemoteClustersPostv2BadRequest with default headers values
func NewRemoteClustersPostv2BadRequest() *RemoteClustersPostv2BadRequest {
	return &RemoteClustersPostv2BadRequest{}
}

/*RemoteClustersPostv2BadRequest handles this case with default header values.

Invalid request
*/
type RemoteClustersPostv2BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *RemoteClustersPostv2BadRequest) Error() string {
	return fmt.Sprintf("[POST /scalemgmt/v2/remotemount/remoteclusters][%d] remoteClustersPostv2BadRequest  %+v", 400, o.Payload)
}

func (o *RemoteClustersPostv2BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *RemoteClustersPostv2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoteClustersPostv2InternalServerError creates a RemoteClustersPostv2InternalServerError with default headers values
func NewRemoteClustersPostv2InternalServerError() *RemoteClustersPostv2InternalServerError {
	return &RemoteClustersPostv2InternalServerError{}
}

/*RemoteClustersPostv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoteClustersPostv2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *RemoteClustersPostv2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /scalemgmt/v2/remotemount/remoteclusters][%d] remoteClustersPostv2InternalServerError  %+v", 500, o.Payload)
}

func (o *RemoteClustersPostv2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *RemoteClustersPostv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}