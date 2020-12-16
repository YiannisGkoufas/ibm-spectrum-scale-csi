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

// FilesystemWatchInlineResponse200 filesystem watch inline response200
//
// swagger:model FilesystemWatchInlineResponse200
type FilesystemWatchInlineResponse200 struct {

	// paging
	Paging *Paging `json:"paging,omitempty"`

	// The status of the request
	Status *RequestStatus `json:"status,omitempty"`

	// watches
	Watches []*Watchfolder `json:"watches"`
}

// Validate validates this filesystem watch inline response200
func (m *FilesystemWatchInlineResponse200) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWatches(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesystemWatchInlineResponse200) validatePaging(formats strfmt.Registry) error {

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

func (m *FilesystemWatchInlineResponse200) validateStatus(formats strfmt.Registry) error {

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

func (m *FilesystemWatchInlineResponse200) validateWatches(formats strfmt.Registry) error {

	if swag.IsZero(m.Watches) { // not required
		return nil
	}

	for i := 0; i < len(m.Watches); i++ {
		if swag.IsZero(m.Watches[i]) { // not required
			continue
		}

		if m.Watches[i] != nil {
			if err := m.Watches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("watches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilesystemWatchInlineResponse200) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesystemWatchInlineResponse200) UnmarshalBinary(b []byte) error {
	var res FilesystemWatchInlineResponse200
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}