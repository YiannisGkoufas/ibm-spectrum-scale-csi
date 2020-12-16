// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Watchfolder Detailed information about a clustered filesystem watch
//
// swagger:model Watchfolder
type Watchfolder struct {

	// The degraded state of the clustered watch
	Degraded bool `json:"degraded,omitempty"`

	// The description of the clustered watch
	Description string `json:"description,omitempty"`

	// The device on which the clustered watch is enabled
	Device string `json:"device,omitempty"`

	// The event handler (only kafkasink is supported)
	EventHandler string `json:"eventHandler,omitempty"`

	// The list of watched events
	Events []string `json:"events"`

	// The fileset on which the clustered watch is enabled
	Fileset string `json:"fileset,omitempty"`

	// The id of the filesystem watch
	ID string `json:"id,omitempty"`

	// node status
	NodeStatus []*WatchfolderNodeStatus `json:"nodeStatus"`

	// oid
	Oid int64 `json:"oid,omitempty"`

	// The path of the clustered watch
	Path string `json:"path,omitempty"`

	// The list of sink brokers
	SinkBrokers []string `json:"sinkBrokers"`

	// The sink topic
	SinkTopic string `json:"sinkTopic,omitempty"`

	// The state of the clustered watch
	State string `json:"state,omitempty"`

	// The type of the clustered watch
	// Enum: [FILESYSTEM INODESPACE INODE FILESET INCOMPLETE FSYS FSET]
	Type string `json:"type,omitempty"`
}

// Validate validates this watchfolder
func (m *Watchfolder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Watchfolder) validateNodeStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeStatus) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeStatus); i++ {
		if swag.IsZero(m.NodeStatus[i]) { // not required
			continue
		}

		if m.NodeStatus[i] != nil {
			if err := m.NodeStatus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeStatus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var watchfolderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FILESYSTEM","INODESPACE","INODE","FILESET","INCOMPLETE","FSYS","FSET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		watchfolderTypeTypePropEnum = append(watchfolderTypeTypePropEnum, v)
	}
}

const (

	// WatchfolderTypeFILESYSTEM captures enum value "FILESYSTEM"
	WatchfolderTypeFILESYSTEM string = "FILESYSTEM"

	// WatchfolderTypeINODESPACE captures enum value "INODESPACE"
	WatchfolderTypeINODESPACE string = "INODESPACE"

	// WatchfolderTypeINODE captures enum value "INODE"
	WatchfolderTypeINODE string = "INODE"

	// WatchfolderTypeFILESET captures enum value "FILESET"
	WatchfolderTypeFILESET string = "FILESET"

	// WatchfolderTypeINCOMPLETE captures enum value "INCOMPLETE"
	WatchfolderTypeINCOMPLETE string = "INCOMPLETE"

	// WatchfolderTypeFSYS captures enum value "FSYS"
	WatchfolderTypeFSYS string = "FSYS"

	// WatchfolderTypeFSET captures enum value "FSET"
	WatchfolderTypeFSET string = "FSET"
)

// prop value enum
func (m *Watchfolder) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, watchfolderTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Watchfolder) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Watchfolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Watchfolder) UnmarshalBinary(b []byte) error {
	var res Watchfolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}