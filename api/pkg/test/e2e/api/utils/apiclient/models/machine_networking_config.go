// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MachineNetworkingConfig MachineNetworkingConfig specifies the networking parameters used for IPAM.
//
// swagger:model MachineNetworkingConfig
type MachineNetworkingConfig struct {

	// c ID r
	CIDR string `json:"cidr,omitempty"`

	// DNS servers
	DNSServers []string `json:"dnsServers"`

	// gateway
	Gateway string `json:"gateway,omitempty"`
}

// Validate validates this machine networking config
func (m *MachineNetworkingConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MachineNetworkingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineNetworkingConfig) UnmarshalBinary(b []byte) error {
	var res MachineNetworkingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
