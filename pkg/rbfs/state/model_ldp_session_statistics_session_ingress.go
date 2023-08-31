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

type LdpSessionStatisticsSessionIngress struct {
	// Number of received init messages.
	Init int `json:"init,omitempty"`
	// Number of received keep-alive messages.
	KeepAlive int `json:"keep_alive,omitempty"`
	// Number of received notify messages.
	Notify int `json:"notify,omitempty"`
	// Number of received label mapping messages.
	LabelMapping int `json:"label_mapping,omitempty"`
	// Number of received label withdraw messages.
	LabelWithdraw int `json:"label_withdraw,omitempty"`
	// Number of received label release messages.
	LabelRelease int `json:"label_release,omitempty"`
	// Number of received address messages.
	LabelAddress int `json:"label_address,omitempty"`
	// Number of received address withdraw messages.
	AddressWithdraw int `json:"address_withdraw,omitempty"`
}
