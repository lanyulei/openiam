package shared

import (
	"context"
	"openops/pkg/plugin/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

/*
  @Author : lanyulei
  @Desc :
*/

type CloudGRPCPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl CloudProvider
}

func (p *CloudGRPCPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterCloudProviderServer(s, &grpcServer{Impl: p.Impl})
	return nil
}

func (p *CloudGRPCPlugin) GRPCClient(ctx context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &grpcClient{
		client:     proto.NewCloudProviderClient(c),
		clientConn: c,
		doneCtx:    ctx,
	}, nil
}
