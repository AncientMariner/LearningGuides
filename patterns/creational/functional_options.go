package creational

import (
	"fmt"
	"time"
)

// Functional Options: preferred over Builder for most cases.
// Add new options without breaking existing callers.

type HTTPServer struct {
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxConns     int
}

type Option func(*HTTPServer)

func WithHTTPReadTimeout(d time.Duration) Option {
	return func(s *HTTPServer) { s.readTimeout = d }
}

func WithHTTPWriteTimeout(d time.Duration) Option {
	return func(s *HTTPServer) { s.writeTimeout = d }
}

func WithHTTPMaxConns(n int) Option {
	return func(s *HTTPServer) { s.maxConns = n }
}

func NewHTTPServer(addr string, opts ...Option) *HTTPServer {
	s := &HTTPServer{
		addr:         addr,
		readTimeout:  5 * time.Second,
		writeTimeout: 10 * time.Second,
		maxConns:     100,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func ExampleFunctionalOptions() {
	srv := NewHTTPServer(":9090",
		WithHTTPMaxConns(500),
		WithHTTPReadTimeout(30*time.Second),
	)
	fmt.Println(srv.addr)
	fmt.Println(srv.maxConns)
	// Output:
	// :9090
	// 500
}
