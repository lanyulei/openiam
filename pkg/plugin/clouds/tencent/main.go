package main

import (
	"context"
	"openops/pkg/plugin/clouds"
	"openops/pkg/plugin/clouds/tencent/pkg"
	"openops/pkg/plugin/shared"

	"github.com/hashicorp/go-plugin"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Tencent struct{}

func (Tencent) List(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	var (
		handler pkg.HandlerInterface
	)

	handler, err = pkg.NewHandler(clouds.CloudResourceType(resource), region, clouds.HandleType(handleType), data)
	if err != nil {
		return
	}

	result, err = handler.Get(ctx)
	if err != nil {
		return
	}

	return
}

func (Tencent) Get(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	var (
		handler pkg.HandlerInterface
	)

	handler, err = pkg.NewHandler(clouds.CloudResourceType(resource), region, clouds.HandleType(handleType), data)
	if err != nil {
		return
	}

	result, err = handler.Get(ctx)
	if err != nil {
		return
	}

	return
}

func (Tencent) Post(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	var (
		handler pkg.HandlerInterface
	)

	handler, err = pkg.NewHandler(clouds.CloudResourceType(resource), region, clouds.HandleType(handleType), data)
	if err != nil {
		return
	}

	result, err = handler.Post(ctx)
	if err != nil {
		return
	}

	return
}

func (Tencent) Put(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	var (
		handler pkg.HandlerInterface
	)

	handler, err = pkg.NewHandler(clouds.CloudResourceType(resource), region, clouds.HandleType(handleType), data)
	if err != nil {
		return
	}

	result, err = handler.Put(ctx)
	if err != nil {
		return
	}

	return
}

func (Tencent) Delete(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	var (
		handler pkg.HandlerInterface
	)

	handler, err = pkg.NewHandler(clouds.CloudResourceType(resource), region, clouds.HandleType(handleType), data)
	if err != nil {
		return
	}

	result, err = handler.Delete(ctx)
	if err != nil {
		return
	}

	return
}

func main() {
	var (
		serveConfig *plugin.ServeConfig
	)

	// PluginSets is the map of plugins we can dispense.
	PluginSets := map[int]plugin.PluginSet{
		1: {
			shared.CloudProviderName: &shared.CloudGRPCPlugin{Impl: &Tencent{}},
		},
	}

	handshake := shared.Handshake
	handshake.MagicCookieValue = "0a1b9f9b-78f1-4634-a94d-3c79cc857608"

	serveConfig = &plugin.ServeConfig{
		HandshakeConfig:  handshake,
		VersionedPlugins: PluginSets,
		GRPCServer:       plugin.DefaultGRPCServer,
	}

	plugin.Serve(serveConfig)
}
