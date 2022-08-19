package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type backEnd struct {
	proxy       *httputil.ReverseProxy
	containerID string
}

type ServiceRegistry struct {
	BackendStore atomic.Value
}

func (s *ServiceRegistry) Init() {
	s.BackendStore.Store([]backEnd{})
}

func (s *ServiceRegistry) Add(containerID, addr string) {
	URL, _ := url.Parse(addr)
	s.BackendStore.Swap(append(s.GetBackends(), backEnd{
		proxy:       httputil.NewSingleHostReverseProxy(URL),
		containerID: containerID,
	}))
}

func (s *ServiceRegistry) GetByContainerID(containerID string) (backEnd, bool) {
	for _, b := range s.GetBackends() {
		if b.containerID == containerID {
			return b, true
		}
	}
	return backEnd{}, false
}

func (s *ServiceRegistry) GetByIndex(index int) backEnd {
	return s.GetBackends()[index]
}

func (s *ServiceRegistry) RemoveByContainerID(containerID string) {
	var backends []backEnd
	for _, b := range s.GetBackends() {
		if b.containerID == containerID {
			continue
		}
		backends = append(backends, b)
	}
	s.BackendStore.Store(backends)
}

func (s *ServiceRegistry) RemoveAll() {
	s.BackendStore.Store([]backEnd{})
}

func (s *ServiceRegistry) Len() int {
	return len(s.GetBackends())
}

func (s *ServiceRegistry) List() {
	backends := s.GetBackends()
	for i := range backends {
		fmt.Println(backends[i].containerID)
	}
}

func (s *ServiceRegistry) GetBackends() []backEnd {
	return s.BackendStore.Load().([]backEnd)
}
