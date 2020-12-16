// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileAuditOperation file audit operation
//
// swagger:model FileAuditOperation
type FileAuditOperation struct {

	// The action that shall be applied [enable, disable, update]
	Action string `json:"action,omitempty"`

	// file audit log config
	FileAuditLogConfig *FileAuditLogConfigEnable `json:"fileAuditLogConfig,omitempty"`
}

// Validate validates this file audit operation
func (m *FileAuditOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileAuditLogConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileAuditOperation) validateFileAuditLogConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FileAuditLogConfig) { // not required
		return nil
	}

	if m.FileAuditLogConfig != nil {
		if err := m.FileAuditLogConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileAuditLogConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileAuditOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileAuditOperation) UnmarshalBinary(b []byte) error {
	var res FileAuditOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}