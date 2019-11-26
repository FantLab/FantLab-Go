package docs

import (
	"reflect"
	"runtime"

	"golang.org/x/tools/go/packages"
)

func unptr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func loadPackage(dir string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, dir)

	if err != nil {
		return nil
	}
	return pkgs[0]
}

func getCallerFrame(i interface{}) *runtime.Frame {
	pc := reflect.ValueOf(i).Pointer()
	frames := runtime.CallersFrames([]uintptr{pc})

	if frames == nil {
		return nil
	}

	frame, _ := frames.Next()

	if frame.Entry == 0 {
		return nil
	}
	return &frame
}
