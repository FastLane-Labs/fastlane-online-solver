package utils

import (
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/fastlaneOnline"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func OpsRelaySolverOpToFastlaneOnlineSolverOp(op *operation.SolverOperation) fastlaneOnline.SolverOperation {

	return fastlaneOnline.SolverOperation{
		From:         op.From,
		To:           op.To,
		Value:        op.Value,
		Gas:          op.Gas,
		MaxFeePerGas: op.MaxFeePerGas,
		Deadline:     op.Deadline,
		Solver:       op.Solver,
		Control:      op.Control,
		UserOpHash:   op.UserOpHash,
		BidToken:     op.BidToken,
		BidAmount:    op.BidAmount,
		Data:         op.Data,
		Signature:    op.Signature,
	}
}

func Domain(conf *config.Config, chainId int64) *apitypes.TypedDataDomain {
	return &apitypes.TypedDataDomain{
		Name:              "AtlasVerification",
		Version:           "1.0",
		ChainId:           math.NewHexOrDecimal256(chainId),
		VerifyingContract: conf.AtlasVerificationAddress,
	}
}
