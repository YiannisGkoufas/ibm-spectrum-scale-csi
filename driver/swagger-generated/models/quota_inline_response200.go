// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaInlineResponse200 quota inline response200
//
// swagger:model QuotaInlineResponse200
type QuotaInlineResponse200 struct {

	// paging
	Paging *Paging `json:"paging,omitempty"`

	// quotas
	Quotas []*Quota `json:"quotas"`

	// The status of the request
	Status *RequestStatus `json:"status,omitempty"`
}

// Validate validates this quota inline response200
func (m *QuotaInlineResponse200) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaInlineResponse200) validatePaging(formats strfmt.Registry) error {

	if swag.IsZero(m.Paging) { // not required
		return nil
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaInlineResponse200) validateQuotas(formats strfmt.Registry) error {

	if swag.IsZero(m.Quotas) { // not required
		return nil
	}

	for i := 0; i < len(m.Quotas); i++ {
		if swag.IsZero(m.Quotas[i]) { // not required
			continue
		}

		if m.Quotas[i] != nil {
			if err := m.Quotas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuotaInlineResponse200) validateStatus(formats strfmt.Registry) error {

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
func (m *QuotaInlineResponse200) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaInlineResponse200) UnmarshalBinary(b []byte) error {
	var res QuotaInlineResponse200
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}