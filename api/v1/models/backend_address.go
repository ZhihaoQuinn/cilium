// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackendAddress Service backend address
//
// swagger:model BackendAddress
type BackendAddress struct {

	// Layer 3 address
	// Required: true
	IP *string `json:"ip"`

	// Optional name of the node on which this backend runs
	NodeName string `json:"nodeName,omitempty"`

	// Layer 4 port number
	Port uint16 `json:"port,omitempty"`

	// Layer 4 protocol (TCP, UDP, etc)
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this backend address
func (m *BackendAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackendAddress) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendAddress) UnmarshalBinary(b []byte) error {
	var res BackendAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
