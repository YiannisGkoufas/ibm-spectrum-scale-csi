// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductPackageUpdateJSONBody product package update Json body
//
// swagger:model ProductPackageUpdateJsonBody
type ProductPackageUpdateJSONBody struct {

	// visibility
	Visibility *DeploymentVisibility `json:"visibility,omitempty"`
}

// Validate validates this product package update Json body
func (m *ProductPackageUpdateJSONBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPackageUpdateJSONBody) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	if m.Visibility != nil {
		if err := m.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductPackageUpdateJSONBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductPackageUpdateJSONBody) UnmarshalBinary(b []byte) error {
	var res ProductPackageUpdateJSONBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}