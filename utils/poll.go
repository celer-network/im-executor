package utils

import (
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
)

type pollOpt struct {
	maxRetry int
	interval time.Duration
	deadline time.Time
	message  string
}

type PollOpts struct {
	apply func(opt *pollOpt)
}

func newPollOpts(f func(opt *pollOpt)) *PollOpts {
	return &PollOpts{apply: f}
}

func WithInterval(interval time.Duration) *PollOpts {
	return newPollOpts(func(opt *pollOpt) { opt.interval = interval })
}

func WithTimeout(timeout time.Duration) *PollOpts {
	return newPollOpts(func(opt *pollOpt) { opt.deadline = time.Now().Add(timeout) })
}

func WithMaxRetry(maxRetry int) *PollOpts {
	return newPollOpts(func(opt *pollOpt) { opt.maxRetry = maxRetry })
}

func WithMessage(msg string, args ...interface{}) *PollOpts {
	return newPollOpts(func(opt *pollOpt) { opt.message = fmt.Sprintf(msg, args...) })
}

func Poll[T any](run func() (result *T, stop bool), opts ...*PollOpts) (*T, error) {
	pollOpts := &pollOpt{
		maxRetry: 10,
		interval: 3 * time.Second,
		message:  "poll attempt",
	}
	for _, opt := range opts {
		opt.apply(pollOpts)
	}
	for i := 0; i < pollOpts.maxRetry; i++ {
		log.Infof(pollOpts.message+" [%d/%d]", i+1, pollOpts.maxRetry)
		result, stop := run()
		if stop {
			return result, nil
		}
		time.Sleep(pollOpts.interval)
	}
	return nil, fmt.Errorf("poll max retry hit")
}
