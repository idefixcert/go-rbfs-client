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

// The hardware chassis.
type Chassis struct {
	// The product name.
	ProductName string `json:"product_name,omitempty"`
	// The vendor name.
	VendorName string `json:"vendor_name,omitempty"`
	// The system MAC address.
	Mac string `json:"mac,omitempty"`
	// The date when the system has been manufactured.
	ManufactureDate string `json:"manufacture_date,omitempty"`
	// The manufacturer name.
	Manufacturer string `json:"manufacturer,omitempty"`
	// The ONIE SW version.
	OnieVersion string `json:"onie_version,omitempty"`
	// The diag SW version.
	DiagVersion string `json:"diag_version,omitempty"`
	// The part number.
	PartNumber string `json:"part_number,omitempty"`
	// The platform name.
	PlatformName string `json:"platform_name,omitempty"`
	// The chassis serial number.
	SerialNumber string `json:"serial_number,omitempty"`
	// The country where the system has been manufactured.
	CountryCode string `json:"country_code,omitempty"`
}
