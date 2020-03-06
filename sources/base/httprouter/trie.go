package httprouter

import (
	"net/http"
	"strings"
)

type pathSegment struct {
	handler http.Handler
	keys    map[string]struct{}
	dynamic *pathSegment
	static  map[string]*pathSegment
}

func (s *pathSegment) insertPathHandler(path []string, handler http.Handler) {
	if len(path) == 0 {
		s.handler = handler

		return
	}

	wildcard, name := parseSegment(path[0])

	var child *pathSegment

	if wildcard {
		child = s.dynamic

		if child == nil {
			child = new(pathSegment)

			s.dynamic = child
		}

		if child.keys == nil {
			child.keys = make(map[string]struct{})
		}

		child.keys[name] = struct{}{}
	} else {
		if s.static != nil {
			child = s.static[name]
		}

		if child == nil {
			child = new(pathSegment)

			if s.static == nil {
				s.static = make(map[string]*pathSegment)
			}

			s.static[name] = child
		}
	}

	child.insertPathHandler(path[1:], handler)
}

func (s *pathSegment) handlerForPath(path []string, saveParam func(key, value string)) http.Handler {
	if len(path) == 0 {
		return s.handler
	}

	name, path := path[0], path[1:]

	if child := s.static[name]; child != nil {
		if handler := child.handlerForPath(path, saveParam); handler != nil {
			return handler
		}
	}

	if s.dynamic != nil {
		if handler := s.dynamic.handlerForPath(path, saveParam); handler != nil {
			for key := range s.dynamic.keys {
				saveParam(key, name)
			}
			return handler
		}
	}

	return nil
}

func parseSegment(s string) (bool, string) {
	if s[0] == ':' {
		return true, s[1:]
	}
	return false, s
}

type pathTrie struct {
	maxDepth         int
	prefix           string
	segmentValidator func(string) bool
	root             *pathSegment
}

func (t *pathTrie) insertPathHandler(path string, handler http.Handler) bool {
	if handler == nil {
		return false
	}

	segments := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	if t.segmentValidator != nil {
		for _, segment := range segments {
			wildcard, name := parseSegment(segment)
			if wildcard {
				continue
			}
			if !t.segmentValidator(name) {
				return false
			}
		}
	}

	t.root.insertPathHandler(segments, handler)

	if len(segments) > t.maxDepth {
		t.maxDepth = len(segments)
	}

	return true
}

func (t *pathTrie) handlerForPath(path string) (http.Handler, map[string]string) {
	segments := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	if t.prefix != "" {
		if len(segments) == 0 || t.prefix != segments[0] {
			return nil, nil
		}
		segments = segments[1:]
	}

	if len(segments) > t.maxDepth {
		return nil, nil
	}

	var params map[string]string

	handler := t.root.handlerForPath(segments, func(key, value string) {
		if params == nil {
			params = make(map[string]string)
		}
		params[key] = value
	})

	return handler, params
}

func newPathTrie(prefix string, segmentValidator func(string) bool) *pathTrie {
	return &pathTrie{
		prefix:           prefix,
		segmentValidator: segmentValidator,
		root:             &pathSegment{},
	}
}
