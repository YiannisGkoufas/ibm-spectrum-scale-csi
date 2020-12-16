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

// NfsExportsGetReader is a Reader for the NfsExportsGet structure.
type NfsExportsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsExportsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsExportsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewNfsExportsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNfsExportsGetOK creates a NfsExportsGetOK with default headers values
func NewNfsExportsGetOK() *NfsExportsGetOK {
	return &NfsExportsGetOK{}
}

/*NfsExportsGetOK handles this case with default header values.

successful operation
*/
type NfsExportsGetOK struct {
	Payload *models.NfsExportInlineResponse200
}

func (o *NfsExportsGetOK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nfs/exports][%d] nfsExportsGetOK  %+v", 200, o.Payload)
}

func (o *NfsExportsGetOK) GetPayload() *models.NfsExportInlineResponse200 {
	return o.Payload
}

func (o *NfsExportsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsExportInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNfsExportsGetInternalServerError creates a NfsExportsGetInternalServerError with default headers values
func NewNfsExportsGetInternalServerError() *NfsExportsGetInternalServerError {
	return &NfsExportsGetInternalServerError{}
}

/*NfsExportsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type NfsExportsGetInternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *NfsExportsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nfs/exports][%d] nfsExportsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *NfsExportsGetInternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *NfsExportsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}