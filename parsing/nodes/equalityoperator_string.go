// Code generated by "stringer -type=EqualityOperator"; DO NOT EDIT.

package nodes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EqualityOperatorOr-0]
	_ = x[EqualityOperatorAnd-1]
}

const _EqualityOperator_name = "EqualityOperatorOrEqualityOperatorAnd"

var _EqualityOperator_index = [...]uint8{0, 18, 37}

func (i EqualityOperator) String() string {
	if i < 0 || i >= EqualityOperator(len(_EqualityOperator_index)-1) {
		return "EqualityOperator(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EqualityOperator_name[_EqualityOperator_index[i]:_EqualityOperator_index[i+1]]
}