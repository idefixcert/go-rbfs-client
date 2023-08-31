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

type LdpSession struct {
	// The LSR lable-space identifier.
	LdpId string `json:"ldp_id,omitempty"`
	// The LSR identifier.
	LsrId                           string                           `json:"lsr_id,omitempty"`
	LdpState                        *LdpSessionState                 `json:"ldp_state,omitempty"`
	LdpLabelAdvertismentMode        *LdpLabelAdvertismentMode        `json:"ldp_label_advertisment_mode,omitempty"`
	LdpLabelDistributionControlMode *LdpLabelDistributionControlMode `json:"ldp_label_distribution_control_mode,omitempty"`
	LdpLabelRetentionMode           *LdpLabelRetentionMode           `json:"ldp_label_retention_mode,omitempty"`
	// LDP session flap count
	LdpSessionFlaps int `json:"ldp_session_flaps,omitempty"`
	// Last session state transition.
	LastStateTransition string   `json:"last_state_transition,omitempty"`
	LdpRole             *LdpRole `json:"ldp_role,omitempty"`
	IflName             string   `json:"ifl_name,omitempty"`
	// The local IPv4 address.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// The local IPv6 address.
	Ipv6Address string                `json:"ipv6_address,omitempty"`
	Peer        *LdpSessionPeer       `json:"peer,omitempty"`
	Statistics  *LdpSessionStatistics `json:"statistics,omitempty"`
	Timers      *LdpSessionTimers     `json:"timers,omitempty"`
}
