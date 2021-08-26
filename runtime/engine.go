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

	for _, node := range program.Source.Nodes {
		value := engine.evaluate(node)
		fmt.Println(value)
	}

	return nil
}

func (engine *Engine) evaluate(rootNode ast.Node) *Value {
	switch node := rootNode.(type) {
	case *ast.VariableDeclaration:
		value := engine.evaluate(node.Expression)
		engine.Delegate.DeclareVariable(node.Identifier, value)
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
	}

	return nil
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
