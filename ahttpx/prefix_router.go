package ahttpx

import (
	"net/http"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 前缀优先匹配路由模块
// 仿照 net/http 的路由匹配模式
// 如果注册路由的最后一个字符是 / 则认为是一个前缀匹配路由

// 使用
// 将：server := rest.MustNewServer(c.RestConf)
// 改为：server := rest.MustNewServer(c.RestConf, rest.WithRouter(ahttpx.NewPrefixPriorityRouter(router.NewRouter())))

type prefixRoute struct {
	h       http.Handler
	pattern string
}

type prefixFirstRouter struct {
	httpx.Router

	prefixRouters []prefixRoute
}

func NewPrefixPriorityRouter(origin httpx.Router) httpx.Router {
	return &prefixFirstRouter{Router: origin}
}

func (r *prefixFirstRouter) Handle(method, path string, handler http.Handler) error {
	if path[len(path)-1] == '/' {
		e := prefixRoute{h: handler, pattern: path}
		r.prefixRouters = appendSorted(r.prefixRouters, e)
		return nil
	}

	return r.Router.Handle(method, path, handler)
}

func appendSorted(es []prefixRoute, e prefixRoute) []prefixRoute {
	n := len(es)
	i := sort.Search(n, func(i int) bool {
		return len(es[i].pattern) < len(e.pattern)
	})
	if i == n {
		return append(es, e)
	}
	es = append(es, prefixRoute{})
	copy(es[i+1:], es[i:])
	es[i] = e
	return es
}

func (r *prefixFirstRouter) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	for _, pr := range r.prefixRouters {
		if strings.HasPrefix(req.URL.Path, pr.pattern) {
			pr.h.ServeHTTP(writer, req)
			return
		}
	}

	r.Router.ServeHTTP(writer, req)
}
