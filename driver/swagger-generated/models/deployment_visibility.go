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

// DeploymentVisibility deployment visibility
//
// swagger:model DeploymentVisibility
type DeploymentVisibility struct {

	// subscribe
	// Required: true
	Subscribe *Subscribe `json:"subscribe"`

	// view
	// Required: true
	View *View `json:"view"`
}

// Validate validates this deployment visibility
func (m *DeploymentVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscribe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateView(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentVisibility) validateSubscribe(formats strfmt.Registry) error {

	if err := validate.Required("subscribe", "body", m.Subscribe); err != nil {
		return err
	}

	if m.Subscribe != nil {
		if err := m.Subscribe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subscribe")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentVisibility) validateView(formats strfmt.Registry) error {

	if err := validate.Required("view", "body", m.View); err != nil {
		return err
	}

	if m.View != nil {
		if err := m.View.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("view")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentVisibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentVisibility) UnmarshalBinary(b []byte) error {
	var res DeploymentVisibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}