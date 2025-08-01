package permission

import (
	"fmt"
	"openops/pkg/respstatus"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/lanyulei/toolkit/redis"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

var enforcer *casbin.SyncedEnforcer

func Setup() *casbin.SyncedEnforcer {
	var (
		err     error
		adapter *gormAdapter.Adapter
		m       model.Model
	)
	adapter, err = gormAdapter.NewAdapterByDBWithCustomTable(db.Orm(), nil, viper.GetString("casbin.tableName"))
	if err != nil {
		logger.Fatalf("failed to create casbin gorm adapter, error：%v", err)
	}

	m, err = model.NewModelFromString(rbacModel)
	if err != nil {
		logger.Fatalf("failed to generate casbin model, error：%v", err)
	}

	enforcer, err = casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		logger.Fatalf("failed to create casbin enforcer, error：%v", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Fatalf("failed to load policy from database, error：%v", err)
	}

	// 通过 redis 发布订阅机制，实现权限策略的实时更新
	go func(enforcer *casbin.SyncedEnforcer) {
		sub := redis.Rc().Subscribe(syncLabel)
		for msg := range sub.Channel() {
			if msg.Payload == "true" {
				err = enforcer.LoadPolicy()
				if err != nil {
					logger.Fatalf("failed to load policy from database, error：%v", err)
				}
			}
		}
	}(enforcer)

	// 定时同步策略
	if viper.GetBool("casbin.isTiming") {
		// 间隔多长时间同步一次权限策略，单位：秒
		enforcer.StartAutoLoadPolicy(time.Second * time.Duration(viper.GetInt("casbin.intervalTime")))
	}

	return enforcer
}

func CheckPermMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ok bool
		)

		//获取资源
		obj := c.Request.URL.Path
		//获取方法
		act := c.Request.Method
		//获取实体
		sub := c.GetString("username")

		isAdmin := c.GetBool("isAdmin")
		if isAdmin {
			c.Next()
		} else {
			//判断策略中是否存在
			if ok, _ = enforcer.Enforce(sub, obj, act); ok {
				c.Next()
			} else {
				response.Error(c, fmt.Errorf("暂无权限通过 %s 方式访问 %s ", act, obj), respstatus.NoPermissionError)
				c.Abort()
			}
		}
	}
}

func Enforcer() *casbin.SyncedEnforcer {
	return enforcer
}

// Sync 兼容多服务权限缓存不同步的问题
func Sync() (err error) {
	err = redis.Rc().Publish(syncLabel, "true")
	if err != nil {
		return
	}
	return
}
