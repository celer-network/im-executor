package types

import (
	fmt "fmt"
	math "math"
	"time"

	"github.com/celer-network/goutils/log"
)

const (
	DEFAULT_INITIAL_BACKOFF_SECONDS = 30
	DEFAULT_INTERVAL_SECONDS        = 30
	DEFAULT_SCALING_FACTOR          = 1  // linear backoff
	DEFAULT_MAX_ATTEMPTS            = 15 // TODO remove types.MaxExecuteRetry and implement retry max attempts
)

type RetryConfig struct {
	// the initial amount of seconds to wait before the first retry
	InitialBackoffSeconds uint64 `mapstructure:"initial_backoff_seconds"`

	// amount of seconds in between retries, affected by scaling factor
	IntervalSeconds uint64 `mapstructure:"interval_seconds"`

	// for any retry count n, x = scaling factor, s = interval seconds,
	// actual retry interval = s * n^x
	// e.g. set x = 0 for constant backoff, 1 for linear backoff, 2 for quadratic backoff, etc...
	ScalingFactor float64 `mapstructure:"scaling_factor"`

	// number of retries before giving up
	MaxAttempts int `mapstructure:"max_attempts"`
}

func (c *RetryConfig) ApplyDefaults() {
	if c.InitialBackoffSeconds == 0 {
		c.InitialBackoffSeconds = DEFAULT_INITIAL_BACKOFF_SECONDS
	}
	if c.IntervalSeconds == 0 {
		log.Infof("[executor] retry initial backoff not set or is 0, using default %d", DEFAULT_INTERVAL_SECONDS)
		c.IntervalSeconds = DEFAULT_INTERVAL_SECONDS
	}
	if c.ScalingFactor == 0 {
		log.Infof("[executor] retry backoff scaling factor not set or is 0, using default %d", DEFAULT_SCALING_FACTOR)
		c.ScalingFactor = DEFAULT_SCALING_FACTOR
	}
	if c.MaxAttempts == 0 {
		log.Infof("[executor] retry max attempts not set or is 0, using default %d", DEFAULT_MAX_ATTEMPTS)
		c.MaxAttempts = DEFAULT_MAX_ATTEMPTS
	}
}

func (c *RetryConfig) GetInitialBackoff() time.Duration {
	return time.Duration(c.InitialBackoffSeconds) * time.Second
}

func (c *RetryConfig) GetIntervalSeconds() time.Duration {
	return time.Duration(c.IntervalSeconds) * time.Second
}

func (c *RetryConfig) ShouldBackoff(attempts uint64, lastTryTime time.Time) bool {
	now := time.Now()
	if attempts == 1 && now.Before(lastTryTime.Add(c.GetInitialBackoff())) {
		// backoff when it's the first retry and initial backoff window has not passed
		return true
	}
	if attempts > 0 {
		curAttempt := float64(attempts + 1)
		// for any retry count n, x = scaling factor, s = interval seconds,
		// actual retry interval = s * n^x
		actualInterval := c.GetIntervalSeconds() * time.Duration(math.Pow(curAttempt, c.ScalingFactor))
		if now.Before(lastTryTime.Add(actualInterval)) {
			return true
		}
	}
	return false
}

func (c *RetryConfig) String() string {
	if c == nil {
		return "retry config: nil"
	}
	return fmt.Sprintf("retry config: initial backoff %ds, interval %ds, scaling factor %f",
		c.InitialBackoffSeconds, c.IntervalSeconds, c.ScalingFactor)
}

func (c *RetryConfig) PrettyLog() {
	log.Infoln(c.String())
}
