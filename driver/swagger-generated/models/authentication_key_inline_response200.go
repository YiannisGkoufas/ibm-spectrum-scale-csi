// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthenticationKeyInlineResponse200 authentication key inline response200
//
// swagger:model AuthenticationKeyInlineResponse200
type AuthenticationKeyInlineResponse200 struct {

	// The list of ciphers of the RSA key, or one of DEFAULT, EMPTY or AUTHONLY
	Ciphers []string `json:"ciphers"`

	// The currently active public RSA key
	Key []string `json:"key"`

	// A new generated public RSA key. This key will become active once it's comitted. To commit this new key use PUT /scalemgmt/v2/remotemount/authenticationkey
	NewKey []string `json:"newKey"`

	// The status of the request
	Status *RequestStatus `json:"status,omitempty"`
}

// Validate validates this authentication key inline response200
func (m *AuthenticationKeyInlineResponse200) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationKeyInlineResponse200) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationKeyInlineResponse200) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationKeyInlineResponse200) UnmarshalBinary(b []byte) error {
	var res AuthenticationKeyInlineResponse200
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}