package webk

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义请求处理函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现http.Handler接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 构造webk.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path

	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET 定义添加GET请求的方法
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST 定义添加POST请求的方法
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
