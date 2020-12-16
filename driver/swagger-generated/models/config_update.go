// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigUpdate Summary information about an update of GPFS config
//
// swagger:model ConfigUpdate
type ConfigUpdate struct {

	// List of attributes and their values to be changed
	Attributes map[string]string `json:"attributes,omitempty"`

	// Changes take effect immediately, they do persist when GPFS is restarted. Valid only for specific attributtes
	Immediately bool `json:"immediately,omitempty"`

	// List of nodes or nodeclasses. If not specified, the default the nodeclass 'all' is used
	Nodes []string `json:"nodes"`
}

// Validate validates this config update
func (m *ConfigUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigUpdate) UnmarshalBinary(b []byte) error {
	var res ConfigUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}