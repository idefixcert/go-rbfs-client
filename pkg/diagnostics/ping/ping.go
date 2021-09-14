package ping

import (
	"fmt"
	"net"
	"time"
)

type (
	// Ping contains all arguments to ping a destination IP address or hostname
	Ping struct {
		destination     string
		sourceInterface string
		sourceIP        net.IP
		instanceName    string
		count           int32
		interval        time.Duration
	}

	// Option applies a ping command argument
	Option func(*Ping) error
)

// NewPing creates a new ping command
func NewPing(destination string, options ...Option) (*Ping, error) {
	p := &Ping{
		destination:  destination,
		instanceName: "default",
		count:        5,
		interval:     time.Second,
	}

	// Apply all given ping option
	for _, option := range options {
		if err := option(p); err != nil {
			return nil, err
		}
	}

	return p, nil
}

// SourceIP specifies the source IP address
func SourceIP(ipAddress net.IP) Option {
	return func(p *Ping) error {
		if ipAddress != nil {
			if p.sourceInterface != "" {
				return fmt.Errorf("source interface and source IP are mutual exclusive")
			}
			p.sourceIP = ipAddress
		}
		return nil
	}
}

// SourceInterface sets the ping source interface name.
// Source interface and source IP are mutual exclusive!
func SourceInterface(name string) Option {
	return func(p *Ping) error {
		if name != "" {
			if p.sourceIP != nil {
				return fmt.Errorf("source interface and source IP are mutual exclusive")
			}
			p.sourceInterface = name
		}
		return nil
	}
}

// Count sets the number of pings to be sent.
func Count(count int32) Option {
	return func(p *Ping) error {
		if count <= 0 {
			return fmt.Errorf("count value must be greater than 0")
		}

		const maxAllowedPings = 10
		if count > maxAllowedPings {
			return fmt.Errorf("count value must be less or equal than %d", maxAllowedPings)
		}

		p.count = count
		return nil
	}
}

// Interval sets the interval between two pings.
// The accepted interval range is between 1ms and 10 seconds.
func Interval(interval time.Duration) Option {
	return func(p *Ping) error {
		if interval < 1*time.Millisecond {
			return fmt.Errorf("interval must not be less than 1ms")
		}
		if interval > 10*time.Second {
			return fmt.Errorf("interval must not exceed 10s")
		}
		p.interval = interval
		return nil
	}
}

// InstanceName sets the routing instance name to run the ping command.
func InstanceName(instanceName string) Option {
	return func(p *Ping) error {
		if instanceName == "" {
			return fmt.Errorf("instance name must not be empty")
		}
		p.instanceName = instanceName
		return nil
	}
}

func (p *Ping) SourceInterface() string {
	return p.sourceInterface
}

func (p *Ping) SourceIP() net.IP {
	return p.sourceIP
}

func (p *Ping) Count() int32 {
	return p.count
}

func (p *Ping) Interval() time.Duration {
	return p.interval
}

func (p *Ping) Destination() string {
	return p.destination
}
