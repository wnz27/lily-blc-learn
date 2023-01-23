// Code generated by: `make actors-gen`. DO NOT EDIT.
package verifreg

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lily/chain/actors/adt"

	"crypto/sha256"

	verifreg2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	adt2 "github.com/filecoin-project/specs-actors/v2/actors/util/adt"

	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	"github.com/filecoin-project/lily/chain/actors"
)

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type state2 struct {
	verifreg2.State
	store adt.Store
}

func (s *state2) ActorKey() string {
	return actors.VerifregKey
}

func (s *state2) ActorVersion() actors.Version {
	return actors.Version2
}

func (s *state2) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func (s *state2) VerifiedClientsMapBitWidth() int {

	return 5

}

func (s *state2) VerifiedClientsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state2) VerifiedClientsMap() (adt.Map, error) {

	return adt2.AsMap(s.store, s.VerifiedClients)

}

func (s *state2) VerifiersMap() (adt.Map, error) {
	return adt2.AsMap(s.store, s.Verifiers)
}

func (s *state2) VerifiersMapBitWidth() int {

	return 5

}

func (s *state2) VerifiersMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state2) AllocationsMap() (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) AllocationsMapBitWidth() int {

	return 5

}

func (s *state2) AllocationsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state2) AllocationMapForClient(clientIdAddr address.Address) (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) RootKey() (address.Address, error) {
	return s.State.RootKey, nil
}

func (s *state2) VerifiedClientDataCap(addr address.Address) (bool, abi.StoragePower, error) {

	return getDataCap(s.store, actors.Version2, s.VerifiedClientsMap, addr)

}

func (s *state2) VerifierDataCap(addr address.Address) (bool, abi.StoragePower, error) {
	return getDataCap(s.store, actors.Version2, s.VerifiersMap, addr)
}

func (s *state2) RemoveDataCapProposalID(verifier address.Address, client address.Address) (bool, uint64, error) {
	return getRemoveDataCapProposalID(s.store, actors.Version2, s.removeDataCapProposalIDs, verifier, client)
}

func (s *state2) ForEachVerifier(cb func(addr address.Address, dcap abi.StoragePower) error) error {
	return forEachCap(s.store, actors.Version2, s.VerifiersMap, cb)
}

func (s *state2) ForEachClient(cb func(addr address.Address, dcap abi.StoragePower) error) error {

	return forEachCap(s.store, actors.Version2, s.VerifiedClientsMap, cb)

}

func (s *state2) removeDataCapProposalIDs() (adt.Map, error) {
	return nil, nil

}

func (s *state2) GetState() interface{} {
	return &s.State
}

func (s *state2) GetAllocation(clientIdAddr address.Address, allocationId verifreg9.AllocationId) (*verifreg9.Allocation, bool, error) {

	return nil, false, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) GetAllocations(clientIdAddr address.Address) (map[verifreg9.AllocationId]verifreg9.Allocation, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) GetClaim(providerIdAddr address.Address, claimId verifreg9.ClaimId) (*verifreg9.Claim, bool, error) {

	return nil, false, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) GetClaims(providerIdAddr address.Address) (map[verifreg9.ClaimId]verifreg9.Claim, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) ClaimsMap() (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

// TODO this could return an error since not all versions have a claims map
func (s *state2) ClaimsMapBitWidth() int {

	return 5

}

// TODO this could return an error since not all versions have a claims map
func (s *state2) ClaimsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state2) ClaimMapForProvider(providerIdAddr address.Address) (adt.Map, error) {

	return nil, fmt.Errorf("unsupported in actors v2")

}

func (s *state2) getInnerHamtCid(store adt.Store, key abi.Keyer, mapCid cid.Cid, bitwidth int) (cid.Cid, error) {

	return cid.Undef, fmt.Errorf("unsupported in actors v2")

}
