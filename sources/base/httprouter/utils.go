package httprouter

import (
	"context"
	"net/http"
)

func GetValueFromContext(ctx context.Context, paramsKey interface{}, valueKey string) (value string, exists bool) {
	if values, ok := ctx.Value(paramsKey).(map[string]string); ok {
		if values != nil {
			value, exists = values[valueKey]
		}
	}
	return
}

var httpMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodPost:    true,
	http.MethodPut:     true,
	http.MethodPatch:   true,
	http.MethodDelete:  true,
	http.MethodConnect: true,
	http.MethodOptions: true,
	http.MethodTrace:   true,
}

func walkGroup(g *Group, mws []Middleware, fn func(mws []Middleware, endpoint *Endpoint)) {
	if g == nil {
		return
	}
	mws = append(mws, g.Middlewares...)
	for _, endpoint := range g.Endpoints {
		if endpoint != nil {
			fn(mws, endpoint)
		}
	}
	for _, sg := range g.Subgroups {
		walkGroup(sg, mws, fn)
	}
}

func chainHandler(endpoint http.Handler, mws ...Middleware) http.Handler {
	if len(mws) == 0 || endpoint == nil {
		return endpoint
	}
	h := mws[len(mws)-1](endpoint)
	for i := len(mws) - 2; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
