package notify

import (
	"openiam/pkg/notify/commom"
	dingtalkNotification "openiam/pkg/notify/sender/dingtalk/notification"
	dingtalkWebhook "openiam/pkg/notify/sender/dingtalk/webhook"
	"openiam/pkg/notify/sender/email"
	larkNotification "openiam/pkg/notify/sender/lark/notification"
	larkWebhook "openiam/pkg/notify/sender/lark/webhook"
	wecomNotification "openiam/pkg/notify/sender/wecom/notification"
	wecomWebhook "openiam/pkg/notify/sender/wecom/webhook"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	Email(to, cc []string, title string, content *commom.Message) error
	DingTalkRobot(content map[string]interface{}, webhook []string) error
	DingTalkNotify(content map[string]interface{}, userIdList, deptIdList string, toAllUser bool) (result string, err error)
	WeComRobot(webhook []string, content map[string]interface{}) error
	WeComNotify(content map[string]interface{}) (result map[string]interface{}, err error)
	LarkRobot(webhook []string, content map[string]interface{}) error
	LarkNotify(mobiles []string, content map[string]interface{}) (result map[string]interface{}, err error)
}

type notify struct{}

func New() Interface {
	return &notify{}
}

// Email
// @Description: 封装邮件通知
// @param to 接收人
// @param cc 抄送人
// @param from 发送人
// @param title 标题
// @param content 内容
// @return err
func (n *notify) Email(to, cc []string, title string, content *commom.Message) (err error) {
	return email.Send(to, cc, title, content)
}

// DingTalkRobot
// @Description: 钉钉机器人通知
// @param content 消息内容
// @param webhook 机器人 webhook 地址
// @return err 错误信息
func (n *notify) DingTalkRobot(content map[string]interface{}, webhook []string) (err error) {
	return dingtalkWebhook.Send(webhook, content)
}

// DingTalkNotify
// @Description: 钉钉工作通知
// @param content 消息内容
// @param userIdList 接收人的用户userid列表，最大列表长度：20
// @param deptIdList 接收者的部门id列表，最大列表长度：20
// @param toAllUser 是否发送给企业全部用户
// @return err 错误信息
func (n *notify) DingTalkNotify(content map[string]interface{}, userIdList, deptIdList string, toAllUser bool) (result string, err error) {
	return dingtalkNotification.Send(userIdList, deptIdList, toAllUser, content)
}

// WeComRobot
// @Description: 企业微信机器人通知
// @param webhook 企业微信机器人 webhook 地址
// @param content 消息内容
// @return error 错误信息
func (n *notify) WeComRobot(webhook []string, content map[string]interface{}) error {
	return wecomWebhook.Send(webhook, content)
}

// LarkRobot
// @Description: 飞书机器人通知
// @param webhook 飞书机器人 webhook 地址
// @param content 消息内容
// @return error 错误信息
func (n *notify) LarkRobot(webhook []string, content map[string]interface{}) error {
	return larkWebhook.Send(webhook, content)
}

// LarkNotify
// @Description: 飞书机器人通知
// @param webhook 飞书机器人 webhook 地址
// @param content 消息内容
// @return error 错误信息
func (n *notify) LarkNotify(mobiles []string, content map[string]interface{}) (result map[string]interface{}, err error) {
	return larkNotification.Send(mobiles, content)
}

// WeComNotify
// @Description: 企业微信工作通知
// @param content 消息内容
// @return error 错误信息
func (n *notify) WeComNotify(content map[string]interface{}) (result map[string]interface{}, err error) {
	return wecomNotification.Send(content)
}
