package runtime

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/AlexGustafsson/bake/ast"
	log "github.com/sirupsen/logrus"
)

type Engine struct {
	Delegate    Delegate
	returnValue *Value
}

func CreateEngine(delegate Delegate) *Engine {
	return &Engine{
		Delegate: delegate,
	}
}

func (engine *Engine) Evaluate(source *ast.Block) (err error) {
	defer engine.recover(&err)

	engine.evaluate(source)

	return nil
}

func (engine *Engine) EvaluateTask(task string) (err error) {
	defer engine.recover(&err)

	log.Debugf("resolving task '%s'", task)
	value := engine.Delegate.Resolve(task)
	log.Debugf("resolved task '%s': %s", task, value)

	log.Debugf("evaluating task '%s'", task)
	engine.evaluateTask(value)

	return nil
}

func (engine *Engine) evaluateTask(value *Value) {
	switch value.Type {
	case ValueTypeFunction:
		if !value.Exported {
			panic(fmt.Errorf("cannot invoke a private function"))
		}

		function := value.Value.(*Function)
		// The function is a builtin - don't call
		if function.Block == nil {
			panic(fmt.Errorf("function not callable"))
		}

		// TODO: parse arguments
		if len(function.Arguments) > 0 {
			panic(fmt.Errorf("missing arguments"))
		}

		// Create a function scope
		engine.Delegate.PushScope()

		// TODO: define arguments

		// Evaluate block
		engine.evaluate(function.Block)

		// Leave function scope
		engine.Delegate.PopScope()
		engine.returnValue = nil
	case ValueTypeAlias:
		if !value.Exported {
			panic(fmt.Errorf("cannot invoke a private alias"))
		}

		// TODO: Build dependency table (don't run twice, watch for changes etc.)
		alias := value.Value.(*Alias)
		for _, dependency := range alias.Dependencies {
			engine.evaluateTask(dependency)
		}
	case ValueTypeString:
		engine.EvaluateTask(value.Value.(string))
	case ValueTypeRule:
		rule := value.Value.(*Rule)
		// Create a function scope
		engine.Delegate.PushScope()

		// Evaluate block
		engine.evaluate(rule.Block)

		// Leave function scope
		engine.Delegate.PopScope()
	default:
		panic(fmt.Errorf("cannot invoke type '%s'", value.Type))
	}
}

func (engine *Engine) evaluate(rootNode ast.Node) *Value {
	switch node := rootNode.(type) {
	case *ast.VariableDeclaration:
		value := engine.evaluate(node.Expression)
		engine.Delegate.Define(node.Identifier, value)
		return nil
	case *ast.Term:
		left := engine.evaluate(node.Left)
		right := engine.evaluate(node.Right)
		switch node.Operator {
		case ast.AdditiveOperatorAddition:
			return engine.Delegate.Add(left, right)
		case ast.AdditiveOperatorSubtraction:
			return engine.Delegate.Subtract(left, right)
		}
	case *ast.Factor:
		left := engine.evaluate(node.Left)
		right := engine.evaluate(node.Right)
		switch node.Operator {
		case ast.MultiplicativeOperatorMultiplication:
			return engine.Delegate.Multiply(left, right)
		case ast.MultiplicativeOperatorDivision:
			return engine.Delegate.Divide(left, right)
		case ast.MultiplicativeOperatorModulo:
			return engine.Delegate.Modulo(left, right)
		}
	case *ast.Integer:
		value := engine.parseInteger(node.Value)
		return &Value{
			Type:  ValueTypeNumber,
			Value: value,
		}
	case *ast.Boolean:
		value := engine.parseBool(node.Value)
		return &Value{
			Type:  ValueTypeBool,
			Value: value,
		}
	case *ast.Comparison:
		left := engine.evaluate(node.Left)
		right := engine.evaluate(node.Right)
		switch node.Operator {
		case ast.ComparisonOperatorEquals:
			return engine.Delegate.Equals(left, right)
		case ast.ComparisonOperatorGreaterThan:
			return engine.Delegate.GreaterThan(left, right)
		case ast.ComparisonOperatorGreaterThanOrEqual:
			return engine.Delegate.GreaterThanOrEqual(left, right)
		case ast.ComparisonOperatorLessThan:
			return engine.Delegate.LessThan(left, right)
		case ast.ComparisonOperatorLessThanOrEqual:
			return engine.Delegate.LessThanOrEqual(left, right)
		}
	case *ast.Equality:
		left := engine.evaluate(node.Left)
		right := engine.evaluate(node.Right)
		switch node.Operator {
		case ast.EqualityOperatorAnd:
			return engine.Delegate.And(left, right)
		case ast.EqualityOperatorOr:
			return engine.Delegate.Or(left, right)
		}
	case *ast.Unary:
		operand := engine.evaluate(node.Primary)
		switch node.Operator {
		case ast.UnaryOperatorNot:
			return engine.Delegate.Not(operand)
		case ast.UnaryOperatorSubtraction:
			return engine.Delegate.Negative(operand)
		}
	case *ast.Identifier:
		return engine.Delegate.Resolve(node.Value)
	case *ast.Assignment:
		value := engine.evaluate(node.Value)
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.Increment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			value = engine.Delegate.Add(value, &Value{Type: ValueTypeNumber, Value: 1})
			engine.Delegate.Define(identifier.Value, value)
			return value
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
	case *ast.Decrement:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			value = engine.Delegate.Subtract(value, &Value{Type: ValueTypeNumber, Value: 1})
			engine.Delegate.Define(identifier.Value, value)
			return value
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
	case *ast.AdditionAssignment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			expression := engine.evaluate(node.Value)
			value = engine.Delegate.Add(value, expression)
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.SubtractionAssignment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			expression := engine.evaluate(node.Value)
			value = engine.Delegate.Subtract(value, expression)
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.MultiplicationAssignment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			expression := engine.evaluate(node.Value)
			value = engine.Delegate.Multiply(value, expression)
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.DivisionAssignment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			expression := engine.evaluate(node.Value)
			value = engine.Delegate.Divide(value, expression)
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.ModuloAssignment:
		// TODO: Implement objects
		if identifier, ok := node.Expression.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)
			expression := engine.evaluate(node.Value)
			value = engine.Delegate.Modulo(value, expression)
			engine.Delegate.Define(identifier.Value, value)
		} else {
			panic(fmt.Errorf("cannot assign to type %s", node.Type()))
		}
		return nil
	case *ast.Invocation:
		// TODO: Implement objects
		if identifier, ok := node.Operand.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)

			if function, ok := value.Value.(*Function); ok && !function.IsRuleFunction {
				if function.Handler != nil {
					// Evaluate arguments
					arguments := make([]*Value, len(node.Arguments))
					for i, argument := range node.Arguments {
						value := engine.evaluate(argument)
						arguments[i] = value
					}

					returnValue := function.Handler(engine, arguments)
					if returnValue == nil {
						returnValue = &Value{Type: ValueTypeNone}
					}
					return returnValue
				} else {
					if len(node.Arguments) != len(function.Arguments) {
						panic(fmt.Errorf("invocation argument mismatch"))
					}

					// Create a function scope
					engine.Delegate.PushScope()

					// Define arguments
					for i, parameter := range function.Arguments {
						value := engine.evaluate(node.Arguments[i])
						engine.Delegate.Define(parameter, value)
					}

					// Evaluate block
					engine.evaluate(function.Block)
					returnValue := engine.returnValue
					if returnValue == nil {
						returnValue = &Value{Type: ValueTypeNone}
					}

					// Leave function scope
					engine.Delegate.PopScope()
					engine.returnValue = nil
					return returnValue
				}
			} else {
				panic(fmt.Errorf("%s is not a function", node.Type()))
			}
		} else {
			panic(fmt.Errorf("cannot call type %s", node.Type()))
		}
	case *ast.IfStatement:
		condition := engine.evaluate(node.Expression)
		if condition.Type != ValueTypeBool {
			panic(fmt.Errorf("invalid condition"))
		}

		if condition.Value.(bool) {
			engine.Delegate.PushScope()
			engine.evaluate(node.PositiveBranch)
			engine.Delegate.PopScope()
		} else if node.NegativeBranch != nil {
			engine.Delegate.PushScope()
			engine.evaluate(node.NegativeBranch)
			engine.Delegate.PopScope()
		}
		return nil
	case *ast.ForStatement:
		value := engine.evaluate(node.Expression)
		if value.Type != ValueTypeArray {
			panic(fmt.Errorf("invalid collection"))
		}

		collection := value.Value.([]*Value)
		for _, currentValue := range collection {
			engine.Delegate.PushScope()
			engine.Delegate.Define(node.Identifier.Value, currentValue)
			engine.evaluate(node.Block)
			engine.Delegate.PopScope()
		}
		return nil
	case *ast.Block:
		// Define all functions, rule functions and rules first
		for _, statement := range node.Statements {
			switch declaration := statement.(type) {
			case *ast.FunctionDeclaration:
				function := &Function{Arguments: make([]string, 0), Block: declaration.Block}
				if declaration.Signature != nil {
					for _, identifier := range declaration.Signature.Arguments {
						function.Arguments = append(function.Arguments, identifier.Value)
					}
				}

				value := &Value{Type: ValueTypeFunction, Value: function, Exported: declaration.Exported}
				engine.Delegate.Define(declaration.Identifier, value)
			case *ast.RuleFunctionDeclaration:
				function := &Function{Arguments: make([]string, 0), IsRuleFunction: true, Block: declaration.Block}
				if declaration.Signature != nil {
					for _, identifier := range declaration.Signature.Arguments {
						function.Arguments = append(function.Arguments, identifier.Value)
					}
				}

				value := &Value{Type: ValueTypeRuleFunction, Value: function, Exported: declaration.Exported}
				engine.Delegate.Define(declaration.Identifier, value)
			case *ast.RuleDeclaration:
				rule := &Rule{
					Outputs:      make([]string, 0),
					Dependencies: make([]*Value, 0),
				}

				value := &Value{Type: ValueTypeRule, Value: rule}

				for _, outputNode := range declaration.Outputs {
					output := engine.evaluate(outputNode)
					if output.Type != ValueTypeString {
						panic(fmt.Errorf("an output evaluate to a string"))
					}

					rule.Outputs = append(rule.Outputs, output.Value.(string))
					engine.Delegate.Define(output.Value.(string), value)
				}

				for _, dependencyNode := range declaration.Dependencies {
					dependency := engine.evaluate(dependencyNode)
					switch dependency.Type {
					case ValueTypeFunction, ValueTypeString, ValueTypeAlias:
						// Do nothing, valid types
						rule.Dependencies = append(rule.Dependencies, dependency)
					default:
						panic(fmt.Errorf("invalid dependency type '%s'", dependency.Type))
					}
				}

				rule.Block = declaration.Block

				// TODO: Add support for derived rules (syntax sugar for another block?)
			case *ast.AliasDeclaration:
				// TODO: This is not compatible with the syntax for derived functions: alias build : c::build(param)
				// For now, only support array dependencies
				expression := engine.evaluate(declaration.Expression)
				if expression.Type != ValueTypeArray {
					panic(fmt.Errorf("invalid dependencies - expected array"))
				}

				dependencies := expression.Value.([]*Value)
				for _, dependency := range dependencies {
					switch dependency.Type {
					case ValueTypeFunction, ValueTypeString, ValueTypeAlias:
						// Do nothing, valid types
					default:
						panic(fmt.Errorf("invalid dependency type '%s'", dependency.Type))
					}
				}

				alias := &Alias{Dependencies: dependencies}
				value := &Value{Type: ValueTypeAlias, Value: alias, Exported: declaration.Exported}
				engine.Delegate.Define(declaration.Identifier, value)
			}
		}

		// Evaluate all statements
		for _, statement := range node.Statements {
			switch statement.Type() {
			case ast.NodeTypeFunctionDeclaration, ast.NodeTypeRuleFunctionDeclaration, ast.NodeTypeRuleDeclaration, ast.NodeTypeAliasDeclaration:
				// Do nothing
			case ast.NodeTypeReturnStatement:
				returnStatement := statement.(*ast.ReturnStatement)
				value := engine.evaluate(returnStatement.Value)
				engine.returnValue = value
			default:
				engine.evaluate(statement)
			}

			// Prematurely stop evaluating the block if returned
			if engine.returnValue != nil {
				return nil
			}
		}
		return nil
	case *ast.RawString:
		return &Value{
			Type:  ValueTypeString,
			Value: node.Content,
		}
	case *ast.EvaluatedString:
		var builder strings.Builder
		for _, part := range node.Parts {
			switch partNode := part.(type) {
			case *ast.StringPart:
				// TODO: Expand escape codes like \n, \t etc.?
				builder.WriteString(partNode.Content)
			default:
				value := engine.evaluate(partNode)
				builder.WriteString(value.String())
			}
		}
		return &Value{
			Type:  ValueTypeString,
			Value: builder.String(),
		}
	case *ast.Array:
		elements := make([]*Value, len(node.Elements))
		for i, element := range node.Elements {
			value := engine.evaluate(element)
			elements[i] = value
		}
		return &Value{Type: ValueTypeArray, Value: elements}
	}

	panic(fmt.Errorf("unimplemented type %s", rootNode.Type()))
}

func (engine *Engine) recover(errp *error) {
	err := recover()
	if err != nil {
		if _, ok := err.(runtime.Error); ok {
			panic(err)
		}

		if engine != nil {
			*errp = err.(error)
		}
	}
}

func (engine *Engine) parseInteger(raw string) int {
	value, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(value)
}

func (engine *Engine) parseBool(raw string) bool {
	value, err := strconv.ParseBool(raw)
	if err != nil {
		panic(err)
	}

	return value
}
