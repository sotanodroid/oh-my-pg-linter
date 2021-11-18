package manager

import (
	"github.com/ozonru/oh-my-pg-linter/internal/dsl/filepath"
	"github.com/ozonru/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/ozonru/oh-my-pg-linter/internal/dsl/json"
	"github.com/ozonru/oh-my-pg-linter/internal/dsl/parser"
	lua "github.com/yuin/gopher-lua"
)

// NewState ...
func NewState() *lua.LState {
	state := lua.NewState()
	filepath.Preload(state)
	inspect.Preload(state)
	json.Preload(state)
	parser.Preload(state)
	return state
}
