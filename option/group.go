package option

type SelectorOutboundOptions struct {
	GroupCommonOption
	Default                   string `json:"default,omitempty"`
	InterruptExistConnections bool   `json:"interrupt_exist_connections,omitempty"`
}

type URLTestOutboundOptions struct {
	GroupCommonOption
	URL       string   `json:"url,omitempty"`
	Interval  Duration `json:"interval,omitempty"`
	Tolerance uint16   `json:"tolerance,omitempty"`
}

// GroupCommonOption is the common options for group outbounds
type GroupCommonOption struct {
	Outbounds []string `json:"outbounds"`
	Providers []string `json:"providers"`
	Exclude   string   `json:"exclude,omitempty"`
	Include   string   `json:"include,omitempty"`
}
