package market

import (
	//#nosec
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/furynet/fury/testutil/sample"
	marketsimulation "github.com/furynet/fury/x/market/simulation"
	"github.com/furynet/fury/x/market/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = marketsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	//#nosec
	opWeightMsgAddMarket          = "op_weight_msg_create_chain"
	defaultWeightMsgAddMarket int = 100

	//#nosec
	opWeightMsgResolveMarket          = "op_weight_msg_create_chain"
	defaultWeightMsgResolveMarket int = 100

	//#nosec
	opWeightMsgUpdateMarket          = "op_weight_msg_create_chain"
	defaultWeightMsgUpdateMarket int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MarketList: []types.Market{
			{
				Creator: sample.AccAddress(),
				UID:     "0",
			},
			{
				Creator: sample.AccAddress(),
				UID:     "1",
			},
		},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddMarket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddMarket, &weightMsgAddMarket, nil,
		func(_ *rand.Rand) {
			weightMsgAddMarket = defaultWeightMsgAddMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddMarket,
		marketsimulation.SimulateMsgAddMarket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgResolveMarket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgResolveMarket, &weightMsgResolveMarket, nil,
		func(_ *rand.Rand) {
			weightMsgResolveMarket = defaultWeightMsgResolveMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgResolveMarket,
		marketsimulation.SimulateMsgResolveMarket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMarket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMarket, &weightMsgUpdateMarket, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMarket = defaultWeightMsgUpdateMarket
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMarket,
		marketsimulation.SimulateMsgUpdateMarket(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}
