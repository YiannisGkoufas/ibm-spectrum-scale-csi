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

// RemoteAccessStatusGetReader is a Reader for the RemoteAccessStatusGet structure.
type RemoteAccessStatusGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoteAccessStatusGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoteAccessStatusGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRemoteAccessStatusGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoteAccessStatusGetOK creates a RemoteAccessStatusGetOK with default headers values
func NewRemoteAccessStatusGetOK() *RemoteAccessStatusGetOK {
	return &RemoteAccessStatusGetOK{}
}

/*RemoteAccessStatusGetOK handles this case with default header values.

successful operation
*/
type RemoteAccessStatusGetOK struct {
	Payload *models.RemoteAccessStatusInlineResponse200
}

func (o *RemoteAccessStatusGetOK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/access/status][%d] remoteAccessStatusGetOK  %+v", 200, o.Payload)
}

func (o *RemoteAccessStatusGetOK) GetPayload() *models.RemoteAccessStatusInlineResponse200 {
	return o.Payload
}

func (o *RemoteAccessStatusGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteAccessStatusInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoteAccessStatusGetInternalServerError creates a RemoteAccessStatusGetInternalServerError with default headers values
func NewRemoteAccessStatusGetInternalServerError() *RemoteAccessStatusGetInternalServerError {
	return &RemoteAccessStatusGetInternalServerError{}
}

/*RemoteAccessStatusGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoteAccessStatusGetInternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *RemoteAccessStatusGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/access/status][%d] remoteAccessStatusGetInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoteAccessStatusGetInternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *RemoteAccessStatusGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}