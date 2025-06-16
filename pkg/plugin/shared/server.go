package shared

import (
	"context"
	"openops/pkg/plugin/proto"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GRPCServer is a gRPC server for Cloud.
type grpcServer struct {
	Impl CloudProvider
	proto.UnimplementedCloudProviderServer
}

func (m *grpcServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	v, err := m.Impl.List(ctx, req.Resource, req.Region, req.HandleType, req.Data)
	return &proto.ListResponse{Result: v}, err
}

func (m *grpcServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(ctx, req.Resource, req.Region, req.HandleType, req.Data)
	return &proto.GetResponse{Result: v}, err
}

func (m *grpcServer) Post(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	v, err := m.Impl.Post(ctx, req.Resource, req.Region, req.HandleType, req.Data)
	return &proto.CreateResponse{Result: v}, err
}

func (m *grpcServer) Put(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	v, err := m.Impl.Put(ctx, req.Resource, req.Region, req.HandleType, req.Data)
	return &proto.UpdateResponse{Result: v}, err
}

func (m *grpcServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	v, err := m.Impl.Delete(ctx, req.Resource, req.Region, req.HandleType, req.Data)
	return &proto.DeleteResponse{Result: v}, err
}
