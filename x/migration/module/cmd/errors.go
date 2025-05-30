package cmd

import sdkerrors "cosmossdk.io/errors"

const codespace = "pocketd/migrate"

var (
	// ErrInvalidUsage usage is returned when the CLI arguments are invalid.
	ErrInvalidUsage = sdkerrors.Register(codespace, 1100, "invalid CLI usage")

	// ErrMorseExportState is returned with the JSON generated from `pocket util export-genesis-for-reset` is invalid.
	ErrMorseExportState = sdkerrors.Register(codespace, 1101, "morse export state failed")

	// ErrMorseStateTransform is returned upon general failure when transforming the MorseExportState into the MorseAccountState.
	ErrMorseStateTransform = sdkerrors.Register(codespace, 1102, "morse export to state transformation invalid")

	// ErrInvalidMorseAccountState is used when the morse account state hash doesn't match the expected hash.
	ErrInvalidMorseAccountState = sdkerrors.Register(codespace, 1103, "morse account state is invalid")
)
