// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"

	"go.uber.org/multierr"
)

type lifeHook interface {
	OnStart(context.Context) error
	OnStop(context.Context) error
}

// Lifecycle interface coordinates application lifecycle hooks.
type Lifecycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Append(h lifeHook)
}

// lifecycleWrapper coordinates application lifecycle hooks.
type lifecycleWrapper struct {
	hooks      []lifeHook
	numStarted int
}

// Append adds a Hook to the lifecycle.
func (l *lifecycleWrapper) Append(hook lifeHook) {
	l.hooks = append(l.hooks, hook)
}

// Start runs all OnStart hooks, returning immediately if it encounters an
// error.
func (l *lifecycleWrapper) Start(ctx context.Context) error {
	for i := range l.hooks {
		if err := l.hooks[i].OnStart(ctx); err != nil {
			if stopErr := l.hooks[i].OnStop(ctx); stopErr != nil {
				return multierr.Append(err, stopErr)
			}
			return err
		}
		l.numStarted++
	}
	return nil
}

// Stop runs any OnStop hooks whose OnStart counterpart succeeded. OnStop
// hooks run in reverse order.
func (l *lifecycleWrapper) Stop(ctx context.Context) error {
	var errs error
	for ; l.numStarted > 0; l.numStarted-- {
		hook := l.hooks[l.numStarted-1]
		if err := hook.OnStop(ctx); err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}
