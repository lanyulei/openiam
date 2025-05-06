package request

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"openiam/pkg/tools/respstatus"
	"regexp"

	"github.com/lanyulei/toolkit/logger"

	"github.com/gin-gonic/gin"
)

func Forward(c *gin.Context, instance, routePrefix, targetPrefix string) {
	target, _ := url.Parse(instance)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Director = func(req *http.Request) {
		// 使用正则表达式替换 URL 前缀
		re := regexp.MustCompile("^" + routePrefix)
		req.URL.Path = re.ReplaceAllString(req.URL.Path, targetPrefix)

		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host

		req.Header.Set("X-Real-IP", req.RemoteAddr)
	}

	// 捕获代理错误
	proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
		logger.Errorf("error occurred during proxying request: %v", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Header().Set("Content-Type", "application/json")
		_, _ = rw.Write([]byte(fmt.Sprintf(`{"code": "%d", "message": "%s, err: %s"}`, respstatus.RequestForwardError.Code, respstatus.RequestForwardError.Message, err.Error())))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
