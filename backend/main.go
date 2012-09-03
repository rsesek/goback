package main

import (
	"fmt"
)

func main() {
	config, _ := loadConfiguration()
	backups, err := readAllBackups(config)
	fmt.Printf("%v\n%+v\n", err, backups)
}
