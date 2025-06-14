package types

import "encoding/json"

type Response struct {
	UniqueId string          `json:"unique_id"`
	Content  json.RawMessage `json:"content"`
}

type AccountType string

const (
	Common AccountType = "common" // 通用
)

type CloudName string

const (
	AliCloud     CloudName = "AliCloud"     // 阿里云
	TencentCloud CloudName = "TencentCloud" // 腾讯云
)

type CouldResourceType string

const (
	CouldResourceHost   CouldResourceType = "Host"   // 云主机
	CouldResourceDisk   CouldResourceType = "Disk"   // 云硬盘
	CouldResourceImages CouldResourceType = "Images" // 镜像
)
