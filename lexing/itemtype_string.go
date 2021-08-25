// Code generated by "stringer -type=ItemType"; DO NOT EDIT.

package lexing

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemStartOfInput-0]
	_ = x[ItemEndOfInput-1]
	_ = x[ItemError-2]
	_ = x[ItemAddition-3]
	_ = x[ItemAdditionAssign-4]
	_ = x[ItemSubtraction-5]
	_ = x[ItemSubtractionAssign-6]
	_ = x[ItemMultiplication-7]
	_ = x[ItemMultiplicationAssign-8]
	_ = x[ItemDivision-9]
	_ = x[ItemDivisionAssign-10]
	_ = x[ItemAssignment-11]
	_ = x[ItemLooseAssignment-12]
	_ = x[ItemEquals-13]
	_ = x[ItemNot-14]
	_ = x[ItemNotEqual-15]
	_ = x[ItemLessThan-16]
	_ = x[ItemLessThanOrEqual-17]
	_ = x[ItemGreaterThan-18]
	_ = x[ItemGreaterThanOrEqual-19]
	_ = x[ItemAnd-20]
	_ = x[ItemOr-21]
	_ = x[ItemSpread-22]
	_ = x[ItemIncrement-23]
	_ = x[ItemDecrement-24]
	_ = x[ItemLeftParentheses-25]
	_ = x[ItemRightParentheses-26]
	_ = x[ItemLeftBracket-27]
	_ = x[ItemRightBracket-28]
	_ = x[ItemLeftCurly-29]
	_ = x[ItemRightCurly-30]
	_ = x[ItemColon-31]
	_ = x[ItemColonColon-32]
	_ = x[ItemComma-33]
	_ = x[ItemDot-34]
	_ = x[ItemSubstitutionStart-35]
	_ = x[ItemSubstitutionEnd-36]
	_ = x[ItemDoubleQuote-37]
	_ = x[ItemKeywordPackage-38]
	_ = x[ItemKeywordImport-39]
	_ = x[ItemKeywordFunc-40]
	_ = x[ItemKeywordRule-41]
	_ = x[ItemKeywordExport-42]
	_ = x[ItemKeywordIf-43]
	_ = x[ItemKeywordElse-44]
	_ = x[ItemKeywordReturn-45]
	_ = x[ItemKeywordLet-46]
	_ = x[ItemKeywordShell-47]
	_ = x[ItemKeywordAlias-48]
	_ = x[ItemIdentifier-49]
	_ = x[ItemNewline-50]
	_ = x[ItemWhitespace-51]
	_ = x[ItemRawString-52]
	_ = x[ItemShellString-53]
	_ = x[ItemStringPart-54]
	_ = x[ItemInteger-55]
	_ = x[ItemBoolean-56]
	_ = x[ItemComment-57]
}

const _ItemType_name = "ItemStartOfInputItemEndOfInputItemErrorItemAdditionItemAdditionAssignItemSubtractionItemSubtractionAssignItemMultiplicationItemMultiplicationAssignItemDivisionItemDivisionAssignItemAssignmentItemLooseAssignmentItemEqualsItemNotItemNotEqualItemLessThanItemLessThanOrEqualItemGreaterThanItemGreaterThanOrEqualItemAndItemOrItemSpreadItemIncrementItemDecrementItemLeftParenthesesItemRightParenthesesItemLeftBracketItemRightBracketItemLeftCurlyItemRightCurlyItemColonItemColonColonItemCommaItemDotItemSubstitutionStartItemSubstitutionEndItemDoubleQuoteItemKeywordPackageItemKeywordImportItemKeywordFuncItemKeywordRuleItemKeywordExportItemKeywordIfItemKeywordElseItemKeywordReturnItemKeywordLetItemKeywordShellItemKeywordAliasItemIdentifierItemNewlineItemWhitespaceItemRawStringItemShellStringItemStringPartItemIntegerItemBooleanItemComment"

var _ItemType_index = [...]uint16{0, 16, 30, 39, 51, 69, 84, 105, 123, 147, 159, 177, 191, 210, 220, 227, 239, 251, 270, 285, 307, 314, 320, 330, 343, 356, 375, 395, 410, 426, 439, 453, 462, 476, 485, 492, 513, 532, 547, 565, 582, 597, 612, 629, 642, 657, 674, 688, 704, 720, 734, 745, 759, 772, 787, 801, 812, 823, 834}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
