package plugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
	"openops/pkg/plugin/shared"
	"os/exec"
)

/*
  @Author : lanyulei
  @Desc :
*/

type RunPlugin struct {
	BaseClient *plugin.Client
	Raw        interface{}
}

func New(pluginPath string) shared.CloudProvider {
	var (
		err       error
		runPlugin RunPlugin
		rpcClient plugin.ClientProtocol
	)

	shared.Handshake.MagicCookieValue = viper.GetString("plugin.magicCookieValue")

	runPlugin.BaseClient = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		VersionedPlugins: shared.PluginSets,
		Cmd:              exec.Command("sh", "-c", pluginPath),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
			plugin.ProtocolGRPC,
		},
	})

	rpcClient, err = runPlugin.BaseClient.Client()
	if err != nil {
		logger.Errorf("failed to generate rpc client instance error: %v", err)
		return nil
	}

	runPlugin.Raw, err = rpcClient.Dispense(shared.CloudProviderName)
	if err != nil {
		logger.Errorf("failed to dispense plugin error: %v", err)
		return nil
	}

	return &runPlugin
}

func (r *RunPlugin) List(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	defer r.BaseClient.Kill()
	result, err = r.Raw.(shared.CloudProvider).List(ctx, resource, region, handleType, data)
	if err != nil {
		logger.Errorf("failed to list result error: %v", err)
	}
	return
}

func (r *RunPlugin) Get(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	defer r.BaseClient.Kill()
	result, err = r.Raw.(shared.CloudProvider).Get(ctx, resource, region, handleType, data)
	if err != nil {
		logger.Errorf("failed to get result error: %v", err)
	}
	return
}

func (r *RunPlugin) Post(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	defer r.BaseClient.Kill()
	result, err = r.Raw.(shared.CloudProvider).Post(ctx, resource, region, handleType, data)
	return
}

func (r *RunPlugin) Put(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	defer r.BaseClient.Kill()
	result, err = r.Raw.(shared.CloudProvider).Put(ctx, resource, region, handleType, data)
	return
}

func (r *RunPlugin) Delete(ctx context.Context, resource, region, handleType string, data []byte) (result []byte, err error) {
	defer r.BaseClient.Kill()
	result, err = r.Raw.(shared.CloudProvider).Delete(ctx, resource, region, handleType, data)
	return
}
