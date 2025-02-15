package testutil

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/params"
)

func TestSetupInitialDeposits(t *testing.T) {
	entries := 1
	deposits, depositDataRoots, privKeys := SetupInitialDeposits(t, uint64(entries))
	if len(deposits) != entries {
		t.Fatalf("incorrect number of deposits returned, wanted %d but received %d", entries, len(deposits))
	}
	if len(privKeys) != entries {
		t.Fatalf("incorrect number of private keys returned, wanted %d but received %d", entries, len(privKeys))
	}
	expectedPublicKeyAt0 := []byte{0xa9, 0x9a, 0x76, 0xed, 0x77, 0x96, 0xf7, 0xbe, 0x22, 0xd5, 0xb7, 0xe8, 0x5d, 0xee, 0xb7, 0xc5, 0x67, 0x7e, 0x88, 0xe5, 0x11, 0xe0, 0xb3, 0x37, 0x61, 0x8f, 0x8c, 0x4e, 0xb6, 0x13, 0x49, 0xb4, 0xbf, 0x2d, 0x15, 0x3f, 0x64, 0x9f, 0x7b, 0x53, 0x35, 0x9f, 0xe8, 0xb9, 0x4a, 0x38, 0xe4, 0x4c}
	if !bytes.Equal(deposits[0].Data.PublicKey, expectedPublicKeyAt0) {
		t.Fatalf("incorrect public key, wanted %x but received %x", expectedPublicKeyAt0, deposits[0].Data.PublicKey)
	}
	expectedWithdrawalCredentialsAt0 := []byte{0x00, 0xec, 0x7e, 0xf7, 0x78, 0x0c, 0x9d, 0x15, 0x15, 0x97, 0x92, 0x40, 0x36, 0x26, 0x2d, 0xd2, 0x8d, 0xc6, 0x0e, 0x12, 0x28, 0xf4, 0xda, 0x6f, 0xec, 0xf9, 0xd4, 0x02, 0xcb, 0x3f, 0x35, 0x94}
	if !bytes.Equal(deposits[0].Data.WithdrawalCredentials, expectedWithdrawalCredentialsAt0) {
		t.Fatalf("incorrect withdrawal credentials, wanted %x but received %x", expectedWithdrawalCredentialsAt0, deposits[0].Data.WithdrawalCredentials)
	}
	expectedDepositDataRootAt0 := []byte{0xc2, 0x58, 0x8b, 0xb0, 0x44, 0xf5, 0xe8, 0xaf, 0xb9, 0xb1, 0xcc, 0xb7, 0xe0, 0x83, 0x30, 0x35, 0x83, 0x18, 0xf2, 0x56, 0x27, 0x96, 0xfa, 0xad, 0xce, 0x92, 0x03, 0x50, 0x64, 0xaa, 0xf1, 0x3d}
	if !bytes.Equal(depositDataRoots[0][:], expectedDepositDataRootAt0) {
		t.Fatalf("incorrect deposit data root, wanted %x but received %x", expectedDepositDataRootAt0, depositDataRoots[0])
	}
	expectedSignatureAt0 := []byte{0xb3, 0xb9, 0x6e, 0xba, 0x50, 0xfa, 0x47, 0x49, 0x26, 0xfa, 0x46, 0xbb, 0xea, 0x3c, 0x8c, 0x73, 0x4c, 0x85, 0xc9, 0x70, 0x4e, 0x54, 0xb7, 0x19, 0xe5, 0x4e, 0x1b, 0xc5, 0x83, 0x77, 0xdd, 0x00, 0x30, 0x0b, 0x9e, 0xe4, 0xb0, 0x5b, 0xb2, 0x7b, 0x81, 0x8b, 0x38, 0xeb, 0xa2, 0x89, 0xcb, 0xe0, 0x06, 0x7a, 0x34, 0x56, 0xbc, 0xb8, 0xad, 0x59, 0xd0, 0x17, 0xfc, 0xf0, 0x04, 0xe5, 0xf1, 0xc5, 0xff, 0x1b, 0xf2, 0xe4, 0x89, 0x6b, 0x53, 0x2f, 0x4a, 0xea, 0x4b, 0x4c, 0x47, 0x06, 0x9a, 0x26, 0xe3, 0x85, 0x98, 0xf3, 0xd3, 0x37, 0x04, 0x7b, 0x8d, 0x0b, 0xd5, 0x25, 0xe4, 0x9f, 0xfc, 0xd2}
	if !bytes.Equal(deposits[0].Data.Signature, expectedSignatureAt0) {
		t.Fatalf("incorrect signature, wanted %x but received %x", expectedSignatureAt0, deposits[0].Data.Signature)
	}

	entries = 1024
	deposits, depositDataRoots, privKeys = SetupInitialDeposits(t, uint64(entries))
	if len(deposits) != entries {
		t.Fatalf("incorrect number of deposits returned, wanted %d but received %d", entries, len(deposits))
	}
	if len(privKeys) != entries {
		t.Fatalf("incorrect number of private keys returned, wanted %d but received %d", entries, len(privKeys))
	}
	// Ensure 0  has not changed
	if !bytes.Equal(deposits[0].Data.PublicKey, expectedPublicKeyAt0) {
		t.Fatalf("incorrect public key, wanted %x but received %x", expectedPublicKeyAt0, deposits[0].Data.PublicKey)
	}
	if !bytes.Equal(deposits[0].Data.WithdrawalCredentials, expectedWithdrawalCredentialsAt0) {
		t.Fatalf("incorrect withdrawal credentials, wanted %x but received %x", expectedWithdrawalCredentialsAt0, deposits[0].Data.WithdrawalCredentials)
	}
	if !bytes.Equal(depositDataRoots[0][:], expectedDepositDataRootAt0) {
		t.Fatalf("incorrect deposit data root, wanted %x but received %x", expectedDepositDataRootAt0, depositDataRoots[0])
	}
	if !bytes.Equal(deposits[0].Data.Signature, expectedSignatureAt0) {
		t.Fatalf("incorrect signature, wanted %x but received %x", expectedSignatureAt0, deposits[0].Data.Signature)
	}
	expectedPublicKeyAt1023 := []byte{0x81, 0x2b, 0x93, 0x5e, 0xc8, 0x4b, 0x0e, 0x9a, 0x83, 0x95, 0x55, 0xaf, 0x33, 0x60, 0xca, 0xfb, 0x83, 0x1b, 0xd6, 0x12, 0xcf, 0xa2, 0x2e, 0x25, 0xea, 0xb0, 0x3c, 0xf5, 0xfd, 0xb0, 0x2a, 0xf5, 0x2b, 0xa4, 0x01, 0x7a, 0xee, 0xa8, 0x8a, 0x2f, 0x62, 0x2c, 0x78, 0x6e, 0x7f, 0x47, 0x6f, 0x4b}
	if !bytes.Equal(deposits[1023].Data.PublicKey, expectedPublicKeyAt1023) {
		t.Fatalf("incorrect public key, wanted %x but received %x", expectedPublicKeyAt1023, deposits[1023].Data.PublicKey)
	}
	expectedWithdrawalCredentialsAt1023 := []byte{0x00, 0x23, 0xd5, 0x76, 0xbc, 0x6c, 0x15, 0xdb, 0xc4, 0x34, 0x70, 0x1f, 0x3f, 0x41, 0xfd, 0x3e, 0x67, 0x59, 0xd2, 0xea, 0x7c, 0xdc, 0x64, 0x71, 0x0e, 0xe2, 0x8d, 0xde, 0xf7, 0xd2, 0xda, 0x28}
	if !bytes.Equal(deposits[1023].Data.WithdrawalCredentials, expectedWithdrawalCredentialsAt1023) {
		t.Fatalf("incorrect withdrawal credentials, wanted %x but received %x", expectedWithdrawalCredentialsAt1023, deposits[1023].Data.WithdrawalCredentials)
	}
	expectedDepositDataRootAt1023 := []byte{0x54, 0x45, 0x80, 0xf3, 0xc3, 0x87, 0xdd, 0xfb, 0x1f, 0xf7, 0x03, 0xab, 0x15, 0xc9, 0x5b, 0x56, 0x2c, 0x29, 0x04, 0x7b, 0x17, 0xb4, 0xa0, 0x19, 0x69, 0xd6, 0x45, 0x7d, 0xec, 0x4e, 0x87, 0xfc}
	if !bytes.Equal(depositDataRoots[1023][:], expectedDepositDataRootAt1023) {
		t.Fatalf("incorrect deposit data root, wanted %x but received %x", expectedDepositDataRootAt1023, depositDataRoots[1023])
	}
	expectedSignatureAt1023 := []byte{0xa2, 0xad, 0x23, 0x3b, 0x6d, 0xa0, 0xd9, 0xf8, 0xb4, 0xac, 0xe0, 0xc9, 0xae, 0x25, 0x81, 0xfb, 0xca, 0x2d, 0x0a, 0xed, 0x6a, 0xdc, 0xd6, 0xda, 0x49, 0x0a, 0x75, 0xab, 0x3a, 0x3c, 0xc6, 0x37, 0xec, 0x65, 0xe3, 0x3d, 0xbc, 0x00, 0xad, 0xd8, 0x5f, 0x1e, 0x7b, 0x93, 0xcd, 0x63, 0x74, 0x8e, 0x0c, 0x28, 0x60, 0x4f, 0x99, 0x33, 0x6a, 0x29, 0x21, 0x57, 0xb6, 0xe0, 0x45, 0x9f, 0xaa, 0x10, 0xe9, 0x78, 0x02, 0x01, 0x68, 0x65, 0xcf, 0x6a, 0x4c, 0x2a, 0xd5, 0x5f, 0x37, 0xa1, 0x66, 0x05, 0x2b, 0x55, 0x86, 0xe7, 0x68, 0xb7, 0xfd, 0x76, 0xd5, 0x91, 0x3e, 0xeb, 0x6e, 0x46, 0x3f, 0x6d}
	if !bytes.Equal(deposits[1023].Data.Signature, expectedSignatureAt1023) {
		t.Fatalf("incorrect signature, wanted %x but received %x", expectedSignatureAt1023, deposits[1023].Data.Signature)
	}
}

func TestSignBlock(t *testing.T) {
	deposits, _, privKeys := SetupInitialDeposits(t, 100)
	validators := make([]*ethpb.Validator, len(deposits))
	for i := 0; i < len(validators); i++ {
		validators[i] = &ethpb.Validator{
			ExitEpoch: params.BeaconConfig().FarFutureEpoch,
			PublicKey: privKeys[i].PublicKey().Marshal()[:],
		}
	}

	beaconState := &pb.BeaconState{
		Slot: 0,
		Fork: &pb.Fork{
			CurrentVersion:  params.BeaconConfig().GenesisForkVersion,
			PreviousVersion: params.BeaconConfig().GenesisForkVersion,
		},
		Validators:       validators,
		RandaoMixes:      make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
		ActiveIndexRoots: make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
	}

	block := &ethpb.BeaconBlock{
		Slot:       0,
		ParentRoot: []byte{0xC0},
	}
	proposerIdx, err := helpers.BeaconProposerIndex(beaconState)
	if err != nil {
		t.Error(err)
	}
	signingRoot, err := ssz.SigningRoot(block)
	if err != nil {
		t.Error(err)
	}
	epoch := helpers.SlotToEpoch(block.Slot)
	domain := helpers.Domain(beaconState.Fork, epoch, params.BeaconConfig().DomainBeaconProposer)
	blockSig := privKeys[proposerIdx].Sign(signingRoot[:], domain).Marshal()

	signedBlock, err := SignBlock(beaconState, block, privKeys)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(blockSig[:], signedBlock.Signature) {
		t.Errorf("Expected block signatures to be equal, received %#x != %#x", blockSig[:], signedBlock.Signature)
	}
}

func TestCreateRandaoReveal(t *testing.T) {
	deposits, _, privKeys := SetupInitialDeposits(t, 100)
	validators := make([]*ethpb.Validator, len(deposits))
	for i := 0; i < len(validators); i++ {
		validators[i] = &ethpb.Validator{
			ExitEpoch: params.BeaconConfig().FarFutureEpoch,
			PublicKey: privKeys[i].PublicKey().Marshal()[:],
		}
	}

	beaconState := &pb.BeaconState{
		Slot: 0,
		Fork: &pb.Fork{
			CurrentVersion:  params.BeaconConfig().GenesisForkVersion,
			PreviousVersion: params.BeaconConfig().GenesisForkVersion,
		},
		Validators:       validators,
		RandaoMixes:      make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
		ActiveIndexRoots: make([][]byte, params.BeaconConfig().EpochsPerHistoricalVector),
	}

	epoch := helpers.CurrentEpoch(beaconState)
	randaoReveal, err := CreateRandaoReveal(beaconState, epoch, privKeys)
	if err != nil {
		t.Error(err)
	}

	proposerIdx, err := helpers.BeaconProposerIndex(beaconState)
	if err != nil {
		t.Error(err)
	}
	buf := make([]byte, 32)
	binary.LittleEndian.PutUint64(buf, epoch)
	domain := helpers.Domain(beaconState.Fork, epoch, params.BeaconConfig().DomainRandao)
	// We make the previous validator's index sign the message instead of the proposer.
	epochSignature := privKeys[proposerIdx].Sign(buf, domain).Marshal()

	if !bytes.Equal(randaoReveal[:], epochSignature[:]) {
		t.Errorf("Expected randao reveals to be equal, received %#x != %#x", randaoReveal[:], epochSignature[:])
	}
}
