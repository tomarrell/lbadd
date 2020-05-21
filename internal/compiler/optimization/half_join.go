package optimization

import "github.com/tomarrell/lbadd/internal/compiler/command"

// OptHalfJoin reduces Joins that are of the form Join(any,nil) or Join(nil,any)
// to just any.
func OptHalfJoin(cmd command.Command) (command.Command, bool) {
	switch c := cmd.(type) {
	case command.Select:
		if optimized, ok := OptHalfJoin(c.Input); ok {
			return command.Select{
				Filter: c.Filter,
				Input:  optimized.(command.List),
			}, true
		}
	case command.Project:
		if optimized, ok := OptHalfJoin(c.Input); ok {
			return command.Project{
				Cols:  c.Cols,
				Input: optimized.(command.List),
			}, true
		}
	case command.Limit:
		if optimized, ok := OptHalfJoin(c.Input); ok {
			return command.Limit{
				Limit: c.Limit,
				Input: optimized.(command.List),
			}, true
		}
	case command.Join:
		if c.Left == nil && c.Right == nil {
			return nil, false
		}
		if c.Left == nil {
			return c.Right, true
		} else if c.Right == nil {
			return c.Left, true
		}
	}
	return nil, false
}
