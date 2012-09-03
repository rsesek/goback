package main

import (
	"fmt"
)

func main() {
	config, _ := loadConfiguration()
	backups, err := readAllBackups(config)
	fmt.Printf("%v\n%+v\n", err, backups)

	backup, err := newActiveBackup(config, createBackup(config), nil)
	fmt.Printf("\n%v\n", err)
	fmt.Println(backup.execute())
	fmt.Println(<-backup.fin)
	fmt.Println(writeBackupMetadata(backup.b))
}
