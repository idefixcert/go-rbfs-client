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

// LDP session
type LdpInstanceSessionsSessions struct {
	// Number of LDP sessions in NONEXISTENT state, i.e. session establishment has not started.
	NonExistent int `json:"non_existent,omitempty"`
	// Number of LDP sessions in OPENSENT state.
	OpensentCount int `json:"opensent_count,omitempty"`
	// Number of LDP sessions in OPENREC state.
	OpenrecCount int `json:"openrec_count,omitempty"`
	// Number of initialized LDP session.
	InitializedCount int `json:"initialized_count,omitempty"`
	// Number of operational LDP sessions.
	OperationalCount int `json:"operational_count,omitempty"`
	// List of LDP sessions.
	Sessions []LdpSession `json:"sessions,omitempty"`
}
