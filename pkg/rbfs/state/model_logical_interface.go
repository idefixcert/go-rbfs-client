/*
 * RBFS Operational State API
 *
 * This contract describes the RBFS Operational State API contract defined by RBMS, the RtBrick Management System. This API is a _consumer-driven_ API, which means that all changes to this API **must be approved** by RBMS, the consumer of this API to avoid compatibility issues.  The API is kept backwards-compatible and anyone is allowed to _use_ this API.  The consumer of the API _must_ ignore additional attributes not explained in this specification. Additional attributes are _not_ considered violating backwards compatibility. In contrary, additional attributes allow extending the API while preserving backward compatibility.
 *
 * API version: 1.0.0
 * Contact: martin@rtbrick.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package state

// A logical interface.
type LogicalInterface struct {
	IflName string `json:"ifl_name,omitempty"`
	IfpName string `json:"ifp_name,omitempty"`
	// The logical interface type.
	IflType string `json:"ifl_type,omitempty"`
	// The logical intreface alias.
	IflAlias     string `json:"ifl_alias,omitempty"`
	InstanceName string `json:"instance_name,omitempty"`
	// The MAC address on the logical interface.
	MacAddress string `json:"mac_address,omitempty"`
	// The operational interface state.
	OperationalState string `json:"operational_state,omitempty"`
	// The administrative interface state.
	AdministrativeState string `json:"administrative_state,omitempty"`
	// The assigned IPv4 address.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// Whether IPv4 is enabled or disabled on this logical interface.
	Ipv4State string `json:"ipv4_state,omitempty"`
	// IPv4 maximum transfer unit (MTU) size in bytes.
	Ipv4Mtu int32 `json:"ipv4_mtu,omitempty"`
	// The assigned IPv6 address.
	Ipv6Address string `json:"ipv6_address,omitempty"`
	// Whether IPv6 is enabled or disabled on this logical interface.
	Ipv6State string `json:"ipv6_state,omitempty"`
	// IPv6 maximum transfer unit (MTU) size in bytes.
	Ipv6Mtu int32 `json:"ipv6_mtu,omitempty"`
	// MPLS maximum transfer unit (MTU) size in bytes.
	MplsMtu int32 `json:"mpls_mtu,omitempty"`
	// Whether MPLS is enabled or disabled on this logical interface.
	MplsState string `json:"mpls_state,omitempty"`
	// ISO maximum transfer unit (MTU) size.
	IsoMtu int32 `json:"iso_mtu,omitempty"`
	// Whether ISO is enabled or disabled on this logical interface.
	IsoState string `json:"iso_state,omitempty"`
	// The assigned VLANs.  The array is filled beginning with the outermost VLANS:  - The array is _empty_ or _omitted_ for untagged interfaces.  - The array contains the VLAN-ID for single tagged interfaces.  - The array contains the outer VLAN-ID followed by the inner VLAN-ID for double tagged interfaces.
	Vlans       []int32                      `json:"vlans,omitempty"`
	IflCounters *LogicalInterfaceIflCounters `json:"ifl_counters,omitempty"`
}