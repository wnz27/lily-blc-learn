package v2

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lily/tasks"
	"github.com/filecoin-project/lily/tasks/actorstate"
)

type ModelVersion int
type ModelType string
type ModelKind string

const (
	ModelActorKind ModelKind = "actor"
	ModelTsKind    ModelKind = "tipset"
)

type ModelMeta struct {
	Version ModelVersion
	Type    ModelType
	Kind    ModelKind
}

const modelMetaVersionSeparator = "@v"
const modelMetaKindSeparator = ":"

func (mm ModelMeta) String() string {
	return fmt.Sprintf("%s%s%s%s%d", mm.Kind, modelMetaKindSeparator, mm.Type, modelMetaVersionSeparator, mm.Version)
}

func (mm ModelMeta) Equals(other ModelMeta) bool {
	if mm.Kind != other.Kind {
		return false
	}
	if mm.Type != other.Type {
		return false
	}
	if mm.Version != other.Version {
		return false
	}
	return true
}

func DecodeModelMeta(str string) (ModelMeta, error) {
	tokens := strings.Split(str, modelMetaKindSeparator)
	if len(tokens) != 2 {
		return ModelMeta{}, fmt.Errorf("invalid")
	}
	kind := tokens[0]
	tokens = strings.Split(tokens[1], modelMetaVersionSeparator)
	if len(tokens) != 2 {
		return ModelMeta{}, fmt.Errorf("invalid")
	}
	mv, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return ModelMeta{}, err
	}
	mt := tokens[0]
	return ModelMeta{
		Version: ModelVersion(mv),
		Type:    ModelType(mt),
		Kind:    ModelKind(kind),
	}, nil
}

type LilyModel interface {
	cbor.Er
	Cid() cid.Cid
	Meta() ModelMeta
	ChainEpochTime() ChainEpochTime
}

type ChainEpochTime struct {
	Height    abi.ChainEpoch
	StateRoot cid.Cid
}

// TODO consider making registry functions generic

var ActorExtractorRegistry map[ModelMeta]ActorExtractorFn
var ActorTypeRegistry map[ModelMeta]*cid.Set
var ExtractorRegistry map[ModelMeta]ExtractorFn
var ModelReflections map[ModelMeta]ModelReflect

type ModelReflect struct {
	Meta ModelMeta
	Type reflect.Type
}

func init() {
	ActorExtractorRegistry = make(map[ModelMeta]ActorExtractorFn)
	ActorTypeRegistry = make(map[ModelMeta]*cid.Set)
	ExtractorRegistry = make(map[ModelMeta]ExtractorFn)
	ModelReflections = make(map[ModelMeta]ModelReflect)
}

type ActorExtractorFn func(ctx context.Context, api tasks.DataSource, current, executed *types.TipSet, a actorstate.ActorInfo) ([]LilyModel, error)

// RegisterActorExtractor associates a model with extractor that produces it.
func RegisterActorExtractor(model LilyModel, efn ActorExtractorFn) {
	ActorExtractorRegistry[model.Meta()] = efn
	ModelReflections[model.Meta()] = ModelReflect{
		Meta: model.Meta(),
		Type: reflect.TypeOf(model),
	}
}

func RegisterActorType(model LilyModel, actors *cid.Set) {
	ActorTypeRegistry[model.Meta()] = actors
}

type ExtractorFn func(ctx context.Context, api tasks.DataSource, current, executed *types.TipSet) ([]LilyModel, error)

func RegisterExtractor(model LilyModel, efn ExtractorFn) {
	ExtractorRegistry[model.Meta()] = efn
	ModelReflections[model.Meta()] = ModelReflect{
		Meta: model.Meta(),
		Type: reflect.TypeOf(model),
	}
}

func LookupExtractor(meta ModelMeta) (ExtractorFn, error) {
	efn, found := ExtractorRegistry[meta]
	if !found {
		return nil, fmt.Errorf("no extractor for %s", meta)
	}
	return efn, nil
}

func LookupActorExtractor(meta ModelMeta) (ActorExtractorFn, error) {
	efn, found := ActorExtractorRegistry[meta]
	if !found {
		return nil, fmt.Errorf("no extractor for %s", meta)
	}
	return efn, nil
}

func MustLookupActorTypeThing(meta ModelMeta) *cid.Set {
	actors, found := ActorTypeRegistry[meta]
	if !found {
		panic(fmt.Sprintf("no actors for %s developer error", meta))
	}
	return actors
}
