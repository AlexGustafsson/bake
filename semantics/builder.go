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
		symbol := CreateSymbol(node.Identifier, TraitAny, node)
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
	case *ast.IfStatement:
		builder.pushScope()
		builder.Build(node.PositiveBranch)
		builder.popScope()

		builder.pushScope()
		if node.NegativeBranch != nil {
			builder.Build(node.NegativeBranch)
		}
		builder.popScope()
	case *ast.RuleFunctionDeclaration:
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
	case *ast.RuleDeclaration:
		if node.Block != nil {
			builder.pushScope()
			builder.Build(node.Block)
			builder.popScope()
		}
	case *ast.AliasDeclaration:
		symbol := CreateSymbol(node.Identifier, TraitAlias, node)
		builder.CurrentScope.SymbolTable.Insert(symbol)
	case *ast.Signature:
		for _, child := range node.Arguments {
			symbol := CreateSymbol(child.Value, TraitAny, child)
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
