package main

type Configuration struct {
	BackupDestination string

	BackupSource []string
}

func loadConfiguration() (*Configuration, error) {
	return &Configuration{
		BackupDestination: "/Volumes/Build/gopath/src/github.com/rsesek/_dest",
		BackupSource:      []string{"/Volumes/Build/gopath/src/github.com/rsesek/goback/"},
	}, nil
}
