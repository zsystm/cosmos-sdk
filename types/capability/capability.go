package capability

import (
	"cosmossdk.io/core/blockinfo"
	"cosmossdk.io/core/event"
	"cosmossdk.io/core/gas"
	"cosmossdk.io/core/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

const BlockInfoCapabilityKey string = "BlockInfoService"

type BlockInfoService interface {
	BlockInfoService() blockinfo.Service
}

const KVStoreCapabilityKey string = "KVStoreService"

type KVStoreService interface {
	KVStoreService(key storetypes.KVStoreKey) store.KVStoreService
}

const EventCapabilityKey string = "EventService"

type EventService interface {
	EventService() event.Service
}

const GasCapabilityKey string = "GasService"

type GasService interface {
	GasService() gas.Service
}
