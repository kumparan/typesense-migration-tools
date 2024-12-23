package config

import "time"

const (
	DefaultHTTPTimeout             = 10 * time.Second
	DefaultHTTPTLSHandshakeTimeout = 5 * time.Second
	DefaultTLSInsecureSkipVerify   = true

	DefaultMigrationSleepInterval             = time.Second
	DefaultBackupSleepInterval                = time.Second
	DefaultRestoreSleepInterval               = time.Second
	DefaultSleepIntervalForCollectionDeletion = time.Second

	DefaultMigrationBatchSize             = 100
	DefaultBackupBatchSize                = 100
	DefaultRestoreBatchSize               = 100
	DefaultBatchSizeForCollectionDeletion = 100
)
