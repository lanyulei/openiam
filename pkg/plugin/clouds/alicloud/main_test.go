package main

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

func TestAliCloud_Get(t *testing.T) {
	_ = os.Setenv("ALIYUN_ACCESS_KEY_ID", "xxx")
	_ = os.Setenv("ALIYUN_ACCESS_KEY_SECRET", "xxx")

	a := AliCloud{}
	result, err := a.Get(context.Background(), "Host", "cn-hangzhou", "DescribeInstances", []byte{})
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
	fmt.Printf("Result: %v\n", result)
}
