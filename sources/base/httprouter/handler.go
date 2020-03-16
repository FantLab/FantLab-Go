package httprouter

import (
	"context"
	"net/http"
)

type httpRouter struct {
	tree             map[string]*pathTrie
	notFoundHandler  http.Handler
	contextParamsKey interface{}
}

func (hr *httpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	trie := hr.tree[r.Method]

	if trie == nil {
		hr.notFoundHandler.ServeHTTP(w, r)
		return
	}

	handler, params := trie.handlerForPath(r.URL.Path)

	if handler == nil {
		hr.notFoundHandler.ServeHTTP(w, r)
		return
	}

	if len(params) > 0 {
		r = r.WithContext(context.WithValue(r.Context(), hr.contextParamsKey, params))
	}

	handler.ServeHTTP(w, r)
}

type Config struct {
	RootGroup               *Group
	NotFoundHandler         http.Handler
	RequestContextParamsKey interface{}
	CommonPrefix            string
	PathSegmentValidator    func(string) bool
	GlobalMiddlewares       []Middleware
}

func NewRouter(cfg *Config) (http.Handler, []*Endpoint) {
	if cfg == nil || cfg.NotFoundHandler == nil || cfg.RootGroup == nil {
		return nil, nil
	}

	router := &httpRouter{
		tree:             make(map[string]*pathTrie),
		notFoundHandler:  cfg.NotFoundHandler,
		contextParamsKey: cfg.RequestContextParamsKey,
	}

	var badEndpoints []*Endpoint

	walkGroup(cfg.RootGroup, cfg.GlobalMiddlewares, func(mws []Middleware, e *Endpoint) {
		if httpMethods[e.Method] {
			tree := router.tree[e.Method]

			if tree == nil {
				tree = newPathTrie(cfg.CommonPrefix, cfg.PathSegmentValidator)

				router.tree[e.Method] = tree
			}

			if tree.insertPathHandler(e.Path, chainHandler(e.Handler, mws...)) {
				return
			}
		}

		badEndpoints = append(badEndpoints, e)
	})

	return router, badEndpoints
}
