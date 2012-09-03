package main

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/rsesek/goback/common"
)

const (
	kMetadataFile = "info"

	kBackupDateFormat = "2006-01-02-150405"
)

func readAllBackups(c *Configuration) ([]common.Backup, error) {
	f, err := os.Open(c.BackupDestination)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	backups := make([]common.Backup, 0)
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		fullPath := path.Join(c.BackupDestination, file.Name())
		backup, err := readBackupMetadata(fullPath)
		if err != nil {
			// TODO: log.Error
			continue
		}

		backup.FullPath = fullPath
		backups = append(backups, backup)
	}
	return backups, nil
}

func readBackupMetadata(backupPath string) (backup common.Backup, err error) {
	mdpath := path.Join(backupPath, kMetadataFile)
	f, err := os.Open(mdpath)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(backup)
	return
}

func writeBackupMetadata(backup common.Backup) error {
	mdpath := path.Join(backup.FullPath, kMetadataFile)
	f, err := os.Open(mdpath)
	if os.IsNotExist(err) {
		f, err = os.Create(mdpath)
	}
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(backup)
}

func createBackup(config *Configuration) common.Backup {
	t := time.Now()
	return common.Backup{
		FullPath:  path.Join(config.BackupDestination, t.Format(kBackupDateFormat)),
		StartTime: t,
	}
}
