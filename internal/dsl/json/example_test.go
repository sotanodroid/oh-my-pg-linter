package json_test

import (
	"log"

	"github.com/ozonru/oh-my-pg-linter/internal/dsl/inspect"
	"github.com/ozonru/oh-my-pg-linter/internal/dsl/json"
	lua "github.com/yuin/gopher-lua"
)

// json.decode(string).
func ExampleDecode() {
	state := lua.NewState()
	json.Preload(state)
	inspect.Preload(state)
	source := `
    local json = require("json")
    local inspect = require("inspect")
    local jsonString = [[{"a":{"b":1}}]]
    local result, err = json.decode(jsonString)
    if err then error(err) end
    print(inspect(result, {newline="", indent=""}))
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// {a = {b = 1}}
}

// json.encode(obj).
func ExampleEncode() {
	state := lua.NewState()
	json.Preload(state)
	inspect.Preload(state)
	source := `
    local json = require("json")
    local inspect = require("inspect")
    local table = {a={b=1}}
    local result, err = json.encode(table)
    if err then error(err) end
    print(inspect(result, {newline="", indent=""}))

	print(inspect( json.encode( {} ) ))
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// '{"a":{"b":1}}'
	// "[]"
}
