package webhook

import (
	"github.com/lanyulei/toolkit/logger"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send lark webhook
// @param url webhook address
// @param msg message content
// @return err
func Send(urlList []string, msg map[string]interface{}) (err error) {
	if len(urlList) > 0 {
		// request webhook
		go func(urlList []string, msg map[string]interface{}) {
			for _, url := range urlList {
				err = gout.POST(url).
					SetHeader(gout.H{
						"Content-Type": "application/json",
					}).
					SetJSON(msg).
					Do()
				if err != nil {
					logger.Errorf("lark webhook send error: %v", err)
				}
			}
		}(urlList, msg)
	}
	return
}
