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

// Contains the number of routes in the RIB grouped by originating protocol
type RouteCounts struct {
	// ARP ND routes count
	ArpNd int `json:"arp_nd,omitempty"`
	// direct routes count
	Direct int `json:"direct,omitempty"`
	// static routes count
	Static int `json:"static,omitempty"`
	// BGP routes count
	Bgp int `json:"bgp,omitempty"`
	// Local BGP routes count
	BgpLocal int `json:"bgp_local,omitempty"`
	// Local origin BGP routes count
	BgpLocalOrigin int `json:"bgp_local_origin,omitempty"`
	// ISIS routes count
	Isis int `json:"isis,omitempty"`
	// OSPF routes count
	Ospf int `json:"ospf,omitempty"`
	// LDP routes count
	Ldp int `json:"ldp,omitempty"`
	// PIM routes count
	Pim int `json:"pim,omitempty"`
	// IGMP routes count
	Igmp int `json:"igmp,omitempty"`
	// L2TPv2 routes count
	L2tpv2 int `json:"l2tpv2,omitempty"`
	// IPoE routes count
	Ipoe int `json:"ipoe,omitempty"`
	// PPP routes count
	Ppp int `json:"ppp,omitempty"`
	// DHCP routes count
	Dhcp int `json:"dhcp,omitempty"`
	// local routes count
	Local int `json:"local,omitempty"`
	// RIB routes count
	Rib int `json:"rib,omitempty"`
	// mRIB routes count
	Mrib int `json:"mrib,omitempty"`
	// L3VPN routes count
	L3All int `json:"l3_all,omitempty"`
	// L2VPN routes count
	L2All int `json:"l2_all,omitempty"`
	// Total routes count
	Total int `json:"total,omitempty"`
}
