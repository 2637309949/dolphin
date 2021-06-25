// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"fmt"
	"net"

	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// RPCHandler defined
type RPCHandler interface {
	RegisterServer(func(*grpc.Server))
	OnStart(context.Context) error
	OnStop(context.Context) error
}

type grpcHandler struct {
	grpc *grpc.Server
	net  net.Listener
}

func (gh *grpcHandler) RegisterServer(f func(*grpc.Server)) {
	f(gh.grpc)
}

func (gh *grpcHandler) OnStart(ctx context.Context) error {
	logrus.Infof("grpc listen on port:%v", viper.GetString("rpc.port"))
	go func() {
		if err := gh.grpc.Serve(gh.net); err != nil {
			logrus.Fatal(err)
		}
	}()
	return nil
}

func (gh *grpcHandler) OnStop(ctx context.Context) error {
	if err := gh.net.Close(); err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}

// NewGRPCHandler defined
func NewGRPCHandler(dol *Dolphin) RPCHandler {
	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(plugin.RpcCliTrace(viper.GetString("app.name"))),
	}
	net := util.EnsureLeft(net.Listen("tcp", fmt.Sprintf(":%v", viper.GetString("rpc.port")))).(net.Listener)
	return &grpcHandler{net: net, grpc: grpc.NewServer(options...)}
}
