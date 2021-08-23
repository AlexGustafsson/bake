package semantics

import (
	"github.com/AlexGustafsson/bake/ast"
)

type Builder struct {
	RootScope    *Scope
	CurrentScope *Scope
	errors       []error
}

func CreateBuilder() *Builder {
	builder := &Builder{
		RootScope: CreateScope(nil),
		errors:    make([]error, 0),
	}
	builder.CurrentScope = builder.RootScope
	return builder
}

// Build builds a symbol table from a parse tree (or sub tree)
func Build(root ast.Node) (*Scope, []error) {
	builder := CreateBuilder()
	builder.Build(root)
	return builder.RootScope, builder.errors
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
		builder.insertInScope(symbol)
	case *ast.FunctionDeclaration:
		symbol := CreateSymbol(node.Identifier, TraitCallable, node)
		if node.Signature != nil {
			symbol.ArgumentCount = len(node.Signature.Arguments)
		}
		builder.insertInScope(symbol)

		builder.pushScope(node)
		if node.Signature != nil {
			builder.Build(node.Signature)
		}

		builder.Build(node.Block)
		builder.popScope()
	case *ast.IfStatement:
		builder.pushScope(node)
		builder.Build(node.PositiveBranch)
		builder.popScope()

		builder.pushScope(node)
		if node.NegativeBranch != nil {
			builder.Build(node.NegativeBranch)
		}
		builder.popScope()
	case *ast.RuleFunctionDeclaration:
		symbol := CreateSymbol(node.Identifier, TraitCallable, node)
		if node.Signature != nil {
			symbol.ArgumentCount = len(node.Signature.Arguments)
		}
		builder.insertInScope(symbol)

		builder.pushScope(node)
		if node.Signature != nil {
			builder.Build(node.Signature)
		}

		builder.Build(node.Block)
		builder.popScope()
	case *ast.RuleDeclaration:
		if node.Block != nil {
			builder.pushScope(node)
			builder.Build(node.Block)
			builder.popScope()
		}
	case *ast.AliasDeclaration:
		symbol := CreateSymbol(node.Identifier, TraitAlias, node)
		builder.insertInScope(symbol)
	case *ast.Signature:
		for _, child := range node.Arguments {
			symbol := CreateSymbol(child.Value, TraitAny, child)
			builder.insertInScope(symbol)
		}
	}
}

func (builder *Builder) errorf(node ast.Node, format string, arguments ...interface{}) {
	r := ast.CreateRange(node.Start(), node.End())
	builder.errors = append(builder.errors, ast.CreateTreeError(&r, format, arguments...))
}

func (builder *Builder) insertInScope(symbol *Symbol) {
	if previous, ok := builder.CurrentScope.SymbolTable.LookupByName(symbol.Name); ok {
		builder.errorf(symbol.Node, "'%s' already declared on line %d", symbol.Name, previous.Node.Start().Line+1)
	} else {
		builder.CurrentScope.SymbolTable.Insert(symbol)
	}
}

func (builder *Builder) pushScope(node ast.Node) {
	builder.CurrentScope = builder.CurrentScope.CreateScope(node)
}

func (builder *Builder) popScope() {
	builder.CurrentScope = builder.CurrentScope.ParentScope
}
