package optimization

import "github.com/tomarrell/lbadd/internal/compiler/command"

// Optimization defines a process that optimizes an input command and outputs a
// modified, optimized version of that command, if the optimization is
// applicable to the input command. If not, ok=false will be returned.
type Optimization func(command.Command) (optimized command.Command, ok bool)
