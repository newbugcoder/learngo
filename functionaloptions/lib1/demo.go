package lib1

import (
	"time"
)

// 基于闭包函数实现的功能选项模式
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

func NewServer(addr string, port int, options ...Option) *Server {
	svr := Server{
		Addr: addr,
		Port: port,
		Protocol: "tcp",
		Timeout: time.Second,
		MaxConns: 10,
	}

	for _, option := range options {
		option(&svr)
	}
	return &svr
}

type Option func(*Server)

func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}
func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func WithMaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}