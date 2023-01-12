package v9

import (
	"context"

	"github.com/filecoin-project/go-address"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/model"
	minermodel "github.com/filecoin-project/lily/model/actors/miner"
	"github.com/filecoin-project/lily/pkg/core"
	v9 "github.com/filecoin-project/lily/pkg/extract/actors/minerdiff/v9"
)

type Sector struct {
}

func (s Sector) Extract(ctx context.Context, current, executed *types.TipSet, addr address.Address, change *v9.StateDiffResult) (model.Persistable, error) {
	var sectors []*miner9.SectorOnChainInfo
	changes := change.SectorChanges
	for _, sector := range changes {
		// only care about modified and added sectors
		if sector.Change == core.ChangeTypeRemove {
			continue
		}
		// change.Current is the newly added sector, or its state after modification.
		if err := core.StateReadDeferred(ctx, sector.Current, func(sector *miner9.SectorOnChainInfo) error {
			sectors = append(sectors, sector)
			return nil
		}); err != nil {
			return nil, err
		}
	}
	return MinerSectorChangesAsModel(ctx, current, addr, sectors)
}

func MinerSectorChangesAsModel(ctx context.Context, current *types.TipSet, addr address.Address, sectors []*miner9.SectorOnChainInfo) (model.Persistable, error) {
	sectorModel := make(minermodel.MinerSectorInfoV7List, len(sectors))
	for i, sector := range sectors {
		sectorKeyCID := ""
		if sector.SectorKeyCID != nil {
			sectorKeyCID = sector.SectorKeyCID.String()
		}
		sectorModel[i] = &minermodel.MinerSectorInfoV7{
			Height:                int64(current.Height()),
			MinerID:               addr.String(),
			StateRoot:             current.ParentState().String(),
			SectorID:              uint64(sector.SectorNumber),
			SealedCID:             sector.SealedCID.String(),
			ActivationEpoch:       int64(sector.Activation),
			ExpirationEpoch:       int64(sector.Expiration),
			DealWeight:            sector.DealWeight.String(),
			VerifiedDealWeight:    sector.VerifiedDealWeight.String(),
			InitialPledge:         sector.InitialPledge.String(),
			ExpectedDayReward:     sector.ExpectedDayReward.String(),
			ExpectedStoragePledge: sector.ExpectedStoragePledge.String(),
			SectorKeyCID:          sectorKeyCID,
		}
	}

	return sectorModel, nil
}