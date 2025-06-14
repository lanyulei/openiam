package pkg

import (
	"context"
	"fmt"
	"testing"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Test_handler_Get(t *testing.T) {
	h, _ := NewHandler("Host", "cn-beijing", "DescribeInstances", []byte{})
	gotResult, _ := h.Get(context.Background())
	fmt.Println(string(gotResult))
}
