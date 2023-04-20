package configobject

import (
	"time"
)

//type Server struct {
//	Addr     string
//	Port     int
//	Protocol string
//	Timeout  time.Duration
//	MaxConns int
//}

//func NewServer(addr string, port int) *Server {
//	return &Server{
//		Addr: addr,
//		Port: port,
//		Protocol: "tcp",
//		Timeout: time.Second,
//		MaxConns: 10,
//	}
//}
//
//func NewServerExt(addr string, port int, protocol string, timeout time.Duration, maxconns int) *Server {
//	return &Server{
//		Addr: addr,
//		Port: port,
//		Protocol: protocol,
//		Timeout: timeout,
//		MaxConns: maxconns,
//	}
//}

// 配置对象方案

type Server struct {
	Addr     string
	Port     int
	Config *Config
}

type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

func NewServer(addr string, port int, config *Config) *Server {
	if config != nil && config.Protocol == "" {
		config.Protocol = "tcp"
	}

	if config != nil && config.Timeout  == 0 {
		config.Timeout = time.Second
	}

	if config != nil && config.MaxConns == 0 {
		config.MaxConns = 10
	}

	return &Server{
		Addr: addr,
		Port: port,
		Config: config,
	}
}