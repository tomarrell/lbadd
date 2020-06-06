# Scanner

## Product requirements

Scanner fits in as a sub-component of the parser which helps in providing a tokenised version of the input. 

## Technical requirements
The scanner must be fast, with simple API which can be used by the parser to extract tokens from the input. 
A scanner must be able to take in a particular stream of input and return a component which has API for the scanner to extract tokens from.

## Technical design

The scanner will have a `Next` and a `Peek` function.
The `Next` function will provide subsequent tokens to the parser by consuming it whereas the `Peek` function will provide the token without consuming it.

## Implementation

The `Scanner` is an interface which is implemented by a `ruleBasedScanner`.

## Testing

* Specific token testing.
* Random generation testing.
