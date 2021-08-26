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
	visited      int
	seenImports  bool
	seenPackage  bool
	mayReturn    bool
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
	case *ast.PackageDeclaration:
		if validator.seenPackage {
			validator.errorf(node, "only one package may be declared in a file")
		} else {
			if validator.visited > 0 {
				validator.errorf(node, "package declaration must be declared at the top")
			}
		}

		validator.seenPackage = true
	case *ast.Block:
		for _, child := range node.Statements {
			validator.Validate(child)
		}
	case *ast.VariableDeclaration:
		validator.Validate(node.Expression)
	case *ast.FunctionDeclaration:

		if scope, ok := validator.CurrentScope.ChildScopes[node]; ok {
			validator.CurrentScope = scope
			previous := validator.mayReturn
			validator.mayReturn = true
			validator.Validate(node.Block)
			validator.mayReturn = previous
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
			previous := validator.mayReturn
			validator.mayReturn = true
			validator.Validate(node.Block)
			validator.mayReturn = previous
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

		// TODO: Implement recursion
		if identifier, ok := node.Operand.(*ast.Identifier); ok {
			if symbol, _, ok := validator.CurrentScope.LookupByName(identifier.Value); ok {
				if symbol.Trait.Has(TraitCallable) {
					if len(node.Arguments) < symbol.ArgumentCount {
						validator.errorf(identifier, "too few arguments. Expected %d, got %d", symbol.ArgumentCount, len(node.Arguments))
					} else if len(node.Arguments) > symbol.ArgumentCount {
						validator.errorf(identifier, "too many arguments. Expected %d, got %d", symbol.ArgumentCount, len(node.Arguments))
					}
				} else if symbol.Trait.Has(TraitAny) {
					// Nothing to check, unknown amount of arguments
				} else {
					validator.errorf(identifier, "not a function")
				}
			}
		}
	case *ast.EvaluatedString:
		for _, part := range node.Parts {
			// If it's not a string part, it's an expression
			if _, ok := part.(*ast.StringPart); !ok {
				validator.Validate(part)
			}
		}
	case *ast.ReturnStatement:
		if !validator.mayReturn {
			validator.errorf(node, "cannot return here")
		}

		validator.Validate(node.Value)

		if validator.CurrentScope.HasSeenReturn {
			validator.errorf(node, "dead return")
		}
		validator.CurrentScope.HasSeenReturn = true
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
	case *ast.ImportsDeclaration:
		if validator.seenPackage {
			if validator.visited > 1 {
				validator.errorf(node, "imports must be declared after the package")
			}
		} else {
			if validator.visited > 0 {
				validator.errorf(node, "imports must be declared at the top")
			}
		}

		if validator.seenImports {
			validator.errorf(node, "only one imports declaration may be used in a file")
		}

		for _, path := range node.Imports {
			for _, part := range path.Parts {
				if _, ok := part.(*ast.StringPart); !ok {
					validator.errorf(part, "imports may not have evaluations")
				}
			}
		}

		validator.seenImports = true
	case *ast.ShellStatement:
		for _, part := range node.Parts {
			validator.Validate(part)
		}
	case *ast.IfStatement:
		validator.Validate(node.Expression)

		if scope, ok := validator.CurrentScope.ChildScopes[node.PositiveBranch]; ok {
			validator.CurrentScope = scope
			validator.Validate(node.PositiveBranch)
			validator.CurrentScope = scope.ParentScope
		}

		if node.NegativeBranch != nil {
			if scope, ok := validator.CurrentScope.ChildScopes[node.NegativeBranch]; ok {
				validator.CurrentScope = scope
				validator.Validate(node.NegativeBranch)
				validator.CurrentScope = scope.ParentScope
			}
		}
	}

	validator.visited++
}

func (validator *Validator) errorf(node ast.Node, format string, arguments ...interface{}) {
	validator.errors = append(validator.errors, ast.CreateTreeError(node.Range(), format, arguments...))
}

func (validator *Validator) checkDefinedInScope(name string, node ast.Node) {
	// Lookup the identifier to make sure it's declared
	if symbol, scope, ok := validator.CurrentScope.LookupByName(name); ok {
		if scope == validator.CurrentScope && symbol.Node != nil && symbol.Node.Range().Start.Line >= node.Range().Start.Line {
			if validator.CurrentScope.ParentScope == nil {
				validator.errorf(node, "'%s' is used before it's declared", name)
			} else {
				// Check the parent scope (in order to support shadowing)
				if _, _, ok := validator.CurrentScope.ParentScope.LookupByName(name); !ok {
					validator.errorf(node, "'%s' is used before it's declared", name)
				}
			}
		}
	} else {
		validator.errorf(node, "'%s' is undefined", name)
	}
}

func (validator *Validator) checkString(node ast.Node) {
	switch node.(type) {
	case *ast.EvaluatedString:
		validator.Validate(node)
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
