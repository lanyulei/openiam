package route

import (
	"strings"
)

/*
  @Author : lanyulei
  @Desc :
*/

func MatchRouter(routes []map[string]interface{}, path, method string) (r map[string]interface{}, ok bool) {
	// 分割路径段
	pathSegments := strings.Split(path, "/")

	// 遍历路由列表
	for _, route := range routes {
		routePath := ""
		if method == "" {
			routePath = route["path"].(string) // 菜单
		} else {
			routePath = route["url"].(string) // 接口
		}

		// 分割路由段
		routeSegments := strings.Split(routePath, "/")

		// 如果路由段数和路径段数不同，则跳过
		if len(routeSegments) != len(pathSegments) {
			continue
		}

		// 比较路由段和路径段
		for i, segment := range routeSegments {
			// 如果路由段和路径段不同，且路由段不是一个路径参数，则跳过
			if segment != pathSegments[i] && !strings.HasPrefix(segment, ":") {
				break
			}

			// 如果已经比较到最后一个路由段，则说明路由匹配
			if i == len(routeSegments)-1 {
				if method == "" { // 菜单
					r = route
					ok = true
				} else { // 接口
					if route["method"].(string) == method {
						r = route
						ok = true
					}
				}
				return
			}
		}
	}
	return
}
