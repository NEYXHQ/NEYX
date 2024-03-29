package token_test

import (
	"testing"

	keepertest "github.com/Palo_Alt0/vanilla-NEYX/testutil/keeper"
	"github.com/Palo_Alt0/vanilla-NEYX/testutil/nullify"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token"
	"github.com/Palo_Alt0/vanilla-NEYX/x/token/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TokenList: []types.Token{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TokenCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenKeeper(t)
	token.InitGenesis(ctx, *k, genesisState)
	got := token.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	require.Equal(t, genesisState.TokenCount, got.TokenCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
