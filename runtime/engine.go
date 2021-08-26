package runtime

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/AlexGustafsson/bake/ast"
)

type Engine struct {
	Delegate Delegate
}

func CreateEngine(delegate Delegate) *Engine {
	return &Engine{
		Delegate: delegate,
	}
}

func (engine *Engine) Evaluate(program *Program) (err error) {
	defer engine.recover(&err)

	engine.evaluate(program.Source)

	return nil
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
	case *ast.Invocation:
		// TODO: Implement objects
		if identifier, ok := node.Operand.(*ast.Identifier); ok {
			value := engine.Delegate.Resolve(identifier.Value)

			if function, ok := value.Value.(*Function); ok && !function.IsRuleFunction {
				if len(node.Arguments) != len(function.Arguments) {
					panic(fmt.Errorf("invocation argument mismatch"))
				}

				// Create a function scope
				engine.Delegate.PushScope()

				// Define parameters
				for i, parameter := range function.Arguments {
					value := engine.evaluate(node.Arguments[i])
					engine.Delegate.Define(parameter, value)
				}

				// Evaluate block
				returnValue := engine.evaluate(function.Block)

				// Leave function scope
				engine.Delegate.PopScope()
				return returnValue
			} else {
				panic(fmt.Errorf("not a function"))
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
	case *ast.Block:
		// Define all functions, rule functions and rules first
		for _, node := range node.Statements {
			switch declaration := node.(type) {
			case *ast.FunctionDeclaration:
				function := &Function{Arguments: make([]string, 0), Block: declaration.Block}
				if declaration.Signature != nil {
					for _, identifier := range declaration.Signature.Arguments {
						function.Arguments = append(function.Arguments, identifier.Value)
					}
				}

				value := &Value{Type: ValueTypeFunction, Value: function}
				engine.Delegate.Define(declaration.Identifier, value)
			case *ast.RuleFunctionDeclaration:
				function := &Function{Arguments: make([]string, 0), IsRuleFunction: true, Block: declaration.Block}
				if declaration.Signature != nil {
					for _, identifier := range declaration.Signature.Arguments {
						function.Arguments = append(function.Arguments, identifier.Value)
					}
				}

				value := &Value{Type: ValueTypeRuleFunction, Value: function}
				engine.Delegate.Define(declaration.Identifier, value)
			case *ast.RuleDeclaration:
				// rule := &Rule{}
				// value := &Value{Type: ValueTypeRule, Value: rule}
				// engine.Delegate.Define(declaration.Identifier, value)
			}
		}

		// Evaluate all statements
		for _, node := range node.Statements {
			switch node.Type() {
			case ast.NodeTypeFunctionDeclaration, ast.NodeTypeRuleFunctionDeclaration, ast.NodeTypeRuleDeclaration:
				// Do nothing
			case ast.NodeTypeReturnStatement:
				// Prematurely stop evaluating the block
				returnStatement := node.(*ast.ReturnStatement)
				value := engine.evaluate(returnStatement.Value)
				return value
			default:
				value := engine.evaluate(node)
				fmt.Println(value)
			}
		}
		return nil
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
