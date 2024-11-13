package checkers

import "cosmossdk.io/collections"

const ModuleName = "checkers"
const MaxIndexLength = 256

var (
	// ParamsKey  = collections.NewPrefix(0)
	CounterKey = collections.NewPrefix(1)
	ParamsKey      = collections.NewPrefix("Params")
	StoredGamesKey = collections.NewPrefix("StoredGames/value/")
)


