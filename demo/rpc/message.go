package rpc

import (
	"demo/rpc/proto"

	"golang.org/x/net/context"
)

// Message defined
type Message struct{}

// SendMail defined
func (s *Message) SendMail(ctx context.Context, in *proto.Article) (*proto.Success, error) {
	return &proto.Success{}, nil
}
