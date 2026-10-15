package creational

import (
	"errors"
	"fmt"
	"time"
)

// Builder pattern: use when construction requires validation across multiple fields.
// Prefer functional options (see functional_options.go) for simpler cases.

type Server struct {
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxConns     int
}

type ServerBuilder struct {
	addr         string
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxConns     int
}

func NewServerBuilder(addr string) *ServerBuilder {
	return &ServerBuilder{
		addr:         addr,
		readTimeout:  5 * time.Second,
		writeTimeout: 10 * time.Second,
		maxConns:     100,
	}
}

func (b *ServerBuilder) ReadTimeout(d time.Duration) *ServerBuilder {
	b.readTimeout = d
	return b
}

func (b *ServerBuilder) WriteTimeout(d time.Duration) *ServerBuilder {
	b.writeTimeout = d
	return b
}

func (b *ServerBuilder) MaxConns(n int) *ServerBuilder {
	b.maxConns = n
	return b
}

func (b *ServerBuilder) Build() (*Server, error) {
	if b.addr == "" {
		return nil, errors.New("addr is required")
	}
	if b.readTimeout > b.writeTimeout {
		return nil, errors.New("readTimeout must not exceed writeTimeout")
	}
	return &Server{
		addr:         b.addr,
		readTimeout:  b.readTimeout,
		writeTimeout: b.writeTimeout,
		maxConns:     b.maxConns,
	}, nil
}

func ExampleBuilder() {
	srv, err := NewServerBuilder(":8080").
		ReadTimeout(5 * time.Second).
		WriteTimeout(10 * time.Second).
		MaxConns(200).
		Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(srv.addr)
	fmt.Println(srv.maxConns)
	// Output:
	// :8080
	// 200
}
