package region

import (
	"context"
	"errors"
	"openops/pkg/plugin/clouds/alicloud/pkg/common"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	DescribeRegions(ctx context.Context, params map[string]interface{}) (result []byte, err error)
}

type region struct {
	client *openapi.Client
}

func New(ak, sk, endpoint string) (Interface, error) {
	config := common.CreateConfig(ak, sk)
	config.Endpoint = tea.String(endpoint)
	_client, _err := openapi.NewClient(config)
	if _err != nil {
		return nil, _err
	}
	if _client == nil {
		return nil, errors.New("alicloud client is nil")
	}
	return &region{
		client: _client,
	}, nil
}
