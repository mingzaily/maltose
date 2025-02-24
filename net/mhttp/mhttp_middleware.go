package mhttp

import (
	"github.com/gin-gonic/gin"
)

// MiddlewareFunc 定义中间件函数类型
type MiddlewareFunc func(*Request)

// Use 添加全局中间件
func (s *Server) Use(middleware ...MiddlewareFunc) {
	// 转换为 gin 的中间件格式
	for _, m := range middleware {
		handler := func(c *gin.Context) {
			r := RequestFromCtx(c.Request.Context())
			if r == nil {
				r = &Request{Context: c, Server: s}
				c.Request = c.Request.WithContext(WithRequest(c.Request.Context(), r))
			}
			m(r)
		}
		s.middleware = append(s.middleware, handler)
		s.Engine.Use(handler)
	}
}
