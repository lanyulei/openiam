package server

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/lanyulei/toolkit/logger"

	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Start() {
	var (
		err  error
		conn net.Conn
	)

	// 创建Authenticator实例
	rpcServer := new(RPCServer)

	err = rpc.Register(rpcServer)
	if err != nil {
		logger.Fatalf("rpc register failed, err: %s", err.Error())
	}

	// 创建TCP监听
	listener, err := net.Listen("tcp", viper.GetString("rpc.server"))
	if err != nil {
		logger.Fatalf("Error starting listener: %s", err.Error())
	}
	logger.Info("RPC server started, waiting for connections...")

	// 接受连接并为每个连接创建一个新的goroutine进行处理
	for {
		conn, err = listener.Accept()
		if err != nil {
			logger.Errorf("Error accepting connection:", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
