package option

// LoadBalanceOutboundOptions is the options for balancer outbound
type LoadBalanceOutboundOptions struct {
	GroupCommonOption
	Check HealthCheckOptions     `json:"check,omitempty"`
	Pick  LoadBalancePickOptions `json:"pick,omitempty"`
}

// LoadBalancePickOptions is the options for balancer outbound picking
type LoadBalancePickOptions struct {
	// load balance objective
	Objective string `json:"objective,omitempty"`
	// pick strategy
	Strategy string `json:"strategy,omitempty"`
	// max acceptable failures
	MaxFail uint `json:"max_fail,omitempty"`
	// max acceptable rtt. defalut 0
	MaxRTT Duration `json:"max_rtt,omitempty"`
	// expected nodes count to select
	Expected uint `json:"expected,omitempty"`
	// ping rtt baselines
	Baselines []Duration `json:"baselines,omitempty"`
}

// HealthCheckOptions is the settings for health check
type HealthCheckOptions struct {
	Interval    Duration `json:"interval"`
	Sampling    uint     `json:"sampling"`
	Destination string   `json:"destination"`
	DetourOf    []string `json:"detour_of,omitempty"`
}
