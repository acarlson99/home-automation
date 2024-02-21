package expr

import (
	"fmt"
	"reflect"

	hpb "github.com/acarlson99/home-automation/proto/go"
	"golang.org/x/exp/constraints"
)

func EvalVar(expr *hpb.Var) (*hpb.Primitive, error) {
	switch e := expr.GetV().(type) {
	case *hpb.Var_DeviceState:
		// TODO: locate corresponding device
		// TODO: determine if device has requested information
		// TODO: return said information in a user-friendly manner
		return nil, nil // TODO: this
	case *hpb.Var_BinOp:
		return EvalBinaryOperator(e.BinOp)
	case *hpb.Var_UnaryOp:
		return EvalUnaryOperator(e.UnaryOp)
	case *hpb.Var_Cmp:
		v, err := EvalComparison(e.Cmp)
		if err != nil {
			return nil, err
		}
		return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: v}}, nil
	case *hpb.Var_Prim:
		return e.Prim, nil
	}
	return nil, fmt.Errorf("no evaluation available for expr %+v", expr)
}

func cmp[T constraints.Ordered](op *hpb.Comparison_Operator, lhs, rhs T) (bool, error) {
	switch *op.Enum() {
	case hpb.Comparison_GT:
		return lhs > rhs, nil
	case hpb.Comparison_LT:
		return lhs < rhs, nil
	case hpb.Comparison_GTEQ:
		return lhs >= rhs, nil
	case hpb.Comparison_LTEQ:
		return lhs <= rhs, nil
	case hpb.Comparison_EQ:
		return lhs == rhs, nil
	}
	return false, fmt.Errorf("invalid operation %v for types %T %T", op, lhs, rhs)
}

// todo: nvidia chatwithrtx

func EvalComparison(expr *hpb.Comparison) (bool, error) {
	lhs, err := EvalVar(expr.Lhs)
	if err != nil {
		return false, err
	}
	rhs, err := EvalVar(expr.Rhs)
	if err != nil {
		return false, err
	}
	if ta, tb := reflect.TypeOf(lhs.GetV()), reflect.TypeOf(rhs.GetV()); ta != tb {
		return false, fmt.Errorf("improper types for operation %v: %v != %v", expr.GetOp().Enum(), ta, tb)
	}

	switch lhs.GetV().(type) {
	case *hpb.Primitive_Bool:
		if expr.GetOp().Enum() != hpb.Comparison_EQ.Enum() {
			return false, fmt.Errorf("bad operation for bools: %v", expr.GetOp().Enum())
		}
		return lhs.GetBool() == rhs.GetBool(), nil
	case *hpb.Primitive_Float:
		return cmp(expr.GetOp().Enum(), lhs.GetFloat(), rhs.GetFloat())
	case *hpb.Primitive_Int32:
		return cmp(expr.GetOp().Enum(), lhs.GetInt32(), rhs.GetInt32())
	case *hpb.Primitive_String_:
		return cmp(expr.GetOp().Enum(), lhs.GetString_(), rhs.GetString_())
	}

	return false, fmt.Errorf("invalid operation %v for types %T %T", expr.GetOp().Enum(), lhs, rhs)
}

func opInt[T constraints.Integer](op *hpb.BinaryOperator_Operation, lhs, rhs T) (T, error) {
	switch *op.Enum() {
	case hpb.BinaryOperator_ADD:
		return lhs + rhs, nil
	case hpb.BinaryOperator_SUB:
		return lhs - rhs, nil
	case hpb.BinaryOperator_MUL:
		return lhs * rhs, nil
	case hpb.BinaryOperator_DIV:
		return lhs / rhs, nil
	case hpb.BinaryOperator_MIN:
		return min(lhs, rhs), nil
	case hpb.BinaryOperator_MAX:
		return max(lhs + rhs), nil
	case hpb.BinaryOperator_OR:
		return lhs | rhs, nil
	case hpb.BinaryOperator_AND:
		return lhs & rhs, nil
	case hpb.BinaryOperator_XOR:
		return lhs ^ rhs, nil
	}
	return *new(T), fmt.Errorf("could not match operation %v for values %T %T", op.Enum(), lhs, rhs)
}

func opFloat[T constraints.Float](op *hpb.BinaryOperator_Operation, lhs, rhs T) (T, error) {
	switch *op.Enum() {
	case hpb.BinaryOperator_ADD:
		return lhs + rhs, nil
	case hpb.BinaryOperator_SUB:
		return lhs - rhs, nil
	case hpb.BinaryOperator_MUL:
		return lhs * rhs, nil
	case hpb.BinaryOperator_DIV:
		return lhs / rhs, nil
	case hpb.BinaryOperator_MIN:
		return min(lhs, rhs), nil
	case hpb.BinaryOperator_MAX:
		return max(lhs + rhs), nil
	}
	return *new(T), fmt.Errorf("could not match operation %v for types %T %T", op.Enum(), lhs, rhs)
}

func EvalBinaryOperator(expr *hpb.BinaryOperator) (*hpb.Primitive, error) {
	lhs, err := EvalVar(expr.Lhs)
	if err != nil {
		return nil, err
	}
	rhs, err := EvalVar(expr.Rhs)
	if err != nil {
		return nil, err
	}
	if ta, tb := reflect.TypeOf(lhs.GetV()), reflect.TypeOf(rhs.GetV()); ta != tb {
		return nil, fmt.Errorf("improper types for operation %v: %v != %v", expr.GetOp().Enum(), ta, tb)
	}
	switch lhs.GetV().(type) {
	case *hpb.Primitive_Int32:
		v, err := opInt(expr.GetOp().Enum(), lhs.GetInt32(), rhs.GetInt32())
		if err != nil {
			return nil, err
		}
		return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: v}}, nil
	case *hpb.Primitive_Float:
		v, err := opFloat(expr.GetOp().Enum(), lhs.GetFloat(), rhs.GetFloat())
		if err != nil {
			return nil, err
		}
		return &hpb.Primitive{V: &hpb.Primitive_Float{Float: v}}, nil
	case *hpb.Primitive_String_:
		lhs, rhs := lhs.GetString_(), rhs.GetString_()
		s := ""
		switch expr.GetOp() {
		case hpb.BinaryOperator_ADD:
			s = lhs + rhs
		case hpb.BinaryOperator_MIN:
			if lhs < rhs {
				s = lhs
			} else {
				s = rhs
			}
		case hpb.BinaryOperator_MAX:
			if lhs > rhs {
				s = lhs
			} else {
				s = rhs
			}
		default:
			return nil, fmt.Errorf("invalid operator for types %T %T: %v", lhs, rhs, expr.GetOp())
		}
		return &hpb.Primitive{V: &hpb.Primitive_String_{String_: s}}, nil
	case *hpb.Primitive_Bool:
		lhs, rhs := lhs.GetBool(), rhs.GetBool()
		b := false
		switch expr.GetOp() {
		case hpb.BinaryOperator_OR:
			b = lhs || rhs
		case hpb.BinaryOperator_AND:
			b = lhs && rhs
		case hpb.BinaryOperator_XOR:
			b = lhs != rhs
		}
		return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: b}}, nil
	}
	return nil, fmt.Errorf("could not match operation %v for types %T %T", expr.GetOp().Enum(), lhs, rhs)
}

func EvalUnaryOperator(expr *hpb.UnaryOperator) (*hpb.Primitive, error) {
	v, err := EvalVar(expr.X)
	if err != nil {
		return nil, err
	}
	switch expr.GetOp() {
	case hpb.UnaryOperator_NEG:
		switch x := v.V.(type) {
		case *hpb.Primitive_Int32:
			return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: -x.Int32}}, nil
		case *hpb.Primitive_Float:
			return &hpb.Primitive{V: &hpb.Primitive_Float{Float: -x.Float}}, nil
		default:
			return nil, fmt.Errorf("unsupported operation %v for type %T", expr.GetOp().Enum(), x)
		}
	case hpb.UnaryOperator_NOT:
		switch x := v.V.(type) {
		case *hpb.Primitive_Bool:
			return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: !x.Bool}}, nil
		default:
			return nil, fmt.Errorf("unsupported operation %v for type %T", expr.GetOp().Enum(), x)
		}
	case hpb.UnaryOperator_CAST_BOOL:
		switch x := v.V.(type) {
		case *hpb.Primitive_Int32:
			return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: x.Int32 != 0}}, nil
		case *hpb.Primitive_Float:
			return &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: x.Float != 0}}, nil
		case *hpb.Primitive_Bool:
			return v, nil
		default:
			return nil, fmt.Errorf("unsupported operation %v for type %T", expr.GetOp().Enum(), x)
		}
	case hpb.UnaryOperator_CAST_INT:
		switch x := v.V.(type) {
		case *hpb.Primitive_Int32:
			return v, nil
		case *hpb.Primitive_Float:
			return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(x.Float)}}, nil
		case *hpb.Primitive_Bool:
			n := int32(0)
			if x.Bool {
				n = 1
			}
			return &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: n}}, nil
		}
	case hpb.UnaryOperator_CAST_FLOAT:
		switch x := v.V.(type) {
		case *hpb.Primitive_Int32:
			return &hpb.Primitive{V: &hpb.Primitive_Float{Float: float32(x.Int32)}}, nil
		case *hpb.Primitive_Float:
			return &hpb.Primitive{V: &hpb.Primitive_Float{Float: float32(x.Float)}}, nil
		case *hpb.Primitive_Bool:
			n := float32(0)
			if x.Bool {
				n = 1
			}
			return &hpb.Primitive{V: &hpb.Primitive_Float{Float: n}}, nil
		default:
			return nil, fmt.Errorf("unsupported operation %v for type %T", expr.GetOp().Enum(), x)
		}
	}
	return nil, fmt.Errorf("unsupported operation %v for expr %v", expr.GetOp().Enum(), v)
}
