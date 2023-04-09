package keeper_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/cast"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simappUtil "github.com/furynet/xfury/testutil/simapp"
	"github.com/furynet/xfury/x/bet/keeper"
	"github.com/furynet/xfury/x/bet/types"
	marketkeeper "github.com/furynet/xfury/x/market/keeper"
	markettypes "github.com/furynet/xfury/x/market/types"
	"github.com/stretchr/testify/require"
)

var (
	testMarketUID  = "5db09053-2901-4110-8fb5-c14e21f8d555"
	testOddsUID1   = "6db09053-2901-4110-8fb5-c14e21f8d666"
	testOddsUID2   = "5e31c60f-2025-48ce-ae79-1dc110f16358"
	testOddsUID3   = "6e31c60f-2025-48ce-ae79-1dc110f16354"
	testMarketOdds = []*markettypes.Odds{
		{UID: testOddsUID1, Meta: "Odds 1"},
		{UID: testOddsUID2, Meta: "Odds 2"},
		{UID: testOddsUID3, Meta: "Odds 3"},
	}
	testSelectedBetOdds = &types.BetOdds{
		UID:               testOddsUID1,
		MarketUID:         testMarketUID,
		Value:             "4.20",
		MaxLossMultiplier: sdk.MustNewDecFromStr("0.1"),
	}
	testCreator   string
	testBet       *types.MsgPlaceBet
	testAddMarket *markettypes.MsgAddMarket

	testMarket = markettypes.Market{
		UID:                    testMarketUID,
		Creator:                simappUtil.TestParamUsers["user1"].Address.String(),
		StartTS:                1111111111,
		EndTS:                  uint64(time.Now().Unix()) + 5000,
		Odds:                   testMarketOdds,
		Status:                 markettypes.MarketStatus_MARKET_STATUS_RESULT_DECLARED,
		SrContributionForHouse: sdk.NewInt(5),
	}
)

func setupKeeperAndApp(t testing.TB) (*simappUtil.TestApp, *keeper.KeeperTest, sdk.Context) {
	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	return tApp, &tApp.BetKeeper, ctx
}

func setupKeeper(t testing.TB) (*keeper.KeeperTest, sdk.Context) {
	_, k, ctx := setupKeeperAndApp(t)

	return k, ctx
}

func addTestMarket(t testing.TB, tApp *simappUtil.TestApp, ctx sdk.Context) {
	testCreator = simappUtil.TestParamUsers["user1"].Address.String()
	testAddMarketClaim := jwt.MapClaims{
		"uid":                       testMarketUID,
		"start_ts":                  1111111111,
		"end_ts":                    uint64(ctx.BlockTime().Unix()) + 1000,
		"odds":                      testMarketOdds,
		"exp":                       9999999999,
		"iat":                       7777777777,
		"meta":                      "Winner of x:y",
		"sr_contribution_for_house": sdk.NewInt(500000),
		"status":                    markettypes.MarketStatus_MARKET_STATUS_ACTIVE,
	}
	testAddMarketTicket, err := createJwtTicket(testAddMarketClaim)
	require.Nil(t, err)

	testAddMarket = &markettypes.MsgAddMarket{
		Creator: testCreator,
		Ticket:  testAddMarketTicket,
	}
	wctx := sdk.WrapSDKContext(ctx)
	marketSrv := marketkeeper.NewMsgServerImpl(tApp.MarketKeeper)
	resAddMarket, err := marketSrv.AddMarket(wctx, testAddMarket)
	require.Nil(t, err)
	require.NotNil(t, resAddMarket)
}

func addTestMarketBatch(t testing.TB, tApp *simappUtil.TestApp, ctx sdk.Context, count int) (uids []string) {
	for i := 0; i < count; i++ {
		testCreator = simappUtil.TestParamUsers["user"+cast.ToString(i)].Address.String()
		uid := uuid.NewString()
		uids = append(uids, uid)
		testAddMarketClaim := jwt.MapClaims{
			"uid":                       uid,
			"start_ts":                  1111111111,
			"end_ts":                    uint64(ctx.BlockTime().Unix()) + 1000,
			"odds":                      testMarketOdds,
			"exp":                       9999999999,
			"iat":                       7777777777,
			"meta":                      "Winner of x:y",
			"sr_contribution_for_house": sdk.NewInt(500000),
			"status":                    markettypes.MarketStatus_MARKET_STATUS_ACTIVE,
		}
		testAddMarketTicket, err := createJwtTicket(testAddMarketClaim)
		require.Nil(t, err)

		testAddMarket = &markettypes.MsgAddMarket{
			Creator: testCreator,
			Ticket:  testAddMarketTicket,
		}
		wctx := sdk.WrapSDKContext(ctx)
		marketSrv := marketkeeper.NewMsgServerImpl(tApp.MarketKeeper)
		resAddMarket, err := marketSrv.AddMarket(wctx, testAddMarket)
		require.Nil(t, err)
		require.NotNil(t, resAddMarket)
	}

	return uids
}

func placeTestBet(ctx sdk.Context, t testing.TB, tApp *simappUtil.TestApp, betUID string, selectedOdds *types.BetOdds) {
	testCreator = simappUtil.TestParamUsers["user1"].Address.String()
	wctx := sdk.WrapSDKContext(ctx)
	betSrv := keeper.NewMsgServerImpl(tApp.BetKeeper)
	testKyc := &types.KycDataPayload{
		Approved: true,
		ID:       testCreator,
	}

	if selectedOdds == nil {
		selectedOdds = testSelectedBetOdds
	}

	testPlaceBetClaim := jwt.MapClaims{
		"exp":           9999999999,
		"iat":           7777777777,
		"selected_odds": selectedOdds,
		"kyc_data":      testKyc,
		"odds_type":     1,
	}
	testPlaceBetTicket, err := createJwtTicket(testPlaceBetClaim)
	require.Nil(t, err)

	testBet = &types.MsgPlaceBet{
		Creator: testCreator,
		Bet: &types.PlaceBetFields{
			UID:    betUID,
			Amount: sdk.NewInt(500),
			Ticket: testPlaceBetTicket,
		},
	}
	resPlaceBet, err := betSrv.PlaceBet(wctx, testBet)
	require.Nil(t, err)
	require.NotNil(t, resPlaceBet)
}

func createJwtTicket(claim jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claim)
	return token.SignedString(simappUtil.TestDVMPrivateKeys[0])
}

func TestLogger(t *testing.T) {
	k, ctx := setupKeeper(t)
	logger := k.Logger(ctx)
	require.NotNil(t, logger)

	logger.Debug("test log")
}
