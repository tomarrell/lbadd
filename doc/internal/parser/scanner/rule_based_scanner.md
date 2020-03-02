# Rule based scanner

## Background
RuleBasedScanner was inspired by Eclipse's syntax highlighting, which works on top of a generic rule based scanner.
Rules are defined and added to a scanner.
For every rune in the input, the scanner tries to apply all rules.
If the rule can be applied (the rule returns a result that contains whether the rule is applicable), the scanner processes the produced token and advances to the next rune.
Every rule can advance the scanner's position, so that the scanner does not really need to apply all rules to all runes (which would be expensive).

The same approach was chosen here.
Rules are defined and added to the scanner.
Before applying a rule, the scanner's state is snapshotted.
After applying a rule, which changes the scanner's state, if the rule wasn't applicable, the scanner's state is reset to the snapshot, then the next rule is applied.
If the rule was applicable, the scanner's position has been changed by the rule, and the scanner can continue trying to apply all rules again, but this time to the next rune.

---

```
Hello World
```
```go
var wordRule = func(s RuneScanner) (typ token.Type, ruleWasApplicable bool) {
	for {
		nextRune, hasMoreRunes = s.Next()
		if !hasMoreRunes {
			ruleWasApplicable = false
			return
		}
		if allowedRunes.Matches(nextRune) {
			s.ConsumeRune()
		} else {
			break
		}
	}
	typ = token.Word
	ruleWasApplicable = true
	return
}
```

The above example is a sample implementation of a word rule, that accepts words that may consist of a defined set of runes (`allowedRunes`).
The above code reads as follows.
```
while
	peek the next rune that lies ahead
	if there are no more runes
		mark the rule as not applicable, because it cannot be applied
		return from the rule
	if the next rune is element of the allowed runes
		consume the rune
	else
		break the loop, which stops reading more runes and returns from the function successfully
set the token type to token.Word (*)
mark the rule as applicable
return from the rule
```
`(*)` the `RuneScanner` remembers all consumed runes, and from all consumed runes and the returned token type, **if** the rule was marked as applicable, the rule based scanner will then create a token.
The `RuneScanner` will be passed in from the rule based scanner, you don't have to and shouldn't worry about what it does.
We won't specify exactly what kind of `RuneScanner` will be passed in, because we probably will change the functionality in the future without notice or documentation.
Also, you shouldn't rely on internal documentation, and just work with the interface you are provided here.
At the moment (first implementation, Feb. 2020), a rule based scanner implements `RuneScanner`, and it passes itself, but that may change in the future, as mentioned.

## Implementation
Rule based scanner implements the scanner interface. 

```
type ruleBasedScanner struct {
	input []rune

	cache token.Token

	whitespaceDetector ruleset.DetectorFunc
	linefeedDetector   ruleset.DetectorFunc
	rules              []ruleset.Rule

	state
}
``` 
Above is the structure of a `ruleBasedScanner`.

* `Next` :
  Internally calls the `Peek` function to get a token and removes from the cache, effectively "removing" from the stream.
 
* `Peek` :
  `Peek` enables the computing of the next token by applying the rules using a `computeNext` function. 

* `applyRule` : Called by the `computeNext` function; applies multiple preset rules provided to the scanner. Rules are applied in a given order. Each state before applying is check-pointed in order to be able to restore the state on a bad rule applied. 
   On success, a token is emitted from the function. If all rules are applied and no token is found, the offending rule is skipped and an error token is emitted, with the value of the offending token.
   
## Testing

