package structural

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware: wraps an http.Handler to add cross-cutting concerns.
// Each middleware calls the next handler, forming a chain.
// This is the Decorator pattern applied to http.Handler.

type Middleware func(http.Handler) http.Handler

// Chain applies middlewares right-to-left so the first listed runs first.
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Auth(token string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") != "Bearer "+token {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Generic function middleware — useful outside HTTP (e.g. CLI commands, gRPC).

type HandlerFunc[Req, Resp any] func(Req) (Resp, error)

type FuncMiddleware[Req, Resp any] func(HandlerFunc[Req, Resp]) HandlerFunc[Req, Resp]

func ChainFunc[Req, Resp any](h HandlerFunc[Req, Resp], middlewares ...FuncMiddleware[Req, Resp]) HandlerFunc[Req, Resp] {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func LoggingFunc[Req, Resp any](label string) FuncMiddleware[Req, Resp] {
	return func(next HandlerFunc[Req, Resp]) HandlerFunc[Req, Resp] {
		return func(req Req) (Resp, error) {
			fmt.Printf("[%s] calling\n", label)
			resp, err := next(req)
			fmt.Printf("[%s] done\n", label)
			return resp, err
		}
	}
}

func ExampleMiddleware() {
	greet := HandlerFunc[string, string](func(name string) (string, error) {
		return "hello, " + name, nil
	})

	wrapped := ChainFunc(greet,
		LoggingFunc[string, string]("greet"),
	)

	result, _ := wrapped("Alice")
	fmt.Println(result)
	// Output:
	// [greet] calling
	// [greet] done
	// hello, Alice
}
