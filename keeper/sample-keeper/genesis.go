package keeper

import (
	"context"

	"github.com/alice/checkers"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *checkers.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	// for _, counter := range data.Counters {
	// 	if err := k.Counter.Set(ctx, counter.Address, counter.Count); err != nil {
	// 		return err
	// 	}
	// }

	for _, indexedStoredGame := range data.IndexedStoredGameList {
		if err := k.StoredGames.Set(ctx, indexedStoredGame.Index, indexedStoredGame.StoredGame); err != nil {
		    return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*checkers.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	// var counters []checkers.Counter
	// if err := k.Counter.Walk(ctx, nil, func(address string, count uint64) (bool, error) {
	// 	counters = append(counters, checkers.Counter{
	// 		Address: address,
	// 		Count:   count,
	// 	})

	// 	return false, nil
	// }); err != nil {
	// 	return nil, err
	// }

	var indexedStoredGames []checkers.IndexedStoredGame
        if err := k.StoredGames.Walk(ctx, nil, func(index string, storedGame checkers.StoredGame) (bool, error) {
			indexedStoredGames = append(indexedStoredGames, checkers.IndexedStoredGame{
				Index:      index,
				StoredGame: storedGame,
			})
			return false, nil
        }); err != nil {
            return nil, err
        }


		return &checkers.GenesisState{
			Params:   params,
			// Counters: counters,
			IndexedStoredGameList: indexedStoredGames,
		}, nil
}
