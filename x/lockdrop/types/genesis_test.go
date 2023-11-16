package types_test

import (
	"testing"

	proto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/dymensionxyz/dymension/x/lockdrop/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGenesisStateMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		state *types.GenesisState
	}{
		{ // default genesis state
			state: types.DefaultGenesisState(),
		},
		{ // empty lock durations
			state: &types.GenesisState{
				Params: types.DefaultParams(),
				DistrInfo: &types.DistrInfo{
					TotalWeight: sdk.ZeroInt(),
					Records:     nil,
				},
			},
		},
		{ // empty array distribution info
			state: &types.GenesisState{
				Params: types.DefaultParams(),
				DistrInfo: &types.DistrInfo{
					TotalWeight: sdk.ZeroInt(),
					Records:     nil,
				},
			},
		},
		{ // one record distribution info
			state: &types.GenesisState{
				Params: types.DefaultParams(),
				DistrInfo: &types.DistrInfo{
					TotalWeight: sdk.NewInt(1),
					Records: []types.DistrRecord{
						{
							GaugeId: 1,
							Weight:  sdk.NewInt(1),
						},
					},
				},
			},
		},
		{ // empty params
			state: &types.GenesisState{
				Params: types.Params{},
				DistrInfo: &types.DistrInfo{
					TotalWeight: sdk.ZeroInt(),
					Records:     nil,
				},
			},
		},
	}

	for _, test := range tests {
		bz, err := proto.Marshal(test.state)
		require.NoError(t, err)
		decoded := types.GenesisState{}
		err = proto.Unmarshal(bz, &decoded)
		require.NoError(t, err)
		require.Equal(t, *test.state, decoded)
	}
}