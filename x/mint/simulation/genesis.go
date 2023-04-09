package simulation

// DONTCOVER

import (
	"encoding/json"
	//#nosec
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/furynet/fury/app/params"
	"github.com/furynet/fury/x/mint/types"
)

// Simulation parameter constants
const (
	Inflation = "inflation"
)

// GenInflation randomized Inflation
func GenInflation(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(int64(r.Intn(99))).QuoInt64(100)
}

// GenBlocksPerYear randomized BlocksPerYear
func GenBlocksPerYear(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(types.BlocksPerYear)
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var inflation sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, Inflation, &inflation, simState.Rand,
		func(r *rand.Rand) { inflation = GenInflation(r) },
	)

	// params
	mintDenom := params.DefaultBondDenom
	params := types.NewParams(mintDenom, types.BlocksPerYear, types.DefaultExcludeAmount, types.DefaultParams().Phases)

	mintGenesis := types.NewGenesisState(types.InitialMinter(inflation), params)

	_, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
