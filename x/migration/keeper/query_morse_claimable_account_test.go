package keeper_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/nullify"
	"github.com/pokt-network/poktroll/x/migration/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMorseClaimableAccountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MigrationKeeper(t)
	msgs := createNMorseClaimableAccount(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryMorseClaimableAccountRequest
		response *types.QueryMorseClaimableAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryMorseClaimableAccountRequest{
				Address: msgs[0].MorseSrcAddress,
			},
			response: &types.QueryMorseClaimableAccountResponse{MorseClaimableAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryMorseClaimableAccountRequest{
				Address: msgs[1].MorseSrcAddress,
			},
			response: &types.QueryMorseClaimableAccountResponse{MorseClaimableAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryMorseClaimableAccountRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
		{
			desc: "Uppercase Address",
			request: &types.QueryMorseClaimableAccountRequest{
				Address: strings.ToUpper(msgs[0].MorseSrcAddress),
			},
			response: &types.QueryMorseClaimableAccountResponse{MorseClaimableAccount: msgs[0]},
		},
		{
			desc: "Lowercase Address",
			request: &types.QueryMorseClaimableAccountRequest{
				Address: strings.ToLower(msgs[0].MorseSrcAddress),
			},
			response: &types.QueryMorseClaimableAccountResponse{MorseClaimableAccount: msgs[0]},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MorseClaimableAccount(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMorseClaimableAccountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MigrationKeeper(t)
	msgs := createNMorseClaimableAccount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMorseClaimableAccountRequest {
		return &types.QueryAllMorseClaimableAccountRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MorseClaimableAccountAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MorseClaimableAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MorseClaimableAccount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MorseClaimableAccountAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MorseClaimableAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MorseClaimableAccount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MorseClaimableAccountAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MorseClaimableAccount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MorseClaimableAccountAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
