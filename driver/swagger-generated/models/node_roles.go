// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeRoles Information about the nodes roles
//
// swagger:model NodeRoles
type NodeRoles struct {

	// Specifies if this node is a CES node
	CesNode bool `json:"cesNode,omitempty"`

	// Specifies if this node is a client node
	ClientNode bool `json:"clientNode,omitempty"`

	// Specifies if this node is a cloud gateway node
	CloudGatewayNode bool `json:"cloudGatewayNode,omitempty"`

	// Specifies if this node is a cNFS node
	CnfsNode bool `json:"cnfsNode,omitempty"`

	// The nodes designation
	// Enum: [quorum quorumManager manager]
	Designation string `json:"designation,omitempty"`

	// Specifies if this node is a gateway node
	GatewayNode bool `json:"gatewayNode,omitempty"`

	// Specifies if this node is a manager node
	ManagerNode bool `json:"managerNode,omitempty"`

	// Other roles of this node
	OtherNodeRoles string `json:"otherNodeRoles,omitempty"`

	// Specifies if this node is a quorum node
	QuorumNode bool `json:"quorumNode,omitempty"`

	// Specifies if this node is a SNMP node
	SnmpNode bool `json:"snmpNode,omitempty"`
}

// Validate validates this node roles
func (m *NodeRoles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesignation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nodeRolesTypeDesignationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["quorum","quorumManager","manager"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeRolesTypeDesignationPropEnum = append(nodeRolesTypeDesignationPropEnum, v)
	}
}

const (

	// NodeRolesDesignationQuorum captures enum value "quorum"
	NodeRolesDesignationQuorum string = "quorum"

	// NodeRolesDesignationQuorumManager captures enum value "quorumManager"
	NodeRolesDesignationQuorumManager string = "quorumManager"

	// NodeRolesDesignationManager captures enum value "manager"
	NodeRolesDesignationManager string = "manager"
)

// prop value enum
func (m *NodeRoles) validateDesignationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeRolesTypeDesignationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeRoles) validateDesignation(formats strfmt.Registry) error {

	if swag.IsZero(m.Designation) { // not required
		return nil
	}

	// value enum
	if err := m.validateDesignationEnum("designation", "body", m.Designation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeRoles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeRoles) UnmarshalBinary(b []byte) error {
	var res NodeRoles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}