package evaluator

import (
	"toy-monkey/ast"
	"toy-monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
