package dialer

import (
	"time"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/option"
	dns "github.com/sagernet/sing-dns"
	"github.com/sagernet/sing/common"
	N "github.com/sagernet/sing/common/network"
)

func New(router adapter.Router, options option.DialerOptions) (N.Dialer, error) {
	return new(router, options, "")
}

func MustNewChainRedirectable(router adapter.Router, tag string, options option.DialerOptions) N.Dialer {
	return common.Must1(NewChainRedirectable(router, tag, options))
}

func NewChainRedirectable(router adapter.Router, tag string, options option.DialerOptions) (N.Dialer, error) {
	return new(router, options, tag)
}

func new(router adapter.Router, options option.DialerOptions, redirectableTag string) (N.Dialer, error) {
	var (
		dialer N.Dialer
		err    error
	)
	if options.Detour == "" {
		dialer, err = NewDefault(router, options)
		if err != nil {
			return nil, err
		}
		if redirectableTag != "" {
			dialer = NewChainRedirectDialer(redirectableTag, dialer, dialer)
		}
		if options.IsWireGuardListener {
			return dialer, nil
		}
	} else {
		dialer = NewDetour(router, options.Detour)
		if redirectableTag != "" {
			defDialer, err := NewDefault(router, options)
			if err != nil {
				return nil, err
			}
			dialer = NewChainRedirectDialer(redirectableTag, dialer, defDialer)
		}
	}
	domainStrategy := dns.DomainStrategy(options.DomainStrategy)
	if domainStrategy != dns.DomainStrategyAsIS || options.Detour == "" {
		dialer = NewResolveDialer(
			router,
			dialer,
			options.Detour == "" && !options.TCPFastOpen,
			domainStrategy,
			time.Duration(options.FallbackDelay))
	}
	return dialer, nil
}
