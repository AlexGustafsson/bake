package semantics

import "github.com/AlexGustafsson/bake/ast"

type Builder struct {
	RootScope    *Scope
	CurrentScope *Scope
}

func CreateBuilder() *Builder {
	builder := &Builder{
		RootScope: CreateScope(nil),
	}
	builder.CurrentScope = builder.RootScope
	return builder
}

// Build builds a symbol table from a parse tree (or sub tree)
func Build(root ast.Node) *Scope {
	builder := CreateBuilder()
	builder.Build(root)
	return builder.RootScope
}

func (builder *Builder) Build(root ast.Node) {
	switch node := root.(type) {
	case *ast.SourceFile:
		for _, child := range node.Nodes {
			builder.Build(child)
		}
	case *ast.Block:
		for _, child := range node.Statements {
			builder.Build(child)
		}
	case *ast.VariableDeclaration:
		// If the variable is declared without an expression, we cannot infer it's type
		symbol := CreateSymbol(node.Identifier, TraitAny, node)
		if node.Expression != nil {
			// Use the trait of the expression
			builder.Build(node.Expression)
			if expressionSymbol, ok := builder.CurrentScope.SymbolTable.LookupNode(node.Expression); ok {
				symbol.Trait = expressionSymbol.Trait
			}
		}
		builder.CurrentScope.SymbolTable.Insert(symbol)
	case *ast.FunctionDeclaration:
		symbol := CreateSymbol(node.Identifier, TraitCallable, node)
		if node.Signature != nil {
			symbol.ArgumentCount = len(node.Signature.Arguments)
		}
		builder.CurrentScope.SymbolTable.Insert(symbol)

		builder.pushScope()
		if node.Signature != nil {
			builder.Build(node.Signature)
		}

		builder.Build(node.Block)
		builder.popScope()
	case *ast.Signature:
		for _, child := range node.Arguments {
			// TODO: Do we need actual nodes for arguments (identifiers)?
			symbol := CreateSymbol(child, TraitAny, node)
			builder.CurrentScope.SymbolTable.Insert(symbol)
		}
	}
}

func (builder *Builder) pushScope() {
	builder.CurrentScope = builder.CurrentScope.CreateScope()
}

func (builder *Builder) popScope() {
	builder.CurrentScope = builder.CurrentScope.ParentScope
}
