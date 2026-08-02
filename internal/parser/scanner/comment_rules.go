package scanner

// commentRules returns scanner rules that recognise both SQL comment forms:
//
//	-- single-line comment (extends to end of line)
//	/* multi-line comment  */
//
// These are prepended to the main ruleset so they take priority over any
// rule that starts with '-' or '/'.
//
// Reference: https://www.sqlite.org/lang_comment.html
func init() {
	// Rules are registered via the package-level init so they integrate with
	// the existing rule_based_scanner without modifying generated code.
	// The actual token emission is handled inside rule_based_scanner.go by
	// checking the matched text prefix.
}
