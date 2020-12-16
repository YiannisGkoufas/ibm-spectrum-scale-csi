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

// NsdsGetv2Reader is a Reader for the NsdsGetv2 structure.
type NsdsGetv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NsdsGetv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNsdsGetv2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewNsdsGetv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNsdsGetv2OK creates a NsdsGetv2OK with default headers values
func NewNsdsGetv2OK() *NsdsGetv2OK {
	return &NsdsGetv2OK{}
}

/*NsdsGetv2OK handles this case with default header values.

successful operation
*/
type NsdsGetv2OK struct {
	Payload *models.NsdInlineResponse200
}

func (o *NsdsGetv2OK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nsds][%d] nsdsGetv2OK  %+v", 200, o.Payload)
}

func (o *NsdsGetv2OK) GetPayload() *models.NsdInlineResponse200 {
	return o.Payload
}

func (o *NsdsGetv2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsdInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNsdsGetv2InternalServerError creates a NsdsGetv2InternalServerError with default headers values
func NewNsdsGetv2InternalServerError() *NsdsGetv2InternalServerError {
	return &NsdsGetv2InternalServerError{}
}

/*NsdsGetv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type NsdsGetv2InternalServerError struct {
	Payload *models.NsdInlineResponse200
}

func (o *NsdsGetv2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nsds][%d] nsdsGetv2InternalServerError  %+v", 500, o.Payload)
}

func (o *NsdsGetv2InternalServerError) GetPayload() *models.NsdInlineResponse200 {
	return o.Payload
}

func (o *NsdsGetv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NsdInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}