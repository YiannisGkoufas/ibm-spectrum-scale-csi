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

// SnapshotCreate Summary information about a Snapshot
//
// swagger:model SnapshotCreate
type SnapshotCreate struct {

	// The name of the snapshot to create
	// Required: true
	SnapshotName *string `json:"snapshotName"`
}

// Validate validates this snapshot create
func (m *SnapshotCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotCreate) validateSnapshotName(formats strfmt.Registry) error {

	if err := validate.Required("snapshotName", "body", m.SnapshotName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotCreate) UnmarshalBinary(b []byte) error {
	var res SnapshotCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}