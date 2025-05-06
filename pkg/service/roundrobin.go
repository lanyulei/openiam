package service

import "sync"

/*
  @Author : lanyulei
  @Desc :
*/

type RoundRobin struct {
	instances []string
	index     int
	mutex     sync.Mutex
}

func NewRoundRobin(instances []string) *RoundRobin {
	return &RoundRobin{
		instances: instances,
		index:     0,
	}
}

func (r *RoundRobin) Next() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if len(r.instances) == 0 {
		return ""
	}

	instance := r.instances[r.index]
	r.index = (r.index + 1) % len(r.instances)
	return instance
}
