package region

import (
	"context"
	"encoding/json"

	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/*
  @Author : lanyulei
  @Desc :
*/

func (r *region) DescribeRegions(ctx context.Context, queries map[string]interface{}) (result []byte, err error) {
	var (
		_result map[string]interface{}
	)

	openapiParams := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-01-01"),
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

	if queries != nil {
		request.Query = openapiutil.Query(queries)
	}

	_result, err = r.client.CallApi(openapiParams, request, runtime)
	if err != nil {
		return
	}

	result, err = json.Marshal(_result)
	if err != nil {
		return
	}

	return
}
