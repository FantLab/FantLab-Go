package docs

import (
	"fantlab/protobuf"
	"go/parser"
	"go/token"
	"reflect"
	"runtime"
)

type handlerInfo struct {
	comment string
	file    string
	line    int
}

func getHandlerInfo(handler protobuf.HandlerFunc) (result handlerInfo) {
	frame := getCallerFrame(handler)

	if frame == nil {
		return
	}

	result = handlerInfo{
		comment: getFuncComment(frame.File, frame.Line),
		file:    frame.File,
		line:    frame.Line,
	}

	return
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

func getFuncComment(file string, line int) string {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return ""
	}

	if len(astFile.Comments) == 0 {
		return ""
	}

	for _, cmt := range astFile.Comments {
		if fset.Position(cmt.End()).Line+1 == line {
			return cmt.Text()
		}
	}

	return ""
}
