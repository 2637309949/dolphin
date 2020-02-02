// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package srv

import (
	"context"

	"go.uber.org/multierr"
)

// A Hook is a pair of start and stop callbacks, either of which can be nil,
// plus a string identifying the supplier of the hook.
type Hook struct {
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
}

// Lifecycle interface coordinates application lifecycle hooks.
type Lifecycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Append(h Hook)
}

// lifecycleWrapper coordinates application lifecycle hooks.
type lifecycleWrapper struct {
	hooks      []Hook
	numStarted int
}

// Append adds a Hook to the lifecycle.
func (l *lifecycleWrapper) Append(hook Hook) {
	l.hooks = append(l.hooks, hook)
}

// Start runs all OnStart hooks, returning immediately if it encounters an
// error.
func (l *lifecycleWrapper) Start(ctx context.Context) error {
	for _, hook := range l.hooks {
		if hook.OnStart != nil {
			if err := hook.OnStart(ctx); err != nil {
				if stopErr := hook.OnStop(ctx); stopErr != nil {
					return multierr.Append(err, stopErr)
				}
				return err
			}
		}
		l.numStarted++
	}
	return nil
}

// Stop runs any OnStop hooks whose OnStart counterpart succeeded. OnStop
// hooks run in reverse order.
func (l *lifecycleWrapper) Stop(ctx context.Context) error {
	var errs []error
	// Run backward from last successful OnStart.
	for ; l.numStarted > 0; l.numStarted-- {
		hook := l.hooks[l.numStarted-1]
		if hook.OnStop == nil {
			continue
		}
		if err := hook.OnStop(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	return multierr.Combine(errs...)
}
