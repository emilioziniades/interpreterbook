package evaluator

import (
	"github.com/emilioziniades/interpreterbook/ast"
	"github.com/emilioziniades/interpreterbook/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
