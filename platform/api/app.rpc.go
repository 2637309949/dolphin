// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// RpcHandler defined
type RpcHandler interface {
	RegisterServer(func(*grpc.Server))
	OnStart(context.Context) error
	OnStop(context.Context) error
}

type gHandler struct {
	grpc *grpc.Server
	net  net.Listener
}

func (gh *gHandler) RegisterServer(f func(*grpc.Server)) {
	f(gh.grpc)
}

func (gh *gHandler) OnStart(ctx context.Context) error {
	go func() {
		logrus.Infof("grpc listen on port:%v", viper.GetString("rpc.port"))
		if err := gh.grpc.Serve(gh.net); err != nil {
			logrus.Fatal(err)
		}
	}()
	return nil
}

func (gh *gHandler) OnStop(ctx context.Context) error {
	if err := gh.net.Close(); err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}

// RpcClientDailTimeOut defined TODO
var RpcClientDailTimeOut = 3 * time.Second

// NewRpcClient defined TODO
func NewRpcClient(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	options := append(opts, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(plugin.RpcSrvTrace),
	}...)
	ctx, cancel := context.WithTimeout(context.Background(), RpcClientDailTimeOut)
	defer cancel()
	conn, err := grpc.DialContext(ctx, viper.GetString("rpc.domain_srv"), options...)
	return conn, err
}

// NewRpcHandler defined TODO
func NewRpcHandler() RpcHandler {
	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(plugin.RpcCliTrace(viper.GetString("app.name"))),
	}
	net := util.EnsureLeft(net.Listen("tcp", fmt.Sprintf(":%v", viper.GetString("rpc.port")))).(net.Listener)
	return &gHandler{net: net, grpc: grpc.NewServer(options...)}
}
