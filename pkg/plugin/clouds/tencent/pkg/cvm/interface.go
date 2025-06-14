package cvm

import (
	"context"
	"errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvmobject "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"openops/pkg/plugin/clouds/tencent/pkg/common"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	DescribeInstances(ctx context.Context, params map[string]interface{}) (result []byte, err error)
}

type cvm struct {
	client *cvmobject.Client
}

func New(ak, sk, region, endpoint string) (Interface, error) {
	credential := common.CreateConfig(ak, sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	_client, err := cvmobject.NewClient(credential, region, cpf)

	if err != nil {
		return nil, err
	}

	if _client == nil {
		return nil, errors.New("tencent client is nil")
	}

	return &cvm{
		client: _client,
	}, nil
}
