package models

import (
	"encoding/json"
	"time"
)

type MaintenanceInfo struct {
	Owner             string              `json:"owner"`
	QuickCycle        CycleParams         `json:"quick"`
	FullCycle         CycleParams         `json:"full"`
	LogRetention      LogRetentionOptions `json:"logRetention"`
	ExtendObjectLocks bool                `json:"extendObjectLocks"`
	ListParallelism   int                 `json:"listParallelism"`
	Schedule          Schedule            `json:"schedule"`
}

type Params struct {
	Owner             string              `json:"owner"`
	QuickCycle        CycleParams         `json:"quick"`
	FullCycle         CycleParams         `json:"full"`
	LogRetention      LogRetentionOptions `json:"logRetention"`
	ExtendObjectLocks bool                `json:"extendObjectLocks"`
	ListParallelism   int                 `json:"listParallelism"`
}
type CycleParams struct {
	Enabled  bool          `json:"enabled"`
	Interval time.Duration `json:"interval"`
}
type LogRetentionOptions struct {
	MaxTotalSize int64         `json:"maxTotalSize"`
	MaxCount     int           `json:"maxCount"`
	MaxAge       time.Duration `json:"maxAge"`
}
type Schedule struct {
	NextFullMaintenanceTime  time.Time `json:"nextFullMaintenance"`
	NextQuickMaintenanceTime time.Time `json:"nextQuickMaintenance"`

	Runs map[TaskType][]RunInfo `json:"runs"`
}

type TaskType string

// Task IDs.
const (
	TaskSnapshotGarbageCollection    = "snapshot-gc"
	TaskDeleteOrphanedBlobsQuick     = "quick-delete-blobs"
	TaskDeleteOrphanedBlobsFull      = "full-delete-blobs"
	TaskRewriteContentsQuick         = "quick-rewrite-contents"
	TaskRewriteContentsFull          = "full-rewrite-contents"
	TaskDropDeletedContentsFull      = "full-drop-deleted-content"
	TaskIndexCompaction              = "index-compaction"
	TaskExtendBlobRetentionTimeFull  = "extend-blob-retention-time"
	TaskCleanupLogs                  = "cleanup-logs"
	TaskEpochAdvance                 = "advance-epoch"
	TaskEpochDeleteSupersededIndexes = "delete-superseded-epoch-indexes"
	TaskEpochCleanupMarkers          = "cleanup-epoch-markers"
	TaskEpochGenerateRange           = "generate-epoch-range-index"
	TaskEpochCompactSingle           = "compact-single-epoch"
)

type RunInfo struct {
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Success bool      `json:"success,omitempty"`
	Error   string    `json:"error,omitempty"`
	Extra   []Extra   `json:"extra,omitempty"`
}

type Extra struct {
	Kind string          `json:"kind,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}
