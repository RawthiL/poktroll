package testsupplier

import (
	"context"
	"testing"

	"cosmossdk.io/depinject"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/supplier"
	"github.com/pokt-network/poktroll/pkg/client/tx"
	"github.com/pokt-network/poktroll/testutil/mockclient"
	"github.com/pokt-network/poktroll/testutil/testclient/testtx"
)

// NewLocalnetClient creates and returns a new supplier client that connects to
// the LocalNet validator.
func NewLocalnetClient(
	t *testing.T,
	signingKeyName string,
) client.SupplierClient {
	t.Helper()

	txClientOpt := tx.WithSigningKeyName(signingKeyName)
	supplierClientOpt := supplier.WithSigningKeyName(signingKeyName)

	txCtx := testtx.NewLocalnetContext(t)
	txClient := testtx.NewLocalnetClient(t, txClientOpt)

	deps := depinject.Supply(
		txCtx,
		txClient,
	)

	supplierClient, err := supplier.NewSupplierClient(deps, supplierClientOpt)
	require.NoError(t, err)
	return supplierClient
}

// NewClaimProofSupplierClientMap creates and returns a map of supplier to supplier
// client mocks. Each supplier client is expected to submit exactly 1 claim and
// exactly proofCount proofs.
func NewClaimProofSupplierClientMap(
	ctx context.Context,
	t *testing.T,
	supplierOperatorAddress string,
	proofCount int,
) *supplier.SupplierClientMap {
	t.Helper()

	ctrl := gomock.NewController(t)
	supplierClientMock := mockclient.NewMockSupplierClient(ctrl)

	supplierClientMock.EXPECT().
		OperatorAddress().
		Return(supplierOperatorAddress).
		AnyTimes()

	supplierClientMock.EXPECT().
		CreateClaims(
			gomock.Eq(ctx),
			gomock.Any(),
			gomock.AssignableToTypeOf(([]client.MsgCreateClaim)(nil)),
		).
		Return(nil).
		Times(1)

	supplierClientMock.EXPECT().
		SubmitProofs(
			gomock.Eq(ctx),
			gomock.Any(),
			gomock.AssignableToTypeOf(([]client.MsgSubmitProof)(nil)),
		).
		Return(nil).
		Times(proofCount)

	supplierClientMap := supplier.NewSupplierClientMap()
	supplierClientMap.SupplierClients[supplierOperatorAddress] = supplierClientMock

	return supplierClientMap
}
