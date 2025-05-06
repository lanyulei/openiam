package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/lanyulei/toolkit/redis"

	"github.com/lanyulei/toolkit/logger"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Discover
// @Description: 服务发现
func Discover(serviceName string) (serviceAddrs []string, err error) {
	keyPattern := Prefix + serviceName + ":*"

	keys, err := redis.Rc().Keys(keyPattern)
	if err != nil {
		logger.Fatalf("discover service failed: %v", err.Error())
	}

	serviceAddrs = make([]string, 0, len(keys))
	for _, key := range keys {
		addr, err := redis.Rc().Get(key)
		if err == nil {
			serviceAddrs = append(serviceAddrs, addr)
		}
	}
	return
}

func isInstanceHealthy(addr string, timeout time.Duration) bool {
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(fmt.Sprintf("http://%s/health", addr))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func DiscoverHealthyService(serviceName string) ([]string, error) {
	var (
		wg sync.WaitGroup
	)

	serviceAddrs, err := Discover(serviceName)
	if err != nil {
		return nil, err
	}

	healthyServiceAddrs := make([]string, 0, len(serviceAddrs))
	for _, addr := range serviceAddrs {
		go func(wg *sync.WaitGroup, addr string, healthyServiceAddrs []string) {
			defer wg.Done()

			if isInstanceHealthy(addr, 1*time.Second) {
				healthyServiceAddrs = append(healthyServiceAddrs, addr)
			}
		}(&wg, addr, healthyServiceAddrs)
	}
	wg.Wait()

	return healthyServiceAddrs, nil
}
