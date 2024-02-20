package sync

import (
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/polygon/heimdall"
)

//go:generate mockgen -destination=./storage_mock.go -package=sync -source=./storage.go
type HeaderStore interface {
	PutHeaders(headers []*types.Header) error
}

type CheckpointStore interface {
	HeaderStore
	heimdall.CheckpointStore
}

type MilestoneStore interface {
	HeaderStore
	heimdall.MilestoneStore
}
