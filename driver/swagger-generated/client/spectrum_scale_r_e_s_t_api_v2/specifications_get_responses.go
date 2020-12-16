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

// SpecificationsGetReader is a Reader for the SpecificationsGet structure.
type SpecificationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecificationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpecificationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSpecificationsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSpecificationsGetOK creates a SpecificationsGetOK with default headers values
func NewSpecificationsGetOK() *SpecificationsGetOK {
	return &SpecificationsGetOK{}
}

/*SpecificationsGetOK handles this case with default header values.

successful operation
*/
type SpecificationsGetOK struct {
	Payload *models.ComponentSpecificationInlineResponse200
}

func (o *SpecificationsGetOK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/components/specifications][%d] specificationsGetOK  %+v", 200, o.Payload)
}

func (o *SpecificationsGetOK) GetPayload() *models.ComponentSpecificationInlineResponse200 {
	return o.Payload
}

func (o *SpecificationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentSpecificationInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecificationsGetInternalServerError creates a SpecificationsGetInternalServerError with default headers values
func NewSpecificationsGetInternalServerError() *SpecificationsGetInternalServerError {
	return &SpecificationsGetInternalServerError{}
}

/*SpecificationsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type SpecificationsGetInternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *SpecificationsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/components/specifications][%d] specificationsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *SpecificationsGetInternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *SpecificationsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}