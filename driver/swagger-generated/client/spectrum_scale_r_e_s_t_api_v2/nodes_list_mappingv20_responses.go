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

// NodesListMappingv20Reader is a Reader for the NodesListMappingv20 structure.
type NodesListMappingv20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesListMappingv20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewNodesListMappingv20Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNodesListMappingv20BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodesListMappingv20InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodesListMappingv20Accepted creates a NodesListMappingv20Accepted with default headers values
func NewNodesListMappingv20Accepted() *NodesListMappingv20Accepted {
	return &NodesListMappingv20Accepted{}
}

/*NodesListMappingv20Accepted handles this case with default header values.

successful operation
*/
type NodesListMappingv20Accepted struct {
	Payload *models.NodeMappingInlineResponse200
}

func (o *NodesListMappingv20Accepted) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nodes/afm/mapping/{mappingName}][%d] nodesListMappingv20Accepted  %+v", 202, o.Payload)
}

func (o *NodesListMappingv20Accepted) GetPayload() *models.NodeMappingInlineResponse200 {
	return o.Payload
}

func (o *NodesListMappingv20Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeMappingInlineResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesListMappingv20BadRequest creates a NodesListMappingv20BadRequest with default headers values
func NewNodesListMappingv20BadRequest() *NodesListMappingv20BadRequest {
	return &NodesListMappingv20BadRequest{}
}

/*NodesListMappingv20BadRequest handles this case with default header values.

Invalid request
*/
type NodesListMappingv20BadRequest struct {
	Payload *models.Http400BadRequest
}

func (o *NodesListMappingv20BadRequest) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nodes/afm/mapping/{mappingName}][%d] nodesListMappingv20BadRequest  %+v", 400, o.Payload)
}

func (o *NodesListMappingv20BadRequest) GetPayload() *models.Http400BadRequest {
	return o.Payload
}

func (o *NodesListMappingv20BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http400BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesListMappingv20InternalServerError creates a NodesListMappingv20InternalServerError with default headers values
func NewNodesListMappingv20InternalServerError() *NodesListMappingv20InternalServerError {
	return &NodesListMappingv20InternalServerError{}
}

/*NodesListMappingv20InternalServerError handles this case with default header values.

Internal Server Error
*/
type NodesListMappingv20InternalServerError struct {
	Payload *models.Http500InternalServerError
}

func (o *NodesListMappingv20InternalServerError) Error() string {
	return fmt.Sprintf("[GET /scalemgmt/v2/nodes/afm/mapping/{mappingName}][%d] nodesListMappingv20InternalServerError  %+v", 500, o.Payload)
}

func (o *NodesListMappingv20InternalServerError) GetPayload() *models.Http500InternalServerError {
	return o.Payload
}

func (o *NodesListMappingv20InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Http500InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}