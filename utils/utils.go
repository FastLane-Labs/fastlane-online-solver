package utils

import (
	"fmt"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	opsRelayUtils "github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/FastLane-Labs/fastlane-online-solver/config"
	"github.com/FastLane-Labs/fastlane-online-solver/contract/fastlaneOnline"
	"github.com/ethereum/go-ethereum/common"
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

func UserOpHash(userOperation fastlaneOnline.UserOperation, domain *apitypes.TypedDataDomain) (common.Hash, error) {
	opsRelayUserOp := &operation.UserOperation{
		From:         userOperation.From,
		To:           userOperation.To,
		Value:        userOperation.Value,
		Gas:          userOperation.Gas,
		MaxFeePerGas: userOperation.MaxFeePerGas,
		Nonce:        userOperation.Nonce,
		Deadline:     userOperation.Deadline,
		Dapp:         userOperation.Dapp,
		Control:      userOperation.Control,
		CallConfig:   userOperation.CallConfig,
		SessionKey:   userOperation.SessionKey,
		Data:         userOperation.Data,
		Signature:    userOperation.Signature,
	}

	userOpHash, relayErr := opsRelayUserOp.Hash(opsRelayUtils.FlagTrustedOpHash(opsRelayUserOp.CallConfig), domain)
	if relayErr != nil {
		return common.Hash{}, fmt.Errorf("failed to hash user operation: %v", relayErr)
	}
	return userOpHash, nil
}
