package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("example middle")
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
