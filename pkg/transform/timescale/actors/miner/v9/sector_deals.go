package v9

import (
	"bytes"
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/model"
	minermodel "github.com/filecoin-project/lily/model/actors/miner"
	"github.com/filecoin-project/lily/pkg/core"
	v9 "github.com/filecoin-project/lily/pkg/extract/actors/minerdiff/v9"
)

type SectorDeal struct{}

func (s SectorDeal) Extract(ctx context.Context, current, executed *types.TipSet, addr address.Address, change *v9.StateDiffResult) (model.Persistable, error) {
	sectors := change.SectorChanges
	out := minermodel.MinerSectorDealList{}
	height := int64(current.Height())
	minerAddr := addr.String()
	for _, sector := range sectors {
		switch sector.Change {
		case core.ChangeTypeAdd:
			if err := core.StateReadDeferred(ctx, sector.Current, func(s *miner9.SectorOnChainInfo) error {
				for _, deal := range s.DealIDs {
					out = append(out, &minermodel.MinerSectorDeal{
						Height:   height,
						MinerID:  minerAddr,
						SectorID: uint64(s.SectorNumber),
						DealID:   uint64(deal),
					})
				}
				return nil
			}); err != nil {
				return nil, err
			}
		case core.ChangeTypeModify:
			previousSector := new(miner9.SectorOnChainInfo)
			if err := previousSector.UnmarshalCBOR(bytes.NewReader(sector.Previous.Raw)); err != nil {
				return nil, err
			}
			currentSector := new(miner9.SectorOnChainInfo)
			if err := currentSector.UnmarshalCBOR(bytes.NewReader(sector.Current.Raw)); err != nil {
				return nil, err
			}
			for _, deal := range compareDealIDs(currentSector.DealIDs, previousSector.DealIDs) {
				out = append(out, &minermodel.MinerSectorDeal{
					Height:   height,
					MinerID:  minerAddr,
					SectorID: uint64(currentSector.SectorNumber),
					DealID:   uint64(deal),
				})
			}
		}
	}
	return out, nil
}

func compareDealIDs(cur, pre []abi.DealID) []abi.DealID {
	var diff []abi.DealID

	// Loop two times, first to find cur dealIDs not in pre,
	// second loop to find pre dealIDs not in cur
	for i := 0; i < 2; i++ {
		for _, s1 := range cur {
			found := false
			for _, s2 := range pre {
				if s1 == s2 {
					found = true
					break
				}
			}
			// DealID not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			cur, pre = pre, cur
		}
	}

	return diff
}