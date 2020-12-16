// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaDefaultsCreate quota defaults create
//
// swagger:model QuotaDefaultsCreate
type QuotaDefaultsCreate struct {

	// The quota hard limit for data. The number can be specified using the suffix K, M, G, or T
	BlockHardLimit string `json:"blockHardLimit,omitempty"`

	// The quota soft limit for data. The number can be specified using the suffix K, M, G, or T
	BlockSoftLimit string `json:"blockSoftLimit,omitempty"`

	// The quota hard limit for files. The number can be specified using the suffix K, M, G, or T
	FilesHardLimit string `json:"filesHardLimit,omitempty"`

	// The quota soft limit for files. The number can be specified using the suffix K, M, G, or T
	FilesSoftLimit string `json:"filesSoftLimit,omitempty"`

	// The type of the Quota defaults to set. Either USR, GRP or FILESET
	QuotaType string `json:"quotaType,omitempty"`
}

// Validate validates this quota defaults create
func (m *QuotaDefaultsCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaDefaultsCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaDefaultsCreate) UnmarshalBinary(b []byte) error {
	var res QuotaDefaultsCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}