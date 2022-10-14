package economics

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"reflect"

	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lily/chain/indexer/v2/transform"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform/persistable"
	messagemodel "github.com/filecoin-project/lily/model/messages"
	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/model/v2/economics"
)

var log = logging.Logger("transform/economics")

type GasEconomyTransform struct {
	meta v2.ModelMeta
}

func NewGasEconomyTransform() *GasEconomyTransform {
	info := economics.ChainEconomics{}
	return &GasEconomyTransform{meta: info.Meta()}
}

func (g *GasEconomyTransform) Run(ctx context.Context, in chan transform.IndexState, out chan transform.Result) error {
	log.Debugf("run %s", g.Name())
	for res := range in {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			log.Debugw("received data", "count", len(res.Models()))
			// add up total and unique gas
			for _, modeldata := range res.Models() {
				m := modeldata.(*economics.ChainEconomics)

				baseFeeRat := new(big.Rat).SetFrac(m.BaseFee.Int, new(big.Int).SetUint64(build.FilecoinPrecision))
				baseFee, _ := baseFeeRat.Float64()

				baseFeeChange := new(big.Rat).SetFrac(m.BaseFee.Int, m.ParentBaseFee.Int)
				baseFeeChangeF, _ := baseFeeChange.Float64()

				sqlModel := &messagemodel.MessageGasEconomy{
					Height:              int64(m.Height),
					StateRoot:           m.StateRoot.String(),
					GasLimitTotal:       m.TotalGasLimit,
					GasLimitUniqueTotal: m.TotalUniqueGasLimit,
					BaseFee:             baseFee,
					BaseFeeChangeLog:    math.Log(baseFeeChangeF) / math.Log(1.125),
					GasFillRatio:        float64(m.TotalGasLimit) / float64(m.NumBlocks*build.BlockGasTarget),
					GasCapacityRatio:    float64(m.TotalUniqueGasLimit) / float64(m.NumBlocks*build.BlockGasTarget),
					GasWasteRatio:       float64(m.TotalGasLimit-m.TotalUniqueGasLimit) / float64(m.NumBlocks*build.BlockGasTarget),
				}
				out <- &persistable.Result{Model: sqlModel}
			}
		}
	}
	return nil
}

func (g *GasEconomyTransform) Name() string {
	info := GasEconomyTransform{}
	return reflect.TypeOf(info).Name()
}

func (g *GasEconomyTransform) ModelType() v2.ModelMeta {
	return g.meta
}

func (g *GasEconomyTransform) Matcher() string {
	return fmt.Sprintf("^%s$", g.meta.String())
}