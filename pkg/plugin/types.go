package plugin

import "encoding/json"

/*
  @Author : lanyulei
  @Desc :
*/

type Response struct {
	UniqueId string          `json:"unique_id"`
	Content  json.RawMessage `json:"content"`
}
