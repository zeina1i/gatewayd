package config

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog"
)

// GetVerificationPolicy returns the hook verification policy from plugin config file.
func (p PluginConfig) GetVerificationPolicy() Policy {
	switch p.VerificationPolicy {
	case "ignore":
		return Ignore
	case "abort":
		return Abort
	case "remove":
		return Remove
	default:
		return PassDown
	}
}

// GetPluginCompatPolicy returns the plugin compatibility policy from plugin config file.
func (p PluginConfig) GetPluginCompatPolicy() CompatPolicy {
	switch p.CompatibilityPolicy {
	case "strict":
		return Strict
	case "loose":
		return Loose
	default:
		return Strict
	}
}

// GetLoadBalancer returns the load balancing algorithm to use.
func (s Server) GetLoadBalancer() gnet.LoadBalancing {
	switch s.LoadBalancer {
	case "roundrobin":
		return gnet.RoundRobin
	case "leastconnections":
		return gnet.LeastConnections
	case "sourceaddrhash":
		return gnet.SourceAddrHash
	default:
		return gnet.RoundRobin
	}
}

// GetTCPNoDelay returns the TCP no delay option from config file.
func (s Server) GetTCPNoDelay() gnet.TCPSocketOpt {
	if s.TCPNoDelay {
		return gnet.TCPNoDelay
	}

	return gnet.TCPDelay
}

// GetSize returns the pool size from config file.
func (p Pool) GetSize() int {
	if p.Size == 0 {
		return DefaultPoolSize
	}

	// Minimum pool size is 2.
	if p.Size < MinimumPoolSize {
		p.Size = MinimumPoolSize
	}

	return p.Size
}

// GetOutput returns the logger output from config file.
func (l Logger) GetOutput() LogOutput {
	switch l.Output {
	case "file":
		return File
	case "stdout":
		return Stdout
	case "stderr":
		return Stderr
	default:
		return Console
	}
}

// GetTimeFormat returns the logger time format from config file.
func (l Logger) GetTimeFormat() string {
	switch l.TimeFormat {
	case "unixms":
		return zerolog.TimeFormatUnixMs
	case "unixmicro":
		return zerolog.TimeFormatUnixMicro
	case "unixnano":
		return zerolog.TimeFormatUnixNano
	case "unix":
		return zerolog.TimeFormatUnix
	default:
		return zerolog.TimeFormatUnix
	}
}

// GetLevel returns the logger level from config file.
func (l Logger) GetLevel() zerolog.Level {
	switch l.Level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disabled":
		return zerolog.Disabled
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
