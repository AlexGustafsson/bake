// Code generated by "stringer -type=AdditiveOperator"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AdditiveOperatorAddition-0]
	_ = x[AdditiveOperatorSubtraction-1]
}

const _AdditiveOperator_name = "AdditiveOperatorAdditionAdditiveOperatorSubtraction"

var _AdditiveOperator_index = [...]uint8{0, 24, 51}

func (i AdditiveOperator) String() string {
	if i < 0 || i >= AdditiveOperator(len(_AdditiveOperator_index)-1) {
		return "AdditiveOperator(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AdditiveOperator_name[_AdditiveOperator_index[i]:_AdditiveOperator_index[i+1]]
}
