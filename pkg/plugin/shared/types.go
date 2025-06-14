package shared

import (
	"github.com/hashicorp/go-plugin"
)

/*
  @Author : lanyulei
  @Desc :
*/

const CloudProviderName = "CloudProvider"

var (
	Handshake = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "OPENOPS_CLOUD_PLUGIN",
		MagicCookieValue: "",
	}

	// PluginSets is the map of plugins we can dispense.
	PluginSets = map[int]plugin.PluginSet{
		1: {
			CloudProviderName: &CloudGRPCPlugin{},
		},
	}
)
