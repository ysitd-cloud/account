package modals

import "fmt"

type IncorrectResultAffectedError struct {
	Row      int64
	Expected int
}

func (e *IncorrectResultAffectedError) Error() string {
	return fmt.Sprintf("Database result row does not match, excepted: %d, affected: %d", e.Expected, e.Row)
}

func (e *IncorrectResultAffectedError) RuntimeError() {}
