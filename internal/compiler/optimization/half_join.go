package optimization

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// OptHalfJoin reduces Joins that are of the form Join(any,nil) or Join(nil,any)
// to just any. The result is a command and a flag that determines whether the
// command has been optimized. If that flag is false, then a nil command is
// returned, and the command, that was passed in, should be used further.
// Otherwise, proceed to work with the returned command. The input command will
// not be modified.
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
		left := c.Left
		right := c.Right
		var optimized bool
		if left != nil {
			if optimizedLeft, ok := OptHalfJoin(left); ok {
				optimized = optimized || ok
				left = optimizedLeft.(command.List)
			}
		}
		if right != nil {
			if optimizedRight, ok := OptHalfJoin(right); ok {
				optimized = optimized || ok
				right = optimizedRight.(command.List)
			}
		}

		// both halfs are optimized, if one of them is nil, don't return a join,
		// but the non-nil part
		if left == nil && right == nil {
			return nil, true
		}
		if left == nil {
			return right, true
		}
		if right == nil {
			return left, true
		}

		// if none of the both halfs are nil return them in a join (both halfs
		// are potentially optimized)
		return command.Join{
			Left:  left,
			Right: right,
		}, optimized
	}
	return nil, false
}
