package main

import "net/http"

// 大写在包之外能够被别人使用
// 什么是接口？  一组行为的抽象
type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (s *sdkHttpServer) start(address string) error {
	//TODO implement me
	panic("implement me")
}

func main() {

}

// 接口就是对实际事务的抽象
