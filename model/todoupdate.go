package model

import "github.com/guregu/null"

// TodoUpdate holds input data for updating ToDos
type TodoUpdate struct {
	TodoNew

	// special Bool type here to treat an absent "Completed" as null and not false (by bool default)
	MaybeCompleted null.Bool `json:"Completed"`
}
