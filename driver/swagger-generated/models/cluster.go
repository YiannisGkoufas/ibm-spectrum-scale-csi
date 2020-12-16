// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cluster Summary information about a cluster
//
// swagger:model Cluster
type Cluster struct {

	// capacity licensing
	CapacityLicensing *CapacityLiable4Licensing `json:"capacityLicensing,omitempty"`

	// ces summary
	CesSummary *CESSummary `json:"cesSummary,omitempty"`

	// cluster summary
	ClusterSummary *ClusterSummary `json:"clusterSummary,omitempty"`

	// cnfs summary
	CnfsSummary *CnfsSummary `json:"cnfsSummary,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityLicensing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCesSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCnfsSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateCapacityLicensing(formats strfmt.Registry) error {

	if swag.IsZero(m.CapacityLicensing) { // not required
		return nil
	}

	if m.CapacityLicensing != nil {
		if err := m.CapacityLicensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacityLicensing")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateCesSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.CesSummary) { // not required
		return nil
	}

	if m.CesSummary != nil {
		if err := m.CesSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cesSummary")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateClusterSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterSummary) { // not required
		return nil
	}

	if m.ClusterSummary != nil {
		if err := m.ClusterSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSummary")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validateCnfsSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.CnfsSummary) { // not required
		return nil
	}

	if m.CnfsSummary != nil {
		if err := m.CnfsSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cnfsSummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}