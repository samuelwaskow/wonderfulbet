package wonderfulbet_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "wonderfulbet/testutil/keeper"
	"wonderfulbet/testutil/nullify"
	"wonderfulbet/x/wonderfulbet"
	"wonderfulbet/x/wonderfulbet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WonderfulbetKeeper(t)
	wonderfulbet.InitGenesis(ctx, *k, genesisState)
	got := wonderfulbet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
