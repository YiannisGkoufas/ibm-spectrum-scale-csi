// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaDefaultsFsConfig quota defaults fs config
//
// swagger:model QuotaDefaultsFsConfig
type QuotaDefaultsFsConfig struct {

	// Assigns the default quota limits to/of existing users, groups, or filesets when quota defaults are enabled.
	Assign bool `json:"assign,omitempty"`

	// Specifies that default quotas for filesets are to be activated.
	Fileset bool `json:"fileset,omitempty"`

	// Specifies that default quotas for groups are to be activated.
	Group bool `json:"group,omitempty"`

	// Reset the default quota limits to/of existing users, groups, or filesets when quota defaults are disabled.
	Reset bool `json:"reset,omitempty"`

	// Specifies that default quotas for users are to be activated.
	User bool `json:"user,omitempty"`
}

// Validate validates this quota defaults fs config
func (m *QuotaDefaultsFsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaDefaultsFsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaDefaultsFsConfig) UnmarshalBinary(b []byte) error {
	var res QuotaDefaultsFsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}