package rpc

import (
	"demo/rpc/proto"

	"golang.org/x/net/context"
)

// Message defined
type Message struct{}

// SendMail defined
func (s *Message) SendMail(ctx context.Context, in *proto.Mail) (*proto.Reply, error) {
	return &proto.Reply{}, nil
}
