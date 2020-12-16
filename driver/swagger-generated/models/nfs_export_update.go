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

// NfsExportUpdate Summary information about a nfs export
//
// swagger:model NfsExportUpdate
type NfsExportUpdate struct {

	// Adds a new client declaration for the specified Path. ClientOptions can be a list of one or more client definitions.
	// Required: true
	Nfsadd *string `json:"nfsadd"`

	// Modifies an existing client declaration for the specified Path. ClientOptions can be a list of one or more client definitions.
	// Required: true
	Nfschange *string `json:"nfschange"`

	// It reorders the client declaration section within the export declaration file. This option can be used only together with --nfsadd or --nfschange.
	// Required: true
	Nfsposition *string `json:"nfsposition"`

	// Removes the NFS client specified by client from the export configuration for the path.
	// Required: true
	Nfsremove *string `json:"nfsremove"`
}

// Validate validates this nfs export update
func (m *NfsExportUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNfsadd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfschange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsremove(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsExportUpdate) validateNfsadd(formats strfmt.Registry) error {

	if err := validate.Required("nfsadd", "body", m.Nfsadd); err != nil {
		return err
	}

	return nil
}

func (m *NfsExportUpdate) validateNfschange(formats strfmt.Registry) error {

	if err := validate.Required("nfschange", "body", m.Nfschange); err != nil {
		return err
	}

	return nil
}

func (m *NfsExportUpdate) validateNfsposition(formats strfmt.Registry) error {

	if err := validate.Required("nfsposition", "body", m.Nfsposition); err != nil {
		return err
	}

	return nil
}

func (m *NfsExportUpdate) validateNfsremove(formats strfmt.Registry) error {

	if err := validate.Required("nfsremove", "body", m.Nfsremove); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsExportUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsExportUpdate) UnmarshalBinary(b []byte) error {
	var res NfsExportUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}