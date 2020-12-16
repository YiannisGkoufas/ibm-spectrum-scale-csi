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

// SmbShare Summary information about a smb share
//
// swagger:model SmbShare
type SmbShare struct {

	// The name of the fileset the share belongs to.
	FilesetName string `json:"filesetName,omitempty"`

	// The name of the filesystem the share belongs to.
	// Required: true
	FilesystemName *string `json:"filesystemName"`

	// Specifies the path for the share.
	Path string `json:"path,omitempty"`

	// The name of the SMB share.
	// Required: true
	ShareName *string `json:"shareName"`

	// smb options
	SmbOptions *SmbOptions `json:"smbOptions,omitempty"`
}

// Validate validates this smb share
func (m *SmbShare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilesystemName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmbOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbShare) validateFilesystemName(formats strfmt.Registry) error {

	if err := validate.Required("filesystemName", "body", m.FilesystemName); err != nil {
		return err
	}

	return nil
}

func (m *SmbShare) validateShareName(formats strfmt.Registry) error {

	if err := validate.Required("shareName", "body", m.ShareName); err != nil {
		return err
	}

	return nil
}

func (m *SmbShare) validateSmbOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.SmbOptions) { // not required
		return nil
	}

	if m.SmbOptions != nil {
		if err := m.SmbOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smbOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbShare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbShare) UnmarshalBinary(b []byte) error {
	var res SmbShare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}