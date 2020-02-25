# Rule based scanner

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

