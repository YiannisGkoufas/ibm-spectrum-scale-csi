// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobResult job result
//
// swagger:model JobResult
type JobResult struct {

	// Executed commands in this job
	Commands []string `json:"commands"`

	// Exit code of command. Zero is success, non-zero failure
	ExitCode int64 `json:"exitCode,omitempty"`

	// Progress information for request.
	Progress []string `json:"progress"`

	// CLI messages from stderr
	Stderr []string `json:"stderr"`

	// CLI messages from stdout
	Stdout []string `json:"stdout"`
}

// Validate validates this job result
func (m *JobResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobResult) UnmarshalBinary(b []byte) error {
	var res JobResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}