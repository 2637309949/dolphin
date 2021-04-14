// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// Dol struct
type Dol struct {
	fx *FX
}

// Use defined use plugin
func Use(opts ...Option) {
	d.Use(opts...)
}

// Use defined use plugin
func (d *Dol) Use(opts ...Option) {
	for _, opt := range opts {
		opt.apply(d)
	}
}

// Done with os.Signal
func Done() <-chan os.Signal {
	return d.Done()
}

// Done with os.Signal
func (d *Dol) Done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// Start with Context
func Start(ctx context.Context) error {
	return d.Start(ctx)
}

// Start with Context
func (d *Dol) Start(ctx context.Context) error {
	d.fx.build()
	// lifecycle start
	return withTimeout(ctx, d.fx.lifecycle.Start)
}

// Stop with Context
func Stop(ctx context.Context) error {
	return d.Stop(ctx)
}

// Stop with Context
func (d *Dol) Stop(ctx context.Context) error {
	// lifecycle stop
	return withTimeout(ctx, d.fx.lifecycle.Stop)
}

// Run defined run application
func Run() error {
	return d.Run()
}

// Run defined run application
func (d *Dol) Run() error {
	done := d.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := d.Start(ctx); err != nil {
		logrus.Fatalf("ERROR\t\tFailed to start: %v", err)
	}
	<-done
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := d.Stop(ctx); err != nil {
		logrus.Fatalf("ERROR\t\tFailed to stop cleanly: %v", err)
	}
	return nil
}

// New defined return dol
func New() *Dol {
	dol := &Dol{}
	dol.fx = NewFX()
	dol.fx.provide(func() Lifecycle { return dol.fx.lifecycle })
	return dol
}
