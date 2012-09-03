package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/rsesek/goback/common"
)

type activeBackup struct {
	b common.Backup

	args []string
	cmd  *exec.Cmd

	stdout *os.File
	stderr *os.File

	fin chan error
}

func newActiveBackup(config *Configuration, current common.Backup, latestComplete *common.Backup) (*activeBackup, error) {
	if current.FullPath == "" {
		return nil, errors.New("current backup does not have a FullPath")
	}

	args := []string{
		"--archive",
		"--verbose",
	}
	if latestComplete != nil {
		args = append(args, "--link-dest", latestComplete.FullPath)
	}
	args = append(args, config.BackupSource, path.Join(current.FullPath, "src"))

	return &activeBackup{
		b:    current,
		args: args,
		fin:  make(chan error),
	}, nil
}

func (ab *activeBackup) execute() (err error) {
	err = os.Mkdir(ab.b.FullPath, 0744)
	if err != nil {
		return
	}

	ab.stdout, err = os.Create(path.Join(ab.b.FullPath, "rsync.OUT"))
	if err != nil {
		return fmt.Errorf("creating rsync stdout: %v", err)
	}

	ab.stderr, err = os.Create(path.Join(ab.b.FullPath, "rsync.ERR"))
	if err != nil {
		return fmt.Errorf("creating rsync stderr: %v", err)
	}

	ab.cmd = exec.Command("rsync", ab.args...)
	ab.cmd.Stdout = ab.stdout
	ab.cmd.Stderr = ab.stderr
	err = ab.cmd.Start()
	go ab.executeSync()
	return
}

func (ab *activeBackup) executeSync() {
	err := ab.cmd.Wait()
	ab.stdout.Close()
	ab.stderr.Close()
	if ab.cmd.ProcessState.Success() {
		ab.b.State = common.BackupStateCompleted
	} else {
		ab.b.State = common.BackupStateFailed
	}
	ab.fin <- err
}
