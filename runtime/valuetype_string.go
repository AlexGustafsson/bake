// Code generated by "stringer -type=ValueType"; DO NOT EDIT.

package runtime

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValueTypeNumber-0]
	_ = x[ValueTypeString-1]
	_ = x[ValueTypeBool-2]
	_ = x[ValueTypeFunction-3]
	_ = x[ValueTypeRuleFunction-4]
	_ = x[ValueTypeRule-5]
	_ = x[ValueTypeNone-6]
	_ = x[ValueTypeArray-7]
}

const _ValueType_name = "ValueTypeNumberValueTypeStringValueTypeBoolValueTypeFunctionValueTypeRuleFunctionValueTypeRuleValueTypeNoneValueTypeArray"

var _ValueType_index = [...]uint8{0, 15, 30, 43, 60, 81, 94, 107, 121}

func (i ValueType) String() string {
	if i < 0 || i >= ValueType(len(_ValueType_index)-1) {
		return "ValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueType_name[_ValueType_index[i]:_ValueType_index[i+1]]
}
