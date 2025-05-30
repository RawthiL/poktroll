package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultNumBlocksPerSession                = 10
	ParamNumBlocksPerSession                  = "num_blocks_per_session"
	DefaultGracePeriodEndOffsetBlocks         = 1
	ParamGracePeriodEndOffsetBlocks           = "grace_period_end_offset_blocks"
	DefaultClaimWindowOpenOffsetBlocks        = 1
	ParamClaimWindowOpenOffsetBlocks          = "claim_window_open_offset_blocks"
	DefaultClaimWindowCloseOffsetBlocks       = 4
	ParamClaimWindowCloseOffsetBlocks         = "claim_window_close_offset_blocks"
	DefaultProofWindowOpenOffsetBlocks        = 0
	ParamProofWindowOpenOffsetBlocks          = "proof_window_open_offset_blocks"
	DefaultProofWindowCloseOffsetBlocks       = 4
	ParamProofWindowCloseOffsetBlocks         = "proof_window_close_offset_blocks"
	DefaultSupplierUnbondingPeriodSessions    = 1 // 1 session
	ParamSupplierUnbondingPeriodSessions      = "supplier_unbonding_period_sessions"
	DefaultApplicationUnbondingPeriodSessions = 1 // 1 session
	ParamApplicationUnbondingPeriodSessions   = "application_unbonding_period_sessions"
	DefaultGatewayUnbondingPeriodSessions     = 1 // 1 session
	ParamGatewayUnbondingPeriodSessions       = "gateway_unbonding_period_sessions"

	// TODO_MAINNET_MIGRATION(@olshansk): Determine the default value.
	// The default compute unit cost multiplier in pPOKT (i.e. 1/1e6 uPOKT)
	DefaultComputeUnitsToTokensMultiplier = 42_000_000
	ParamComputeUnitsToTokensMultiplier   = "compute_units_to_tokens_multiplier"

	// The default compute unit cost granularity corresponding to 1pPOKT (i.e. 1/1e6 uPOKT)
	DefaultComputeUnitCostGranularity = 1_000_000
	ParamComputeUnitCostGranularity   = "compute_unit_cost_granularity"
)

var (
	_                                     paramtypes.ParamSet = (*Params)(nil)
	KeyNumBlocksPerSession                                    = []byte("NumBlocksPerSession")
	KeyGracePeriodEndOffsetBlocks                             = []byte("GracePeriodEndOffsetBlocks")
	KeyClaimWindowOpenOffsetBlocks                            = []byte("ClaimWindowOpenOffsetBlocks")
	KeyClaimWindowCloseOffsetBlocks                           = []byte("ClaimWindowCloseOffsetBlocks")
	KeyProofWindowOpenOffsetBlocks                            = []byte("ProofWindowOpenOffsetBlocks")
	KeyProofWindowCloseOffsetBlocks                           = []byte("ProofWindowCloseOffsetBlocks")
	KeySupplierUnbondingPeriodSessions                        = []byte("SupplierUnbondingPeriodSessions")
	KeyApplicationUnbondingPeriodSessions                     = []byte("ApplicationUnbondingPeriodSessions")
	KeyGatewayUnbondingPeriodSessions                         = []byte("GatewayUnbondingPeriodSessions")
	KeyComputeUnitsToTokensMultiplier                         = []byte("ComputeUnitsToTokensMultiplier")
	KeyComputeUnitCostGranularity                             = []byte("ComputeUnitCostGranularity")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		NumBlocksPerSession:                DefaultNumBlocksPerSession,
		ClaimWindowOpenOffsetBlocks:        DefaultClaimWindowOpenOffsetBlocks,
		ClaimWindowCloseOffsetBlocks:       DefaultClaimWindowCloseOffsetBlocks,
		ProofWindowOpenOffsetBlocks:        DefaultProofWindowOpenOffsetBlocks,
		ProofWindowCloseOffsetBlocks:       DefaultProofWindowCloseOffsetBlocks,
		GracePeriodEndOffsetBlocks:         DefaultGracePeriodEndOffsetBlocks,
		SupplierUnbondingPeriodSessions:    DefaultSupplierUnbondingPeriodSessions,
		ApplicationUnbondingPeriodSessions: DefaultApplicationUnbondingPeriodSessions,
		GatewayUnbondingPeriodSessions:     DefaultGatewayUnbondingPeriodSessions,
		ComputeUnitsToTokensMultiplier:     DefaultComputeUnitsToTokensMultiplier,
		ComputeUnitCostGranularity:         DefaultComputeUnitCostGranularity,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (params *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			KeyNumBlocksPerSession,
			&params.NumBlocksPerSession,
			ValidateNumBlocksPerSession,
		),
		paramtypes.NewParamSetPair(
			KeyGracePeriodEndOffsetBlocks,
			&params.GracePeriodEndOffsetBlocks,
			ValidateGracePeriodEndOffsetBlocks,
		),
		paramtypes.NewParamSetPair(
			KeyClaimWindowOpenOffsetBlocks,
			&params.ClaimWindowOpenOffsetBlocks,
			ValidateClaimWindowOpenOffsetBlocks,
		),
		paramtypes.NewParamSetPair(
			KeyClaimWindowCloseOffsetBlocks,
			&params.ClaimWindowCloseOffsetBlocks,
			ValidateClaimWindowCloseOffsetBlocks,
		),
		paramtypes.NewParamSetPair(
			KeyProofWindowOpenOffsetBlocks,
			&params.ProofWindowOpenOffsetBlocks,
			ValidateProofWindowOpenOffsetBlocks,
		),
		paramtypes.NewParamSetPair(
			KeyProofWindowCloseOffsetBlocks,
			&params.ProofWindowCloseOffsetBlocks,
			ValidateProofWindowCloseOffsetBlocks,
		),
		paramtypes.NewParamSetPair(
			KeySupplierUnbondingPeriodSessions,
			&params.SupplierUnbondingPeriodSessions,
			ValidateSupplierUnbondingPeriodSessions,
		),
		paramtypes.NewParamSetPair(
			KeyApplicationUnbondingPeriodSessions,
			&params.ApplicationUnbondingPeriodSessions,
			ValidateApplicationUnbondingPeriodSessions,
		),
		paramtypes.NewParamSetPair(
			KeyGatewayUnbondingPeriodSessions,
			&params.GatewayUnbondingPeriodSessions,
			ValidateGatewayUnbondingPeriodSessions,
		),
		paramtypes.NewParamSetPair(
			KeyComputeUnitsToTokensMultiplier,
			&params.ComputeUnitsToTokensMultiplier,
			ValidateComputeUnitsToTokensMultiplier,
		),
	}
}

// Validate validates the set of params
func (params *Params) ValidateBasic() error {
	if err := ValidateNumBlocksPerSession(params.NumBlocksPerSession); err != nil {
		return err
	}

	if err := ValidateClaimWindowOpenOffsetBlocks(params.ClaimWindowOpenOffsetBlocks); err != nil {
		return err
	}

	if err := ValidateClaimWindowCloseOffsetBlocks(params.ClaimWindowCloseOffsetBlocks); err != nil {
		return err
	}

	if err := ValidateProofWindowOpenOffsetBlocks(params.ProofWindowOpenOffsetBlocks); err != nil {
		return err
	}

	if err := ValidateProofWindowCloseOffsetBlocks(params.ProofWindowCloseOffsetBlocks); err != nil {
		return err
	}

	if err := ValidateGracePeriodEndOffsetBlocks(params.GracePeriodEndOffsetBlocks); err != nil {
		return err
	}

	if err := ValidateSupplierUnbondingPeriodSessions(params.SupplierUnbondingPeriodSessions); err != nil {
		return err
	}

	if err := ValidateApplicationUnbondingPeriodSessions(params.ApplicationUnbondingPeriodSessions); err != nil {
		return err
	}

	if err := ValidateGatewayUnbondingPeriodSessions(params.GatewayUnbondingPeriodSessions); err != nil {
		return err
	}

	if err := ValidateComputeUnitsToTokensMultiplier(params.ComputeUnitsToTokensMultiplier); err != nil {
		return err
	}

	if err := ValidateComputeUnitCostGranularity(params.ComputeUnitCostGranularity); err != nil {
		return err
	}

	if err := validateGracePeriodOffsetBlocksIsLessThanNumBlocksPerSession(params); err != nil {
		return err
	}

	if err := validateClaimWindowOpenOffsetIsAtLeastGracePeriodEndOffset(params); err != nil {
		return err
	}

	if err := validateSupplierUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks(params); err != nil {
		return err
	}

	if err := validateApplicationUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks(params); err != nil {
		return err
	}

	// TODO_MAINNET_MIGRATION(@bryanchriswhite): Add validation which ensures that
	// SessionEndToProofWindowCloseBlocks is a multiple of NumBlocksPerSession.

	return nil
}

// ValidateNumBlocksPerSession validates the NumBlocksPerSession param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateNumBlocksPerSession(numBlocksPerSessionAny any) error {
	numBlocksPerSession, err := validateIsUint64(numBlocksPerSessionAny)
	if err != nil {
		return err
	}

	if numBlocksPerSession < 1 {
		return ErrSharedParamInvalid.Wrapf("invalid NumBlocksPerSession: (%v)", numBlocksPerSession)
	}

	return nil
}

// ValidateClaimWindowOpenOffsetBlocks validates the ClaimWindowOpenOffsetBlocks param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateClaimWindowOpenOffsetBlocks(claimWindowOpenOffsetBlocksAny any) error {
	_, err := validateIsUint64(claimWindowOpenOffsetBlocksAny)
	return err
}

// ValidateClaimWindowCloseOffsetBlocks validates the ClaimWindowCloseOffsetBlocks param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateClaimWindowCloseOffsetBlocks(claimWindowCloseOffsetBlocksAny any) error {
	_, err := validateIsUint64(claimWindowCloseOffsetBlocksAny)
	return err
}

// ValidateProofWindowOpenOffsetBlocks validates the ProofWindowOpenOffsetBlocks param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateProofWindowOpenOffsetBlocks(proofWindowOpenOffsetBlocksAny any) error {
	_, err := validateIsUint64(proofWindowOpenOffsetBlocksAny)
	return err
}

// ValidateProofWindowCloseOffsetBlocks validates the ProofWindowCloseOffsetBlocks param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateProofWindowCloseOffsetBlocks(proofWindowCloseOffsetBlocksAny any) error {
	_, err := validateIsUint64(proofWindowCloseOffsetBlocksAny)
	return err
}

// ValidateGracePeriodEndOffsetBlocks validates the GracePeriodEndOffsetBlocks param
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateGracePeriodEndOffsetBlocks(gracePeriodEndOffsetBlocksAny any) error {
	_, err := validateIsUint64(gracePeriodEndOffsetBlocksAny)
	return err
}

// ValidateSupplierUnbondingPeriodSessions validates the SupplierUnbondingPeriodSessions
// governance parameter.
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateSupplierUnbondingPeriodSessions(supplierUnbondingPeriodSessionsAny any) error {
	supplierUnbondingPeriodSessions, err := validateIsUint64(supplierUnbondingPeriodSessionsAny)
	if err != nil {
		return err
	}

	if supplierUnbondingPeriodSessions < 1 {
		return ErrSharedParamInvalid.Wrapf("invalid SupplierUnbondingPeriodSessions: (%v)", supplierUnbondingPeriodSessions)
	}

	return nil
}

// ValidateApplicationUnbondingPeriodSessions validates the ApplicationUnbondingPeriodSessions
// governance parameter.
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateApplicationUnbondingPeriodSessions(applicationUnboindingPeriodSessionsAny any) error {
	applicationUnbondingPeriodSessions, err := validateIsUint64(applicationUnboindingPeriodSessionsAny)
	if err != nil {
		return err
	}

	if applicationUnbondingPeriodSessions < 1 {
		return ErrSharedParamInvalid.Wrapf("invalid ApplicationUnbondingPeriodSessions: (%v)", applicationUnbondingPeriodSessions)
	}

	return nil
}

// ValidateGatewayUnbondingPeriodSessions validates the GatewayUnbondingPeriodSessions
// governance parameter.
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateGatewayUnbondingPeriodSessions(gatewayUnboindingPeriodSessionsAny any) error {
	gatewayUnbondingPeriodSessions, err := validateIsUint64(gatewayUnboindingPeriodSessionsAny)
	if err != nil {
		return err
	}

	if gatewayUnbondingPeriodSessions < 1 {
		return ErrSharedParamInvalid.Wrapf("invalid GatewayUnbondingPeriodSessions: (%v)", gatewayUnbondingPeriodSessions)
	}

	return nil
}

// ValidateComputeUnitsToTokensMultiplier validates the ComputeUnitsToTokensMultiplier governance parameter.
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateComputeUnitsToTokensMultiplier(computeUnitsToTokensMultiplierAny any) error {
	ComputeUnitsToTokensMultiplier, ok := computeUnitsToTokensMultiplierAny.(uint64)
	if !ok {
		return ErrSharedParamInvalid.Wrapf("invalid parameter type: %T", computeUnitsToTokensMultiplierAny)
	}

	if ComputeUnitsToTokensMultiplier <= 0 {
		return ErrSharedParamInvalid.Wrapf("invalid ComputeUnitsToTokensMultiplier: (%v)", ComputeUnitsToTokensMultiplier)
	}

	return nil
}

// ValidateComputeUnitCostGranularity validates the ComputeUnitCostGranularity governance parameter.
// NB: The argument is an interface type to satisfy the ParamSetPair function signature.
func ValidateComputeUnitCostGranularity(computeUnitCostGranularityAny any) error {
	computeUnitCostGranularity, ok := computeUnitCostGranularityAny.(uint64)
	if !ok {
		return ErrSharedParamInvalid.Wrapf("invalid parameter type: %T", computeUnitCostGranularityAny)
	}

	if computeUnitCostGranularity <= 0 {
		return ErrSharedParamInvalid.Wrapf("invalid ComputeUnitCostGranularity: (%v)", computeUnitCostGranularity)
	}

	return nil
}

// validateIsUint64 returns the casted uin64 value or an error if value is not
// type assertable to uint64.
func validateIsUint64(value any) (uint64, error) {
	uint64Value, ok := value.(uint64)
	if !ok {
		return 0, ErrSharedParamInvalid.Wrapf("invalid parameter type: %T", value)
	}

	return uint64Value, nil
}

// validateClaimWindowOpenOffsetIsAtLeastGracePeriodEndOffset validates that the ClaimWindowOpenOffsetBlocks
// is at least as big as GracePeriodEndOffsetBlocks. The claim window cannot open until the grace period ends
// to ensure that the seed for the earliest supplier claim commit height is only observed after the last relay
// for a given session could be serviced.
func validateClaimWindowOpenOffsetIsAtLeastGracePeriodEndOffset(params *Params) error {
	if params.ClaimWindowOpenOffsetBlocks < params.GracePeriodEndOffsetBlocks {
		return ErrSharedParamInvalid.Wrapf(
			"ClaimWindowOpenOffsetBlocks (%v) must be at least GracePeriodEndOffsetBlocks (%v)",
			params.ClaimWindowOpenOffsetBlocks,
			params.GracePeriodEndOffsetBlocks,
		)
	}
	return nil
}

// validateGracePeriodOffsetBlocksIsLessThanNumBlocksPerSession validates that the
// GracePeriodEndOffsetBlocks is less than NumBlocksPerSession; i.e., one session.
func validateGracePeriodOffsetBlocksIsLessThanNumBlocksPerSession(params *Params) error {
	if params.GracePeriodEndOffsetBlocks >= params.NumBlocksPerSession {
		return ErrSharedParamInvalid.Wrapf(
			"GracePeriodEndOffsetBlocks (%v) must be less than NumBlocksPerSession (%v)",
			params.GracePeriodEndOffsetBlocks,
			params.NumBlocksPerSession,
		)
	}
	return nil
}

// validateSupplierUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks
// validates that the SupplierUnbondingPeriodSession blocks is greater than the cumulative
// proof window close blocks.
// It ensures that a supplier cannot unbond before the pending claims are settled.
func validateSupplierUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks(params *Params) error {
	cumulativeProofWindowCloseBlocks := GetSessionEndToProofWindowCloseBlocks(params)
	supplierUnbondingPeriodSessions := int64(params.SupplierUnbondingPeriodSessions * params.NumBlocksPerSession)

	if supplierUnbondingPeriodSessions < cumulativeProofWindowCloseBlocks {
		return ErrSharedParamInvalid.Wrapf(
			"SupplierUnbondingPeriodSessions (%v session) (%v blocks) must be greater than the cumulative ProofWindowCloseOffsetBlocks (%v)",
			params.SupplierUnbondingPeriodSessions,
			supplierUnbondingPeriodSessions,
			cumulativeProofWindowCloseBlocks,
		)
	}

	return nil
}

// validateApplicationUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks
// ensures that an application cannot unbond before the pending claims are settled.
func validateApplicationUnbondingPeriodIsGreaterThanCumulativeProofWindowCloseBlocks(params *Params) error {
	cumulativeProofWindowCloseBlocks := GetSessionEndToProofWindowCloseBlocks(params)
	applicationUnbondingPeriodSessions := int64(params.ApplicationUnbondingPeriodSessions * params.NumBlocksPerSession)

	if applicationUnbondingPeriodSessions < cumulativeProofWindowCloseBlocks {
		return ErrSharedParamInvalid.Wrapf(
			"ApplicationUnbondingPeriodSessions (%v session) (%v blocks) must be greater than the cumulative ProofWindowCloseOffsetBlocks (%v)",
			params.ApplicationUnbondingPeriodSessions,
			applicationUnbondingPeriodSessions,
			cumulativeProofWindowCloseBlocks,
		)
	}

	return nil
}
