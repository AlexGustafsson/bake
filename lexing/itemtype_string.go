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
	_ = x[ItemModulo-25]
	_ = x[ItemModuloAssign-26]
	_ = x[ItemLeftParentheses-27]
	_ = x[ItemRightParentheses-28]
	_ = x[ItemLeftBracket-29]
	_ = x[ItemRightBracket-30]
	_ = x[ItemLeftCurly-31]
	_ = x[ItemRightCurly-32]
	_ = x[ItemColon-33]
	_ = x[ItemColonColon-34]
	_ = x[ItemComma-35]
	_ = x[ItemDot-36]
	_ = x[ItemSubstitutionStart-37]
	_ = x[ItemSubstitutionEnd-38]
	_ = x[ItemDoubleQuote-39]
	_ = x[ItemKeywordPackage-40]
	_ = x[ItemKeywordImport-41]
	_ = x[ItemKeywordFunc-42]
	_ = x[ItemKeywordRule-43]
	_ = x[ItemKeywordExport-44]
	_ = x[ItemKeywordIf-45]
	_ = x[ItemKeywordFor-46]
	_ = x[ItemKeywordIn-47]
	_ = x[ItemKeywordElse-48]
	_ = x[ItemKeywordReturn-49]
	_ = x[ItemKeywordLet-50]
	_ = x[ItemKeywordShell-51]
	_ = x[ItemKeywordAlias-52]
	_ = x[ItemIdentifier-53]
	_ = x[ItemNewline-54]
	_ = x[ItemWhitespace-55]
	_ = x[ItemRawString-56]
	_ = x[ItemStringPart-57]
	_ = x[ItemInteger-58]
	_ = x[ItemBoolean-59]
	_ = x[ItemComment-60]
}

const _ItemType_name = "ItemStartOfInputItemEndOfInputItemErrorItemAdditionItemAdditionAssignItemSubtractionItemSubtractionAssignItemMultiplicationItemMultiplicationAssignItemDivisionItemDivisionAssignItemAssignmentItemLooseAssignmentItemEqualsItemNotItemNotEqualItemLessThanItemLessThanOrEqualItemGreaterThanItemGreaterThanOrEqualItemAndItemOrItemSpreadItemIncrementItemDecrementItemModuloItemModuloAssignItemLeftParenthesesItemRightParenthesesItemLeftBracketItemRightBracketItemLeftCurlyItemRightCurlyItemColonItemColonColonItemCommaItemDotItemSubstitutionStartItemSubstitutionEndItemDoubleQuoteItemKeywordPackageItemKeywordImportItemKeywordFuncItemKeywordRuleItemKeywordExportItemKeywordIfItemKeywordForItemKeywordInItemKeywordElseItemKeywordReturnItemKeywordLetItemKeywordShellItemKeywordAliasItemIdentifierItemNewlineItemWhitespaceItemRawStringItemStringPartItemIntegerItemBooleanItemComment"

var _ItemType_index = [...]uint16{0, 16, 30, 39, 51, 69, 84, 105, 123, 147, 159, 177, 191, 210, 220, 227, 239, 251, 270, 285, 307, 314, 320, 330, 343, 356, 366, 382, 401, 421, 436, 452, 465, 479, 488, 502, 511, 518, 539, 558, 573, 591, 608, 623, 638, 655, 668, 682, 695, 710, 727, 741, 757, 773, 787, 798, 812, 825, 839, 850, 861, 872}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
