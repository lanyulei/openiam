package apis

import (
	"encoding/base64"
	"errors"
	"fmt"
	"openiam/pkg/tools/respstatus"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Upload
// @Description: 上传文件
func Upload(c *gin.Context) {
	var (
		urlPrefix    string
		tag          string
		fileType     string
		saveFilePath string
		err          error
		protocol     = "http"
		requestHost  string
	)

	tag, _ = c.GetPostForm("type")
	fileType = c.DefaultQuery("file_type", "")
	if fileType != "images" && fileType != "files" {
		response.Error(c, fmt.Errorf("上传接口目前，仅支持图片上传和文件上传"), respstatus.InvalidParameterError)
		return
	}

	if strings.HasPrefix(c.Request.Header.Get("Origin"), "https") {
		protocol = "https"
	}

	requestHostList := strings.Split(c.Request.Host, ":")
	if len(requestHostList) > 1 && requestHostList[1] == "80" {
		requestHost = requestHostList[0]
	} else {
		requestHost = c.Request.Host
	}

	if viper.GetBool("upload.isGetHost") {
		urlPrefix = fmt.Sprintf("%s://%s/", protocol, requestHost)
	} else {
		urlPrefix = fmt.Sprintf("%s://%s", protocol, viper.GetString("upload.url"))
		if !strings.HasSuffix(viper.GetString("upload.url"), "/") {
			urlPrefix = urlPrefix + "/"
		}
	}

	if tag == "" {
		tag = "1"
	}

	saveFilePath = "static/uploadfile/" + fileType + "/"
	_, err = os.Stat(saveFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(saveFilePath, 0755)
		if err != nil {
			response.Error(c, err, respstatus.UploadFileError)
			return
		}
	}

	guid := strings.ReplaceAll(uuid.New().String(), "-", "")

	switch tag {
	case "1": // 单图
		files, err := c.FormFile("file")
		if err != nil {
			response.Error(c, err, respstatus.UploadFileError)
			return
		}
		// 上传文件至指定目录
		singleFile := saveFilePath + guid + "-" + files.Filename
		_ = c.SaveUploadedFile(files, singleFile)
		response.OK(c, urlPrefix+singleFile, "上传成功")
		return
	case "2": // 多图
		files := c.Request.MultipartForm.File["file"]
		multipartFile := make([]string, len(files))
		for _, f := range files {
			guid = strings.ReplaceAll(uuid.New().String(), "-", "")
			multipartFileName := saveFilePath + guid + "-" + f.Filename
			_ = c.SaveUploadedFile(f, multipartFileName)
			multipartFile = append(multipartFile, urlPrefix+multipartFileName)
		}
		response.OK(c, multipartFile, "上传成功")
		return
	case "3": // base64
		files, _ := c.GetPostForm("file")
		ddd, _ := base64.StdEncoding.DecodeString(files)
		_ = os.WriteFile(saveFilePath+guid+".jpg", ddd, 0666)
		response.OK(c, urlPrefix+saveFilePath+guid+".jpg", "上传成功")
	default:
		response.Error(c, errors.New(""), respstatus.UploadFileError)
		return
	}
}
