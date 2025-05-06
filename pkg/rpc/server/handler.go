package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"openiam/app/system/models"
	"openiam/pkg/jwtauth"
	"openiam/server/login"

	"github.com/lanyulei/toolkit/db"

	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

type RPCServer struct{}

// CheckToken 验证函数
func (r *RPCServer) CheckToken(token string) bool {
	_, err := jwtauth.ParseToken(token, viper.GetString("jwt.rpc.secret"))
	if err != nil {
		return false
	}
	return true
}

func (r *RPCServer) GetUserByUsername(args *map[string]interface{}, reply *map[string]interface{}) error {
	var (
		err error
	)

	if token, ok := (*args)["token"]; ok {
		status := r.CheckToken(token.(string))
		if !status {
			return errors.New("token invalid")
		}
	}

	if username, ok := (*args)["username"]; ok {
		var (
			userInfo models.UserRequest
			user     []byte
		)
		err = db.Orm().Model(&models.UserRequest{}).Where("username = ?", username).First(&userInfo).Error
		if err != nil {
			err = fmt.Errorf("failed to get user info, error: %s", err.Error())
			return err
		}

		user, err = json.Marshal(userInfo)
		if err != nil {
			err = fmt.Errorf("failed to marshal user info, error: %s", err.Error())
			return err
		}

		err = json.Unmarshal(user, reply)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal user info, error: %s", err.Error())
			return err
		}
	}

	return nil
}

func (r *RPCServer) Authenticate(args *map[string]interface{}, reply *bool) (err error) {
	if _, ok := (*args)["username"]; !ok {
		err = fmt.Errorf("username is empty")
		return
	}

	if _, ok := (*args)["password"]; !ok {
		err = fmt.Errorf("password is empty")
		return
	}

	if (*args)["username"].(string) == "" || (*args)["password"].(string) == "" {
		err = fmt.Errorf("username or password is empty")
		return
	}

	_, err = login.LoginHandler((*args)["username"].(string), (*args)["password"].(string), (*args)["isLdap"].(bool))
	if err != nil {
		err = fmt.Errorf("login error: %v", err)
		return
	}

	*reply = true

	return
}
