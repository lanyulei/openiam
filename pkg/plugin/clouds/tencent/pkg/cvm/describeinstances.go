package cvm

import (
	"context"
	"encoding/json"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvmobject "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"openops/pkg/plugin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func (c *cvm) DescribeInstances(ctx context.Context, queries map[string]interface{}) (result []byte, err error) {
	var (
		response    *cvmobject.DescribeInstancesResponse
		_resultList []plugin.Response
		offset      int64
	)

	request := cvmobject.NewDescribeInstancesRequest()
	request.Limit = common.Int64Ptr(100)

	for {
		request.Offset = common.Int64Ptr(offset)

		response, err = c.client.DescribeInstances(request)
		if err != nil {
			return
		}

		if len(response.Response.InstanceSet) == 0 {
			break
		}

		for _, _instance := range response.Response.InstanceSet {
			_instanceBytes, _err := json.Marshal(_instance)
			if _err != nil {
				return
			}
			_resultList = append(_resultList, plugin.Response{
				UniqueId: *_instance.InstanceId,
				Content:  _instanceBytes,
			})
		}

		offset += 1
	}

	if len(_resultList) > 0 {
		result, err = json.Marshal(_resultList)
		if err != nil {
			return
		}
	}

	return
}
