package routing

import (
	"fantlab/base/protobuf"
	"net/http"
)

type Endpoint struct {
	method  string
	path    string
	info    string
	handler protobuf.HandlerFunc
}

func (e *Endpoint) Method() string {
	return e.method
}

func (e *Endpoint) Path() string {
	return e.path
}

func (e *Endpoint) Info() string {
	return e.info
}

func (e *Endpoint) Handler() protobuf.HandlerFunc {
	return e.handler
}

// *******************************************************

type Group struct {
	info        string
	middlewares []func(http.Handler) http.Handler
	endpoints   []Endpoint
	subgroups   []*Group
}

func (g *Group) Walk(fn func(g *Group)) {
	fn(g)

	for _, sg := range g.subgroups {
		sg.Walk(fn)
	}
}

func (g *Group) Info() string {
	return g.info
}

func (g *Group) Middlewares() []func(http.Handler) http.Handler {
	return g.middlewares
}

func (g *Group) Endpoints() []Endpoint {
	return g.endpoints
}

func (g *Group) Subgroups() []*Group {
	return g.subgroups
}

func (g *Group) middleware(fn func(http.Handler) http.Handler) {
	g.middlewares = append(g.middlewares, fn)
}

func (g *Group) endpoint(method string, path string, handler protobuf.HandlerFunc, info string) {
	g.endpoints = append(g.endpoints, Endpoint{
		method:  method,
		path:    path,
		info:    info,
		handler: handler,
	})
}

func (g *Group) subgroup(info string, fn func(g *Group)) {
	sg := new(Group)
	sg.info = info
	fn(sg)
	g.subgroups = append(g.subgroups, sg)
}
