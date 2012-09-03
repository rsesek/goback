package common

import (
	"time"
)

type BackupState int

const (
	BackupStateUnknown BackupState = iota
	BackupStateInProgress
	BackupStateCompleted
	BackupStateDiscarded
)

// Backup holds the metadata state about an individual backup instance.
type Backup struct {
	// The full path to the backup.
	FullPath string
	// The time at which the backup was started.
	StartTime time.Time
	// The time at which the backup completed.
	EndTime time.Time
	// The state of the backup.
	State BackupState
}
