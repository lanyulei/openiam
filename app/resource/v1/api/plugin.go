package api

import (
	"openops/pkg/respstatus"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func PluginList(c *gin.Context) {
	var (
		err    error
		result []string
		files  []os.DirEntry
	)

	// 获取目录的文件列表
	dirPath := viper.GetString("plugin.path")
	files, err = os.ReadDir(dirPath)
	if err != nil {
		response.Error(c, err, respstatus.GetPluginListError)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}

	response.OK(c, result, "")
}
