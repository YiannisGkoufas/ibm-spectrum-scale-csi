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

// OwningClustersGetv2Reader is a Reader for the OwningClustersGetv2 structure.
type OwningClustersGetv2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OwningClustersGetv2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOwningClustersGetv2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewOwningClustersGetv2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOwningClustersGetv2OK creates a OwningClustersGetv2OK with default headers values
func NewOwningClustersGetv2OK() *OwningClustersGetv2OK {
	return &OwningClustersGetv2OK{}
}

/*OwningClustersGetv2OK handles this case with default header values.

successful operation
*/
type OwningClustersGetv2OK struct {
	Payload *models.OwningClusterInlineResponse200
}

func (o *OwningClustersGetv2OK) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/remotemount/owningclusters][%d] owningClustersGetv2OK  %+v", 200, o.Payload)
}

func (o *OwningClustersGetv2OK) GetPayload() *models.OwningClusterInlineResponse200 {
	return o.Payload
}

func (o *OwningClustersGetv2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OwningClusterInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOwningClustersGetv2InternalServerError creates a OwningClustersGetv2InternalServerError with default headers values
func NewOwningClustersGetv2InternalServerError() *OwningClustersGetv2InternalServerError {
	return &OwningClustersGetv2InternalServerError{}
}

/*OwningClustersGetv2InternalServerError handles this case with default header values.

Internal Server Error
*/
type OwningClustersGetv2InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *OwningClustersGetv2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/remotemount/owningclusters][%d] owningClustersGetv2InternalServerError  %+v", 500, o.Payload)
}

func (o *OwningClustersGetv2InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *OwningClustersGetv2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}