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

// TX-related state.
type LaneTx struct {
	// The status of the Loss-of-Signal (LoS) alert flag.
	LosAlert bool `json:"los_alert,omitempty"`
	// The laser power in milliwatts.
	PowerMw float64 `json:"power_mw,omitempty"`
	// The laser power in dbm.
	PowerDbm float64 `json:"power_dbm,omitempty"`
	// The power high alarm threshold in milliwatts.
	PowerHighAlarmThresholdMw float64 `json:"power_high_alarm_threshold_mw,omitempty"`
	// The power high alarm threshold in dBm.
	PowerHighAlarmThresholdDbm float64 `json:"power_high_alarm_threshold_dbm,omitempty"`
	// The power warning threshold in milliwatts.
	PowerHighWarnThresholdMw float64 `json:"power_high_warn_threshold_mw,omitempty"`
	// The power high warning threshold in dBm.
	PowerHighWarnThresholdDbm float64 `json:"power_high_warn_threshold_dbm,omitempty"`
	// The power low alarm threshold in milliwatts.
	PowerLowAlarmThresholdMw float64 `json:"power_low_alarm_threshold_mw,omitempty"`
	// The power low alarm threshold in dBm.
	PowerLowAlarmThresholdDbm float64 `json:"power_low_alarm_threshold_dbm,omitempty"`
	// The power low warning threshold in milliwatts.
	PowerLowWarnThresholdMw float64 `json:"power_low_warn_threshold_mw,omitempty"`
	// The power low warning threshold in dBm.
	PowerLowWarnThresholdDbm float64 `json:"power_low_warn_threshold_dbm,omitempty"`
	// Laser bias current in mA.
	BiasCurrentMa float64 `json:"bias_current_ma,omitempty"`
	// Laser bias current high alarm threshold in mA.
	BiasCurrentHighAlarmThresholdMa float64 `json:"bias_current_high_alarm_threshold_ma,omitempty"`
	// Laser bias current high warning threshold in mA.
	BiasCurrentHighWarningThresholdMa float64 `json:"bias_current_high_warning_threshold_ma,omitempty"`
	// Laser bias current low alarm threshold in mA.
	BiasCurrentLowAlarmThresholdMa float64 `json:"bias_current_low_alarm_threshold_ma,omitempty"`
	// Laser bias current low warning threshold in mA.
	BiasCurrentLowWarningThresholdMa float64 `json:"bias_current_low_warning_threshold_ma,omitempty"`
}
