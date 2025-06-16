package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"openops/pkg/plugin/clouds"
	"openops/pkg/plugin/clouds/alicloud/pkg/ecs"
	"os"

	"github.com/lanyulei/toolkit/cipher"
)

/*
  @Author : lanyulei
  @Desc :
*/

type HandlerInterface interface {
	List(ctx context.Context) (result []byte, err error)
	Get(ctx context.Context) (result []byte, err error)
	Post(ctx context.Context) (result []byte, err error)
	Put(ctx context.Context) (result []byte, err error)
	Delete(ctx context.Context) (result []byte, err error)
}

type handler struct {
	resource   clouds.CloudResourceType
	region     string
	ak         string
	sk         string
	data       map[string]interface{}
	handleType clouds.HandleType
}

func NewHandler(resource clouds.CloudResourceType, region string, handleType clouds.HandleType, data []byte) (HandlerInterface, error) {
	var (
		err              error
		akBytes, skBytes []byte
	)

	_handler := handler{
		resource:   resource,
		region:     region,
		handleType: handleType,
		ak:         os.Getenv("ALIYUN_ACCESS_KEY_ID"),
		sk:         os.Getenv("ALIYUN_ACCESS_KEY_SECRET"),
	}

	_handler.data = make(map[string]interface{})

	cryptoEnable := os.Getenv("OPENOPS_CRYPTO_ENABLE") // true or false
	if cryptoEnable == "true" {
		cryptoAesKey := os.Getenv("OPENOPS_CRYPTO_AES_KEY")
		if cryptoAesKey != "" && _handler.ak != "" && _handler.sk != "" {
			// ak, sk 任何一个为空，则表示需要走默认权限认证
			akBytes, err = cipher.AesDecryptCBC([]byte(cryptoAesKey), []byte(_handler.ak))
			if err == nil {
				_handler.ak = string(akBytes)
			}
			skBytes, err = cipher.AesDecryptCBC([]byte(cryptoAesKey), []byte(_handler.sk))
			if err == nil {
				_handler.sk = string(skBytes)
			}
		}
	}

	if data != nil && string(data) != "" {
		err = json.Unmarshal(data, &_handler.data)
		if err != nil {
			return nil, err
		}
	}

	return &_handler, nil
}

func (h *handler) List(ctx context.Context) (result []byte, err error) {
	var (
		_ecs ecs.Interface
	)

	switch h.resource {
	case clouds.CloudResourceHost:
		switch h.handleType {
		case clouds.DescribeInstances:
			endpoint := fmt.Sprintf("ecs.%s.aliyuncs.com", h.region)
			_ecs, err = ecs.New(h.ak, h.sk, endpoint)
			if err != nil {
				return
			}

			if _, ok := h.data["RegionId"]; !ok {
				h.data["RegionId"] = h.region
			}

			result, err = _ecs.DescribeInstances(ctx, h.data)
		default:
			err = fmt.Errorf("handle type %s not support", h.handleType)
		}
	default:
		err = fmt.Errorf("resource type %s not support", h.resource)
	}
	return
}

func (h *handler) Get(ctx context.Context) (result []byte, err error) {
	return
}

func (h *handler) Post(ctx context.Context) (result []byte, err error) {
	return
}

func (h *handler) Put(ctx context.Context) (result []byte, err error) {
	return
}

func (h *handler) Delete(ctx context.Context) (result []byte, err error) {
	return
}
