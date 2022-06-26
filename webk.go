package webk

import (
	"net/http"
)

// HandlerFunc 定义请求处理函数类型
type HandlerFunc func(*Context)

// Engine 实现http.Handler接口
type Engine struct {
	router *router
}

// New 构造webk.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
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
