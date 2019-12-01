package analysis

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/packages"
)

type commentsTable = map[string]map[string]string

func makeModelCommentsTable(pkg *packages.Package, isValidField func(f *ast.Field) bool) commentsTable {
	table := make(commentsTable)

	for _, f := range pkg.Syntax {
		ast.Inspect(f, func(node ast.Node) bool {
			if typeSpec, ok := node.(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					subtable := make(map[string]string)
					for _, f := range structType.Fields.List {
						if isValidField(f) {
							comment := strings.TrimRight(f.Doc.Text(), "\n")
							if "" != comment {
								subtable[f.Names[0].Name] = comment
							}
						}
					}
					if len(subtable) > 0 {
						table[typeSpec.Name.Name] = subtable
					}
					return false
				}
			}
			return true
		})
	}
	return table
}
