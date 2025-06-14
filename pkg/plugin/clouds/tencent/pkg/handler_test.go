package pkg

import (
	"context"
	"fmt"
	"os"
	"testing"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Test_handler_Get(t *testing.T) {
	_ = os.Setenv("TENCENT_ACCESS_KEY_ID", "xxx")
	_ = os.Setenv("TENCENT_ACCESS_KEY_SECRET", "xxx")

	h, _ := NewHandler("Host", "ap-beijing", "DescribeInstances", []byte{})
	gotResult, _ := h.Get(context.Background())
	fmt.Println(string(gotResult))
}
