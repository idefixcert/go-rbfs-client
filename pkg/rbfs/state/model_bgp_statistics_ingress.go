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

type BgpStatisticsIngress struct {
	// Number of received open messages.
	OpenCount int32 `json:"open_count,omitempty"`
	// Number of received update messages.
	UpdateCount int32 `json:"update_count,omitempty"`
	// Number of received keep-alive messages.
	KeepAliveCount int32 `json:"keep_alive_count,omitempty"`
	// Number of received notify messages.
	NotifyCount int32 `json:"notify_count,omitempty"`
	// Number of received route-refresh messages.
	RouteRefreshCount int32 `json:"route_refresh_count,omitempty"`
}