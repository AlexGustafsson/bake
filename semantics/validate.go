package semantics

import "go/ast"

// Validate validates a parse tree (or sub tree), returning a root scope if successfull
// or an error indicating semantic issues otherwise
func Validate(root ast.Node) (*Scope, error) {
	rootScope := CreateScope(nil)
	return rootScope, nil
}
