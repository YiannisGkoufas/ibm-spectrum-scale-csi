// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Http503ServiceUnavailable http503 service unavailable
//
// swagger:model Http503ServiceUnavailable
type Http503ServiceUnavailable struct {

	// The HTTP status code that was returned by the request
	Code int32 `json:"code,omitempty"`

	// The detailed success/error message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this http503 service unavailable
func (m *Http503ServiceUnavailable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Http503ServiceUnavailable) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Http503ServiceUnavailable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Http503ServiceUnavailable) UnmarshalBinary(b []byte) error {
	var res Http503ServiceUnavailable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}