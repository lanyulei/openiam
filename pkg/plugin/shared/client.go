package shared

import (
	"context"
	"openops/pkg/plugin/proto"

	"github.com/lanyulei/toolkit/logger"
	"google.golang.org/grpc"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GRPCClient is a gRPC client for Cloud.
type grpcClient struct {
	client     proto.CloudProviderClient
	clientConn *grpc.ClientConn
	doneCtx    context.Context
}

func (m *grpcClient) List(ctx context.Context, resource, region, handleType string, data []byte) ([]byte, error) {
	var (
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	resp, err := m.client.List(ctx, &proto.ListRequest{
		Resource:   resource,
		Region:     region,
		HandleType: handleType,
		Data:       data,
	})
	if err != nil {
		return []byte(""), err
	}

	return resp.Result, nil
}

func (m *grpcClient) Get(ctx context.Context, resource, region, handleType string, data []byte) ([]byte, error) {
	var (
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	resp, err := m.client.Get(ctx, &proto.GetRequest{
		Resource:   resource,
		Region:     region,
		HandleType: handleType,
		Data:       data,
	})
	if err != nil {
		return []byte(""), err
	}

	return resp.Result, nil
}

func (m *grpcClient) Post(ctx context.Context, resource, region, handleType string, data []byte) ([]byte, error) {
	var (
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	resp, err := m.client.Post(ctx, &proto.CreateRequest{
		Resource:   resource,
		Region:     region,
		HandleType: handleType,
		Data:       data,
	})
	if err != nil {
		logger.Errorf("failed to create resource error: %v", err)
		return []byte(""), err
	}

	return resp.Result, nil
}

func (m *grpcClient) Put(ctx context.Context, resource, region, handleType string, data []byte) ([]byte, error) {
	var (
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	resp, err := m.client.Put(ctx, &proto.UpdateRequest{
		Resource:   resource,
		Region:     region,
		HandleType: handleType,
		Data:       data,
	})
	if err != nil {
		logger.Errorf("failed to update resource error: %v", err)
		return []byte(""), err
	}

	return resp.Result, nil
}

func (m *grpcClient) Delete(ctx context.Context, resource, region, handleType string, data []byte) ([]byte, error) {
	var (
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	resp, err := m.client.Delete(ctx, &proto.DeleteRequest{
		Resource:   resource,
		Region:     region,
		HandleType: handleType,
		Data:       data,
	})
	if err != nil {
		logger.Errorf("failed to delete resource error: %v", err)
		return []byte(""), err
	}

	return resp.Result, nil
}
