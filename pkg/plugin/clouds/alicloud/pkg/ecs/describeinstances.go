package ecs

import (
	"context"
	"encoding/json"
	"openops/pkg/plugin"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/*
  @Author : lanyulei
  @Desc :
*/

func (e *ecs) DescribeInstances(ctx context.Context, queries map[string]interface{}) (result []byte, err error) {
	var (
		_result     map[string]interface{}
		_resultList []plugin.Response
		nextToken   string
	)

	openapiParams := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2014-05-26"),
		Protocol:    tea.String("HTTPS"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		Pathname:    tea.String("/"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}

	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{}

	queries["PageSize"] = tea.Int(100)

	for {
		if nextToken != "" {
			queries["NextToken"] = tea.String(nextToken)
		}

		if queries != nil {
			request.Query = openapiutil.Query(queries)
		}

		_result, err = e.client.CallApi(openapiParams, request, runtime)
		if err != nil {
			return
		}

		nextToken = _result["body"].(map[string]interface{})["NextToken"].(string)
		if nextToken == "" {
			break
		}

		_instances := _result["body"].(map[string]interface{})["Instances"].(map[string]interface{})["Instance"].([]interface{})

		for _, _instance := range _instances {
			_instanceBytes, _err := json.Marshal(_instance)
			if _err != nil {
				return
			}
			_resultList = append(_resultList, plugin.Response{
				UniqueId: _instance.(map[string]interface{})["InstanceId"].(string),
				Content:  _instanceBytes,
			})
		}
	}

	if len(_resultList) > 0 {
		result, err = json.Marshal(_resultList)
		if err != nil {
			return
		}
	}

	return
}
