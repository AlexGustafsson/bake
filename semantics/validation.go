package semantics

import (
	"strings"

	"github.com/AlexGustafsson/bake/ast"
)

type Validator struct {
	RootNode     ast.Node
	RootScope    *Scope
	CurrentScope *Scope
	errors       []error
}

func CreateValidator(rootNode ast.Node, rootScope *Scope) *Validator {
	validator := &Validator{
		RootNode:  rootNode,
		RootScope: rootScope,
		errors:    make([]error, 0),
	}
	validator.CurrentScope = validator.RootScope
	return validator
}

func Validate(rootNode ast.Node, rootScope *Scope) []error {
	validator := CreateValidator(rootNode, rootScope)
	validator.Validate(rootNode)
	return validator.errors
}

func (validator *Validator) Validate(root ast.Node) {
	switch node := root.(type) {
	case *ast.SourceFile:
		for _, child := range node.Nodes {
			validator.Validate(child)
		}
	case *ast.Block:
		for _, child := range node.Statements {
			validator.Validate(child)
		}
	case *ast.VariableDeclaration:
		validator.Validate(node.Expression)
	case *ast.FunctionDeclaration:

		if scope, ok := validator.CurrentScope.ChildScopes[node]; ok {
			validator.CurrentScope = scope
			validator.Validate(node.Block)
			validator.CurrentScope = scope.ParentScope
		}
	case *ast.RuleDeclaration:
		for _, output := range node.Outputs {
			validator.Validate(output)
		}

		for _, dependency := range node.Dependencies {
			validator.Validate(dependency)
		}

		if node.Block != nil {
			if scope, ok := validator.CurrentScope.ChildScopes[node]; ok {
				validator.CurrentScope = scope
				validator.Validate(node.Block)
				validator.CurrentScope = scope.ParentScope
			}
		}

		// TODO: Ensure that node.Derived is a rule function or rule function invocation if set
	case *ast.AliasDeclaration:
		validator.Validate(node.Expression)
	case *ast.RuleFunctionDeclaration:

		if scope, ok := validator.CurrentScope.ChildScopes[node]; ok {
			validator.CurrentScope = scope
			validator.Validate(node.Block)
			validator.CurrentScope = scope.ParentScope
		}
	case *ast.Equality:
		validator.Validate(node.Left)
		validator.Validate(node.Right)
		// TODO: Validate operand
	case *ast.Comparison:
		validator.Validate(node.Left)
		validator.Validate(node.Right)
		// TODO: Validate operand
	case *ast.Factor:
		validator.Validate(node.Left)
		validator.Validate(node.Right)
		// TODO: Validate operand
	case *ast.Term:
		validator.Validate(node.Left)
		validator.Validate(node.Right)
		// TODO: Validate operand
	case *ast.Unary:
		validator.Validate(node.Primary)
		// TODO: Validate operand
	case *ast.Primary:
		validator.Validate(node.Operand)
	case *ast.Identifier:
		validator.checkDefinedInScope(node.Value, node)
	case *ast.Invocation:
		validator.Validate(node.Operand)
		for _, argument := range node.Arguments {
			validator.Validate(argument)
		}
	case *ast.InterpretedString:
		// TODO: parse and check expressions
	case *ast.ReturnStatement:
		validator.Validate(node.Value)
		// TODO: Validate that the return statement belongs in a function
	case *ast.Assignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.AdditionAssignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.SubtractionAssignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.MultiplicationAssignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.DivisionAssignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.LooseAssignment:
		validator.Validate(node.Expression)
		validator.Validate(node.Value)
		// TODO: validate that expression is assignable
	case *ast.ImportSelector:
		if _, _, ok := validator.CurrentScope.LookupByName(node.From); !ok {
			validator.errorf(node, "'%s' is undefined. Did you forget to import it?", node.From)
		}
		validator.checkForTrait(node.From, node, TraitImport)
	}
}

func (validator *Validator) errorf(node ast.Node, format string, arguments ...interface{}) {
	r := ast.CreateRange(node.Start(), node.End())
	validator.errors = append(validator.errors, ast.CreateTreeError(&r, format, arguments...))
}

func (validator *Validator) checkDefinedInScope(name string, node ast.Node) {
	if symbol, scope, ok := validator.CurrentScope.LookupByName(name); !ok {
		validator.errorf(node, "'%s' is undefined", name)
	} else if scope == validator.CurrentScope && symbol.Node.Start().Line >= node.Start().Line {
		if validator.CurrentScope.ParentScope != nil {
			// Check the parent scope (in order to support shadowing)
			if _, _, ok := validator.CurrentScope.ParentScope.LookupByName(name); !ok {
				validator.errorf(node, "'%s' is used before it's declared", name)
			}
		} else {
			validator.errorf(node, "'%s' is used before it's declared", name)
		}
	}
}

func (validator *Validator) checkString(node ast.Node) {
	switch node.(type) {
	case *ast.InterpretedString:
		// TODO: validate
	case *ast.RawString:
		// Valid, do nothing
	default:
		validator.errorf(node, "'%s' is not a valid string")
	}
}

func (validator *Validator) checkForTrait(name string, node ast.Node, trait Trait) {
	if symbol, _, ok := validator.CurrentScope.LookupByName(name); ok {
		if !symbol.Trait.Has(trait) {
			labels := strings.Join(trait.Strings(), " or ")
			validator.errorf(node, "'%s' is not of %s", name, labels)
		}
	}
}
