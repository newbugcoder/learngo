package lib2

import (
	"time"
)

// 基于接口函数实现的功能选项模式
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
		option.apply(&svr)
	}
	return &svr
}

type Option interface {
	apply(*Server)
}

type protocolOption string

func (c protocolOption) apply(server *Server) {
	server.Protocol = string(c)
}

func WithProtocol(protocol string) Option {
	return protocolOption(protocol)
}

type timeoutOption time.Duration

func (t timeoutOption) apply(server *Server) {
	server.Timeout = time.Duration(t)
}

func WithTimeout(timeout time.Duration) Option {
	return timeoutOption(timeout)
}

type maxConnsOption int

func (m maxConnsOption) apply(server *Server) {
	server.MaxConns = int(m)
}

func WithMaxConns(maxconns int) Option {
	return maxConnsOption(maxconns)
}