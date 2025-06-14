package shared

import "context"

/*
  @Author : lanyulei
  @Desc :
*/

type CloudProvider interface {
	Get(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error)
	Post(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error)
	Put(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error)
	Delete(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error)
}
